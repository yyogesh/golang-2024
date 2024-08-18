package main

import (
	"job_portal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	// server.GET("/hello", helloWorld);

	server.Run(":8080")
}
