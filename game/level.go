package game

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/ereminiu/final-cut/game/engine"
	dt "github.com/ereminiu/final-cut/game/engine/datastruct"
)

const (
	levelPath = "internal/assets/level.txt"
)

func getInitialLevel() [][]dt.Chunk {
	file, err := os.OpenFile(levelPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("can't open file")
	}
	defer file.Close()

	return engine.ParseLevel(file)
}
