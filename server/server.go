package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/db"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/repository"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	db := db.Get()
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db,
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepository: userRepo,
	})

	txRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB: db,
	})
	txUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
		TransactionRepository: txRepo,
	})

	return NewRouter(&RouterConfig{
		UserUsecase: userUsecase,
		TransactionUsecase: txUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}