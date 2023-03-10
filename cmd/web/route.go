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
	mux.Use(MyMiddleWare)
	mux.Use(Csrf)
	mux.Use(SessionLoad)
	fileServer := http.FileServer(http.Dir("./statics/"))
	mux.Handle("/statics/*", http.StripPrefix("/statics/", fileServer))
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/majors", handlers.Repo.Major)
	mux.Get("/SearchRooms", handlers.Repo.SearchRooms)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation/{id}", handlers.Repo.ReservationPage)
	mux.Post("/make-reservation", handlers.Repo.HandleReservation)
	mux.Post("/search", handlers.Repo.AvailabilityJson)
	mux.Post("/searchAvailability", handlers.Repo.SearchAvailability)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	mux.Get("/bookRoom", handlers.Repo.BookRoom)
	mux.Get("/RenderResPage", handlers.Repo.RenderReservationPage)
	return mux
}
