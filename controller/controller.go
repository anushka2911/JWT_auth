package controller

import (
	"fmt"
	"log"

	"github.com/anushka2911/jwt/middleware"
	"github.com/anushka2911/jwt/model"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
}

func Signup(c *gin.Context) {
	middleware.ValidateUserInput(c, &model.User{})
	log.Println("I am in signup and validated user input")
}

func GetAllUsers(c *gin.Context) {
	fmt.Println("I am in get all users")
}

func GetUser(c *gin.Context) {
}
