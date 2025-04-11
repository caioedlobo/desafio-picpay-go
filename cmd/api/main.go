package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

type application struct {
	logger *zerolog.Logger
}

func main() {
	logger := zerolog.New(os.Stdout)
	app := &application{
		logger: &logger,
	}
	app.logger.Info().Msg("Hello World")
	err := app.serve()
	if err != nil {
		fmt.Println(err)
	}

}
