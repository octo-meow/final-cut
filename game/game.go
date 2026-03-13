package game

import (
	"time"

	"github.com/ereminiu/final-cut/xlog"
	"github.com/gdamore/tcell/v3"
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

		err := recover()
		if err != nil {
			xlog.Debug("caught panic: %v\n", err)
		}
	}()

	g.screen.PutStr(10, 10, "evelina")
	g.screen.Show()
	time.Sleep(2 * time.Second)
	g.screen.Clear()
}
