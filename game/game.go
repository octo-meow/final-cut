package game

import (
	"time"

	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
	"github.com/rs/zerolog/log"
)

type Game struct {
	screen tcell.Screen
}

func New(screen tcell.Screen) *Game {
	return &Game{
		screen: screen,
	}
}

func (g *Game) Run() {
	defer func() {
		g.screen.Fini()

		if err := recover(); err != nil {
			log.Fatal().Any("err", err).Msg("caught panic")
		}
	}()

	g.screen.SetStyle(tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset))

	level := getInitialLevel()
	for x, row := range level {
		for y, chunk := range row {
			g.screen.PutStrStyled(
				y, x,
				string(chunk.Char),
				tcell.StyleDefault.Background(color.Reset).Foreground(chunk.Color),
			)
		}
	}
	g.screen.Show()

	time.Sleep(2 * time.Second)
}

func (g *Game) update() {
}

func (g *Game) render() {
}

func (g *Game) readUserInput() {
}
