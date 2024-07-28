package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", helloWorldHandler)
	server.GET("/home", homeHandler)
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvents)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
