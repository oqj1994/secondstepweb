package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitaLemoTea/secondstepweb/pkg/handlers"
)

func NewRouter() chi.Router {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(mymiddleware)
	mux.Use(csrf)
	mux.Use(sessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
