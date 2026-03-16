package gameprovider

import (
	"os"

	"github.com/gdamore/tcell/v3"
	"github.com/rs/zerolog/log"
)

const (
	levelPath = "internal/assets/level.txt"
)

type GameProvider struct {
	screen    tcell.Screen
	levelFile *os.File
}

func New() *GameProvider {
	return &GameProvider{}
}

func (gp *GameProvider) GetScreen() tcell.Screen {
	if gp.screen == nil {
		screen, err := tcell.NewScreen()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get screen")
		}

		if err := screen.Init(); err != nil {
			log.Fatal().Err(err).Msg("failed to init screen")
		}

		gp.screen = screen
	}

	return gp.screen
}

func (gp *GameProvider) GetLevelFile() *os.File {
	if gp.levelFile == nil {
		file, err := os.OpenFile(levelPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal().Err(err).Msg("can't open file")
		}

		gp.levelFile = file
	}

	return gp.levelFile
}

func (gp *GameProvider) Close() {
	gp.levelFile.Close()

}
