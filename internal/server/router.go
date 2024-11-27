package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"backend/internal/handler"
)

func NewRouter() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5, "application/json"))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.LimitByRealIP(32, time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://axolotlmaid.com"},
		AllowedMethods: []string{"GET", "PATCH"},
		AllowedHeaders: []string{"Content-Type"},
		MaxAge:         300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://axolotlmaid.com", http.StatusPermanentRedirect)
	})

	r.Get("/visit-counter", handler.HandleGetVisitCounter)
	r.With(httprate.LimitByRealIP(1, time.Hour)).Patch("/visit-counter", handler.HandlePatchVisitCounter)

	r.Get("/currently-playing", handler.HandleGetCurrentlyPlaying)
	r.Get("/status", handler.HandleGetStatus)

	r.Get("/ws/computer", handler.HandleComputerWebsocket)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	slog.Info("Starting server", slog.Any("port", port))
	http.ListenAndServe(":"+port, r)
}
