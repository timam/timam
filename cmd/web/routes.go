package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/timam/timam/pkg/config"
	"github.com/timam/timam/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig)  http.Handler{
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
