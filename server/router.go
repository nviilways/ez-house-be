package server

import (
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/handler"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
	})
	
	v1API := router.Group("/api/v1")
	{
		v1API.Group("register", h.UserRegister)
	}

	return router
}

