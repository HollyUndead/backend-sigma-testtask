package main

import (
	"log"
	. "testtry2/database"
	. "testtry2/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	defer Disconnect()
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/users", GetUsersList)
	router.GET("/userId/:Id", FindUserById)
	router.GET("/userName/:FirstName", FindUserByFirstName)
	router.POST("/createUser", CreateUser)
	router.PUT("/updateUser", UpdateUser)
	router.DELETE("/deletUser/:Id", DeleteUser)

	router.Run("localhost:8080")
}
