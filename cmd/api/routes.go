package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)
	mux.Get("/", app.Home)
	mux.Get("/api/skill/list", app.getSkills)
	mux.Get("/api/project/list", app.getProjects)
	mux.Get("/api/project/type/{id}", app.getProjectByTypeId)

	return mux
}
