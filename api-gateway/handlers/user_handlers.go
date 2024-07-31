package handlers

import (
	us "api-gateway/proto/user"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userID := c.GetString("userID")

	resp, err := userService.GetUserProfile(context.Background(), &us.GetUserProfileRequest{
		UserId: userID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
