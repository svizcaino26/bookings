package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/svizcaino26/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Errorf("Expected type *chi.Mux. Received: %T", v)
	}
}
