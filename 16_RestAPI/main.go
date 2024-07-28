package main

import (
	"restapi_example/db"
	"restapi_example/routes"

	"github.com/gin-gonic/gin"
)

func main() { //com component, remoting, webservice soap based service, rest service, graphql, gRpc
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
