package utils

import (
	"strings"
	"time"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/config"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	
	return string(hash), nil
}

func ComparePassword(hashedPwd string, inputPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(inputPwd))
	return err == nil
}

func GenerateAccessToken(user *entity.User) (*dto.Token, error) {
	claims := &entity.Claim{
		ID: user.ID,
		Email: user.Email,
		RoleID: user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "Ez House",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.SecretKey))

	result := &dto.Token{
		AccessToken: tokenString,
	}

	return result, nil
}

func GetToken(c *gin.Context) string {
	token := c.GetHeader("authorization")
	token = strings.Replace(token, "Bearer ", "", -1)
	return token 
}