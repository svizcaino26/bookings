package main

import (
	"net/http"

	"github.com/svizcaino26/bookings/internal/config"
	"github.com/svizcaino26/bookings/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// [GET ROUTES]
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/colonels-suite", handlers.Repo.Colonels)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/htmx", handlers.Repo.Htmx)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// [POST ROUTES]
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
