package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ereminiu/final-cut/internal/game"
	gameprovider "github.com/ereminiu/final-cut/internal/game_provider"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("can't open file")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        file,
		TimeFormat: "",
	})

	gameProvider := gameprovider.New()

	game := game.New(gameProvider.GetScreen())
	game.Run()

	gameProvider.Close()
}
