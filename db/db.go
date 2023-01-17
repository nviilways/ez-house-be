package db

import (
	"fmt"
	"log"
	"os"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	cfg = config.DBConfig
	db *gorm.DB
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", 
	cfg.Host,
	cfg.User,
	cfg.Password,
	cfg.DBName,
	cfg.Port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	Logger: getLogger(),
	})

	return
}

func Get() *gorm.DB {
	return db
}