package main

import (
	"database/sql"
	"github.com/caioedlobo/desafio-picpay-go/internal/application/handlers"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/api"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/eventstore"
	"github.com/caioedlobo/desafio-picpay-go/internal/infrastructure/persistence"
	"github.com/gofiber/fiber/v2"
	log "github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	db, err := sql.Open("postgres", "postgres://postgres:picpay123@localhost/picpay?sslmode=disable")
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
	defer db.Close()

	userRepo := persistence.NewPostgresUserRepository(db)
	eventStore := eventstore.NewPostgresEventStore(db)
	commandHandler := handlers.NewCommandHandler(userRepo, eventStore)
	httpHandler := api.NewHTTPHandler(commandHandler)

	app := fiber.New()
	app.Use(log.New())
	app.Post("/api/v1/users", httpHandler.CreateUser)
	app.Listen(":3000")
}
