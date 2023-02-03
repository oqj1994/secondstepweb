package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitaLemoTea/secondstepweb/pkg/handlers"
	"net/http"
)

func NewRouter() chi.Router {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(mymiddleware)
	mux.Use(csrf)
	mux.Use(sessionLoad)
	fileServer := http.FileServer(http.Dir("./statics/"))
	mux.Handle("/statics/*", http.StripPrefix("/statics/", fileServer))
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
