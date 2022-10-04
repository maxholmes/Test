package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to DB first
	env := new(Env)
	var err error
	env.DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	router := gin.Default()

	router.GET("/Futbol_Fun/:id", env.GetMemberByID)
	router.GET("/Futbol_Fun", env.GetMembers)
	router.POST("/Futbol_Fun", env.PostMember)
	router.PUT("/Futbol_Fun", env.UpdateMember)
	router.DELETE("/Futbol_Fun/:id", env.DeleteMemberByID)

	router.Run("localhost:8080")
}
