package main

import (
	"net/http"
	"restapi_example/models"

	"github.com/gin-gonic/gin"
)

func main() { //com component, remoting, webservice soap based service, rest service, graphql, gRpc
	server := gin.Default()

	server.GET("/", helloWorldHandler)
	server.GET("/home", homeHandler)
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
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

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) // 201 Created status code

}

func helloWorldHandler(context *gin.Context) { // request, response
	context.String(200, "Hello, World!")
}

func homeHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
