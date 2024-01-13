package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/project/get/:id", getProject)
	server.GET("/api/project/list", getProjects)
	server.POST("/api/project/create", createProject)
	server.PUT("/api/project/put/:id", updateProject)
	server.DELETE("/api/project/delete/:id", deleteProject)
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
