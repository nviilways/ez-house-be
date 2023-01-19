package middleware

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthorization(c *gin.Context) {
	auth := c.GetHeader("authorization")

	if auth == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code" : "STATUS_UNAUTHORIZED",
			"message" : "unable to find user auth",
		})
		return
	}

	auth = strings.Replace(auth, "Bearer ", "", -1)

	parsedToken, err := jwt.ParseWithClaims(auth, &entity.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code" : "STATUS_UNAUTHORIZED",
			"message" : "unable to parse token",
		})
	}

	claim, ok := parsedToken.Claims.(*entity.Claim)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code" : "STATUS_UNAUTHORIZED",
			"message" : "invalid token",
		})
		return
	}

	c.Set("claim", *claim)
}