package main

import (
	"github.com/anushka2911/jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
}
