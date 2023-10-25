package middleware

import (
	"log"
	"net/http"

	"github.com/anushka2911/jwt/model"
	"github.com/gin-gonic/gin"
)

func ValidateUserInput(c *gin.Context, user *model.User) {
	err := c.ShouldBindJSON(user)
	if err != nil {
		log.Printf("Error in binding user data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidateUserInput()
	if err != nil {
		log.Printf("Validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
