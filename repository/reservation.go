package repository

import (
	"fmt"
	"math"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"gorm.io/gorm"
)

const hourPerDay = 24
const debitType = 2
const commissionType = 4
const commissionPay = 0.8

type ReservationRepository interface {
	AddReservation(*entity.Reservation) (*entity.Reservation, error)
}

type reservationRepositoryImpl struct {
	db *gorm.DB
}

type ReservationRConfig struct {
	DB *gorm.DB
}

func NewReservationRepository(cfg *ReservationRConfig) ReservationRepository {
	return &reservationRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *reservationRepositoryImpl) ValidateReservation(checkIn time.Time, checkOut time.Time, house_id uint) error {
	if(checkIn.Before(time.Now())) {
		return errs.ErrInvalidCheckDate
	}

	if checkOut.Sub(checkIn).Hours() / hourPerDay < 1 {
		return errs.ErrBookForZeroDays
	}

	err := r.db.Where("check_in_date <= ? AND check_out_date > ?", checkOut, checkIn).First(&entity.Reservation{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		return err
	}

	return errs.ErrAlreadyReserved
}

func (r *reservationRepositoryImpl) ValidateHouse(house_id uint) (*entity.House, error) {
	var house entity.House
	err := r.db.Where("id = ?", house_id).First(&house).Error
	if err != nil {
		return nil, err
	}

	return &house, nil
}

func (r *reservationRepositoryImpl) ValidateUserBalance(user_id uint, price int) (*entity.Wallet, error) {
	var wallet entity.Wallet
	err := r.db.Where("user_id = ?", user_id).First(&wallet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, err
	}

	if(wallet.Balance >= price) {
		return &wallet, nil
	}

	return nil, errs.ErrInsufficientBalance
}

func (r *reservationRepositoryImpl) GetHostWallet(user_id uint) (*entity.Wallet, error) {
	var wallet *entity.Wallet
	err := r.db.Where("user_id = ?", user_id).First(&wallet).Error
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (r *reservationRepositoryImpl) AddReservation(res *entity.Reservation) (*entity.Reservation, error) {
	res.CommissionStatus = "PENDING"
	res.BookingCode = fmt.Sprintf("BOOK-EZ-%d-%d", res.UserId, time.Now().UnixMicro())
	err := r.ValidateReservation(res.CheckInDate, res.CheckOutDate, res.HouseId)
	if err != nil {
		return nil, err
	}

	house, err := r.ValidateHouse(res.HouseId)
	if err != nil {
		return nil, errs.ErrHouseNotFound
	}

	if house.UserID == res.UserId {
		return nil, errs.ErrReserveOwnHouse
	}
	
	totalNight := res.CheckOutDate.Sub(res.CheckInDate).Hours() / hourPerDay
	totalPrice := house.Price * int(totalNight)

	wallet, walletErr := r.ValidateUserBalance(res.UserId, totalPrice)
	if walletErr != nil {
		return nil, errs.ErrInsufficientBalance
	}

	hostWallet, hwErr := r.GetHostWallet(house.UserID)
	if hwErr != nil {
		return nil, hwErr
	}

	res.TotalPrice = totalPrice

	err2 := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&res).Error; err != nil {
			return err
		}

		transaction := &entity.Transaction{
			WalletID: wallet.ID,
			TransactionTypeID: debitType,
			Balance: totalPrice,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("id = ?", wallet.ID).Update("balance", gorm.Expr("balance - ?", totalPrice)).Error; err != nil {
			return err
		}

		floatPrice := float64(totalPrice)

		commission := &entity.Transaction{
			WalletID: hostWallet.ID,
			TransactionTypeID: commissionType,
			Balance: int(math.Round(floatPrice * commissionPay)),
		}

		if err := tx.Create(&commission).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("user_id = ?", house.UserID).Update("balance", gorm.Expr("balance + ?", commission.Balance)).Error; err != nil {
			return err
		}

		return nil
	})

	if err2 != nil {
		return nil, err2
	}

	return res, nil
}
