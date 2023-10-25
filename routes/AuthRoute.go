package routes

import (
	controller "github.com/anushka2911/jwt/controller"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("user/signup", controller.Signup)
	router.POST("user/login", controller.Login)
}
