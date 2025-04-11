package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type application struct {
	logger *zerolog.Logger
	config config
}

type config struct {
	db struct {
		dsn string
	}
}

func main() {
	logger := zerolog.New(os.Stdout)
	var cfg config
	cfg.db.dsn = "postgres://postgres:picpay123@localhost/picpay?sslmode=disable"

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer db.Close()

	logger.Info().Msg("database connection established")

	app := &application{
		logger: &logger,
		config: cfg,
	}

	err = app.serve()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
