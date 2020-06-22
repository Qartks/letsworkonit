package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"go-backend/pkg/api"
	"go-backend/pkg/app"
	"go-backend/pkg/config"
	"go-backend/pkg/data"
	"go-backend/pkg/log"
)

func main() {
	ctx := context.Background()
	logger := log.GetLogger(ctx)

	cfg, err := config.FromEnv()
	if err != nil {
		logger.WithError(err).Fatal("error getting config from env")
	}

	router, err := newRouter(ctx)
	if err != nil {
		logger.WithError(err).Fatal("router initialization failed")
	}

	store, err := initDB(ctx, cfg)
	if err != nil {
		logger.WithError(err).Fatal("database initialization failed")
	}
	logger.Info(store)

	red, err := app.NewRedisClient(ctx, cfg)
	if err != nil {
		logger.WithError(err).Fatal("redis client initialization failed")
	}
	logger.Info(red)

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.WithError(err).Fatal("app initialization failed")
	}
	logger.Info(a)

	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.PORT), router)
	if err == http.ErrServerClosed {
		logger.Print("Server closed")
	} else {
		logger.WithError(err).Fatal("Server returned error")
	}
}

func initDB(ctx context.Context, cfg *config.Config) (data.Store, error) {
	logger := log.GetLogger(ctx)
	logger.Info("Initiating Db")

	dbConn, err := data.ConnectDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	db, err := data.InitializeDatabase(ctx, cfg, dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newRouter(ctx context.Context) (chi.Router, error) {
	logger := log.GetLogger(ctx)
	logger.Info("Initiating router")

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
