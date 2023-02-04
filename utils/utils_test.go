package utils_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashAndSalt(t *testing.T) {
	t.Run("should return hashed password", func(t *testing.T) {
		pw := "password"
		
		result, _ := utils.HashAndSalt(pw)

		assert.NotNil(t, result)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("should return true when password correct", func(t *testing.T) {
		pw := "password"
		hash, _ := utils.HashAndSalt(pw)
		
		result := utils.ComparePassword(hash, pw)

		assert.True(t, result)
	})
}

func TestGenerateAccessToken(t *testing.T) {
	t.Run("should return token with claim", func(t *testing.T) {
		wallet := &entity.Wallet{
			ID: 1,
		}
		user := &entity.User{
			ID: 1,
			Email: "user@mail.com",
			RoleID: 2,
			Wallet: wallet,
		}

		result, _ := utils.GenerateAccessToken(user)

		assert.NotNil(t, result)
	})
}