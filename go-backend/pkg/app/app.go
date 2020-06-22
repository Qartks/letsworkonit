package app

import (
	"context"

	"go-backend/pkg/log"
)

type App struct {
}

func NewApp(ctx context.Context) (App, error) {
	logger := log.GetLogger(ctx)
	logger.Info("Creating App")

	return App{}, nil
}
