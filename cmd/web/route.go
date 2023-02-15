package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vitaLemoTea/secondstepweb/internal/handlers"
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
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Major)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/book", handlers.Repo.Book)
	mux.Post("/book", handlers.Repo.SearchAvailability)
	mux.Post("/search", handlers.Repo.AvailabilityJson)
	return mux
}
