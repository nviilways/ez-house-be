package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/dto"
	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/entity"
	"github.com/gin-gonic/gin"
)

const superAdminRoleId = 4
const hostRoleId = 3

func AuthorizeAdmin(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	if(parsedClaim.RoleID != superAdminRoleId){
		c.AbortWithStatusJSON(http.StatusForbidden, dto.JSONResponse{
			Code: http.StatusForbidden,
			Message: "forbidden access",
		})
	}
}

func AuthorizeHost(c *gin.Context) {
	claim, _ := c.Get("claim")
	parsedClaim := entity.Claim(claim.(entity.Claim))

	if(parsedClaim.RoleID != hostRoleId){
		c.AbortWithStatusJSON(http.StatusForbidden, dto.JSONResponse{
			Code: http.StatusForbidden,
			Message: "forbidden access",
		})
	}
}