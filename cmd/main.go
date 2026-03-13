package main

import (
	"log"
	"os"

	"github.com/ereminiu/final-cut/game"
	"github.com/ereminiu/final-cut/xlog"
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

// Game dependencies:
// Screen

func main() {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("can't init logger: %v", err)
	}
	xlog.InitLogger(file)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalln(err)
	}

	if err = screen.Init(); err != nil {
		log.Fatalln(err)
	}

	screen.SetStyle(tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset))

	game := game.New(screen)

	game.Run()
}
