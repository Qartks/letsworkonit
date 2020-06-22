package data

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"go-backend/pkg/config"
	"go-backend/pkg/log"
)

func ConnectDB(ctx context.Context, cfg *config.Config) (db *sqlx.DB, err error) {
	logger := log.GetLogger(ctx)
	logger.Info("Connecting to DB")

	db, err = sqlx.Connect("postgres", "host=postgres user=postgres dbname=letsworkonit port=5432 password=postgres sslmode=disable")
	if err != nil {
		logger.WithError(err).Fatal("connecting to database failed")
	}
	return db, nil
}

func InitializeDatabase(ctx context.Context, cfg *config.Config, conn *sqlx.DB) (db Store, err error) {
	return db, nil
}
