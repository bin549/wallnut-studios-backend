package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"wallnut-studios.com/rest-api/db"
	"wallnut-studios.com/rest-api/models"
)

func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/api/project/list", getProjects)
	server.POST("/api/project/create", createProject)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Run(":8000")
}

func getProjects(context *gin.Context) {
	projects, err := models.GetAllProjects()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	fmt.Print(projects)
	context.JSON(http.StatusOK, projects)
}

func createProject(context *gin.Context) {
	var project models.Project
	err := context.ShouldBindJSON(project)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cound no parse"})
	}
	project.ID = 1
	project.Name = "1"
	project.ProfileId = 1
	project.TypeId = 1
	err = project.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create project. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Project created", "project": project})
}
