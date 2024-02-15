package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"wallnut-studios.com/rest-api/internal/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Go",
		Version: "1.0.0",
	}
	fmt.Println(payload)
}

func (app *application) getSkills(w http.ResponseWriter, r *http.Request) {
	skills, err := app.DB.GetSkills()
	skill_tags, err := app.DB.GetSkillTags()
	skill_tag_map := map[int]string{}
	for _, skill_tag := range skill_tags {
		skill_tag_map[skill_tag.ID] = skill_tag.Name
	}
	var skill_dtos []models.SkillDto
	for _, skill := range skills {
		var skill_dto models.SkillDto
		skill_dto.Id = skill.ID
		skill_dto.Name = skill.Name
		skill_dto.Tag = skill_tag_map[skill.TagId]
		skill_dtos = append(skill_dtos, skill_dto)
	}
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, skill_dtos)
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := app.DB.GetProjects()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, projects)
}

func (app *application) getProjectByTypeId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	projects, err := app.DB.GetProjects(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, projects)
}
