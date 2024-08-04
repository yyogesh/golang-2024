package routes

import (
	"restapi_example/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", helloWorldHandler)
	server.GET("/home", homeHandler)
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	autenticated := server.Group("/")
	autenticated.Use(middlewares.Authenticate)
	autenticated.POST("/events", createEvents)
	autenticated.PUT("/events/:id", updateEvents)
	autenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
