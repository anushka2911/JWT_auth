package routes

import (
	controller "github.com/anushka2911/jwt/controller"
	middleware "github.com/anushka2911/jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.USE(middleware.AuthenticateUser())
	router.GET("/users", controller.GetAllUsers())
	router.GET("/user/:id", controller.GetUser())
}
