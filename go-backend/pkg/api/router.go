package api

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"go-backend/pkg/log"
)

const (
	RouteV1 = "/api/v1"
)

func MountAPIRoutes(ctx context.Context, rtr *chi.Mux) {
	logger := log.GetLogger(ctx)
	logger.Info("loading api routes")

	rtr.Get("/healthCheck", HandleHealthCheck)

	rtr.Route(RouteV1, func(r chi.Router) {
		r.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "PUT", "DELETE", "OPTION", "POST"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			MaxAge:           300,
		}).Handler)

	})
}
