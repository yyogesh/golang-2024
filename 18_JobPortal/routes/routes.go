package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signUpHandler)
	server.POST("/login", loginHandler)

	server.GET("/jobs", getJobsHandler)
	server.GET("/job/:id", getJobByIdHandler)
	server.POST("/job", createJobHandler)
	server.PUT("/job/:id", updateJobHandler)
	server.DELETE("/job/:id", deleteJobHandler)
}
