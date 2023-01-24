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
	AdminUsecase usecase.AdminUsecase
	HouseUsecase usecase.HouseUsecase
	ReservationUsecase usecase.ReservationUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		UserUsecase: cfg.UserUsecase,
		TransactionUsecase: cfg.TransactionUsecase,
		AdminUsecase: cfg.AdminUsecase,
		HouseUsecase: cfg.HouseUsecase,
		ReservationUsecase: cfg.ReservationUsecase,
	})
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST", "GET", "PATCH"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
	}))

	v1API := router.Group("/api/v1")
	{
		admin := v1API.Group("/admin")
		{
			admin.POST("/login", h.AdminSignIn)
			admin.POST("/register", middleware.JWTAuthorization, middleware.AuthorizeAdmin, h.AdminSignUp)
		}
		house := v1API.Group("/house")
		{
			house.GET("/:id", h.UserGetHouseById)
			house.POST("/add", middleware.JWTAuthorization, middleware.AuthorizeHost, h.HostAddHouse)
			house.DELETE("/:id", middleware.JWTAuthorization, middleware.AuthorizeAdminOrHost, h.HostDeleteHouse)
		}
		reservation := v1API.Group("/reservations")
		{
			reservation.POST("/create", middleware.JWTAuthorization, h.UseAddReservation)
		}
		v1API.POST("/register", h.UserRegister)
		v1API.POST("/login", h.UserLogin)
		v1API.GET("/me", middleware.JWTAuthorization, h.UserDetails)
		v1API.PATCH("/update", middleware.JWTAuthorization, h.UserUpdate)
		v1API.POST("/logout", middleware.JWTAuthorization, h.UserLogout)
		v1API.POST("/topup", middleware.JWTAuthorization, h.UserTopUp)
		v1API.PATCH("/update/role", middleware.JWTAuthorization, h.UserUpdateRole)
	}

	return router
}

