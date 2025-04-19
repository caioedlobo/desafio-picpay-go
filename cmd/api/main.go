package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/caioedlobo/desafio-picpay-go/cmd/internal/data"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	version string
)

type application struct {
	logger *zerolog.Logger
	config config
	models data.Models
}

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	limiter bool
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 3000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "", "PostgreSQL DSN")
	flag.BoolVar(&cfg.limiter, "limiter-enable", true, "Enable rate limiter")
	flag.Parse()

	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer db.Close()

	logger.Info().Msg("database connection established")

	app := &application{
		logger: &logger,
		config: cfg,
		models: data.NewModels(db),
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
