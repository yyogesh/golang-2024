package main

import (
	"net/http"
	"restapi_example/db"
	"restapi_example/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() { //com component, remoting, webservice soap based service, rest service, graphql, gRpc
	db.InitDB()
	server := gin.Default()

	server.GET("/", helloWorldHandler)
	server.GET("/home", homeHandler)
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/events", createEvents)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) { // request postman body raw json format
	var event models.Event
	err := context.ShouldBindJSON(&event) // fmt.scanF

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) // 201 Created status code

}

func helloWorldHandler(context *gin.Context) { // request, response
	context.String(200, "Hello, World!")
}

func homeHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
