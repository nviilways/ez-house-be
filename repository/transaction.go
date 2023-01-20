package repository

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"gorm.io/gorm"
)

const MinTopUpGames = 500000

type TransactionRepository interface {
	TopUp(int, *entity.Transaction) (*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(cfg *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl {
		db: cfg.DB,
	}
}

func (t *transactionRepositoryImpl) TopUp(target int, tr *entity.Transaction) (*entity.Transaction, error) {
	var wallet entity.Wallet
	walletErr := t.db.Where("user_id = ?", target).First(&wallet).Error
	if walletErr != nil {
		return nil, walletErr
	}

	tr.WalletID = wallet.ID

	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&tr).Error; err != nil {
			return err
		}

		if err := tx.Model(&entity.Wallet{}).Where("id = ?", tr.WalletID).Update("balance", gorm.Expr("balance + ?", tr.Balance)).Error; err != nil {
			return err
		}

		gameBonus := tr.Balance / MinTopUpGames
		if(gameBonus > 0) {
			if err := tx.Model(&entity.Game{}).Where("user_id = ?", target).Update("chance", gorm.Expr("chance + ?", gameBonus)).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tr, nil
}