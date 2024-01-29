package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wallnut-studios.com/rest-api/models"
)

func getProjects(context *gin.Context) {
	projects, err := models.GetAllProjects()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	fmt.Print(projects)
	context.JSON(http.StatusOK, projects)
}

func getProject(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	project, err := models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, project)
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

func updateProject(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := context.GetInt64("userId")
	project, err := models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}
	if project.ProfileId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}
	err = project.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}

func deleteProject(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := context.GetInt64("userId")
	project, err := models.GetProjectById(projectId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}
	if project.ProfileId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}
	err = project.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
