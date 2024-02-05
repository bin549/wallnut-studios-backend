package main

import (
	"fmt"
	"net/http"
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
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, skills)
}
