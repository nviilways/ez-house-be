package middleware

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	errs "git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/error"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthorization(c *gin.Context) {
	auth := c.GetHeader("authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.JSONResponse{
			Code: http.StatusUnauthorized,
			Message: errs.ErrorCode[http.StatusUnauthorized],
			Data: nil,
		})
		return
	}

	auth = strings.Replace(auth, "Bearer ", "", -1)

	parsedToken, err := jwt.ParseWithClaims(auth, &entity.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.JSONResponse{
			Code: http.StatusUnauthorized,
			Message: errs.ErrorCode[http.StatusUnauthorized],
			Data: nil,
		})
		return
	}

	claim, ok := parsedToken.Claims.(*entity.Claim)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.JSONResponse{
			Code: http.StatusUnauthorized,
			Message: errs.ErrorCode[http.StatusUnauthorized],
			Data: nil,
		})
		return
	}

	c.Set("claim", *claim)
}