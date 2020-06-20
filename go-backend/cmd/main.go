package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"go-backend/pkg/api"
	"go-backend/pkg/log"
)

func main() {
	ctx := context.Background()
	logger := log.GetLogger(ctx)

	router, err := newRouter(ctx)
	if err != nil {
		logger.WithError(err).Fatal("Router initialization failed")
	}

	err = http.ListenAndServe(":3333", router)
	if err == http.ErrServerClosed {
		logger.Print("Server closed")
	} else {
		logger.WithError(err).Fatal("Server returned error")
	}
}

func newRouter(ctx context.Context) (chi.Router, error) {
	logger := log.GetLogger(ctx)

	rtr := chi.NewRouter()

	rtr.Use(middleware.RequestID)
	rtr.Use(middleware.RealIP)
	rtr.Use(middleware.Logger)
	rtr.Use(middleware.Recoverer)

	rtr.Use(middleware.Timeout(60 * time.Second))

	rtr.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hi"))
	})

	api.MountAPIRoutes(ctx, rtr)

	return rtr, nil
}
