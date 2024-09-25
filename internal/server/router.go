package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"backend/internal/handler"
)

func NewRouter() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://axolotlmaid.com", http.StatusPermanentRedirect)
	})

	r.Get("/visitor-counter", handler.HandleGetVisitorCounter)
	r.Patch("/visitor-counter", handler.HandlePatchVisitorCounter)

	slog.Info("Starting server", slog.Any("port", os.Getenv("PORT")))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
