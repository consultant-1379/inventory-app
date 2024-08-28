package main

import (
	"net/http"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/config"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/instance/{instance}", handlers.Repo.Instance)
	mux.Get("/vpod/{vpod}", handlers.Repo.Vpod)
	mux.Get("/server/{server}", handlers.Repo.Server)
	mux.Get("/cluster/{cluster}", handlers.Repo.Cluster)
	mux.Get("/deployment/{deployment}", handlers.Repo.Deployment)
	mux.Get("/api/deployments", handlers.Repo.GetAllDeployments)
	mux.Route("/api/deployments/{id}", func(mux chi.Router) {
		mux.Get("/", handlers.Repo.GetOneDeployment) // GET /posts?id={id} - Read a single post by :id.
	})

	//mux.Post("/instance", handlers.Repo.PostInstance)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
