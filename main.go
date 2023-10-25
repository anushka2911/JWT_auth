package main

import (
	"log"

	"github.com/anushka2911/jwt/database"
	"github.com/anushka2911/jwt/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
}

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	routes.AuthRoutes(router)

	port := ":8080"
	log.Printf("Server running on port %s...\n", port)

	err := router.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
