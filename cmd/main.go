package main

import (
	"context"
	"database/sql"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/handlers"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/api"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/eventstore"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/persistence"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	connString := "postgres://postgres:picpay123@localhost/picpay?sslmode=disable"
	db, err := sql.Open("postgres", connString)
	dbpool, err := pgxpool.New(context.Background(), connString)
	defer dbpool.Close()

	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer db.Close()

	userRepo := persistence.NewPostgresUserRepository(db)
	eventStore := eventstore.NewPostgresEventStore(dbpool)
	commandHandler := handlers.NewCommandHandler(userRepo, eventStore)
	queryHandler := handlers.NewQueryHandler(userRepo)
	validate := validator.New(validator.WithRequiredStructEnabled())
	httpHandler := api.NewHTTPHandler(commandHandler, queryHandler, validate)

	app := fiber.New()
	app.Use(log.New())
	app.Post("/api/v1/users", httpHandler.CreateUser)
	app.Listen(":3000")
}
