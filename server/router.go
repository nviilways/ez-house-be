package server

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/handler"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/middleware"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUsecase usecase.UserUsecase
	TransactionUsecase usecase.TransactionUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
	})
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST", "GET"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
	}))

	v1API := router.Group("/api/v1")
	{
		v1API.POST("/register", h.UserRegister)
		v1API.POST("/login", h.UserLogin)
		v1API.GET("/me", middleware.JWTAuthorization, h.UserDetails)
		v1API.PATCH("/update", middleware.JWTAuthorization, h.UserUpdate)
		v1API.POST("/logout", middleware.JWTAuthorization, h.UserLogout)
		v1API.POST("/topup", middleware.JWTAuthorization, h.UserTopUp)
	}

	return router
}

