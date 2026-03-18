package game

import (
	"context"
	"math/rand"
	"sync"
	"time"

	dt "github.com/ereminiu/final-cut/internal/game/engine/datastruct"
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
	g.screen.SetStyle(tcell.StyleDefault.Background(color.Reset).Foreground(color.Reset))

	keys := make(chan string, 1)
	frames := make(chan dt.Frame, 1)

	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		defer wg.Done()
		g.update(ctx, keys, frames)
	}()

	go func() {
		defer wg.Done()
		g.render(frames)
	}()

	wg.Wait()
}

func (g *Game) update(ctx context.Context, keys <-chan string, frames chan<- dt.Frame) {
	ticker := time.NewTicker(250 * time.Millisecond)
	currentFrame := getInitialLevel()

	for {
		select {
		case <-ctx.Done():
			close(frames)
			return

		case <-ticker.C:
			nextFrame := currentFrame
			i, j := rand.Int()%10, rand.Int()%15

			nextFrame[i][j] = dt.Chunk{
				Char:  '$',
				Color: color.AliceBlue,
			}

			frames <- nextFrame
			currentFrame = nextFrame
		}
	}
}

func (g *Game) render(frames <-chan dt.Frame) {
	defer func() {
		g.screen.Fini()

		if err := recover(); err != nil {
			log.Fatal().Any("err", err).Msg("caught panic")
		}
	}()

	for frame := range frames {
		g.drawFrame(frame)
	}
}

func (g *Game) readUserInput() {
}

func (g *Game) drawFrame(frame dt.Frame) {
	g.screen.Clear()

	for x, row := range frame {
		for y, chunk := range row {
			g.screen.PutStrStyled(
				y, x,
				string(chunk.Char),
				tcell.StyleDefault.Background(color.Reset).Foreground(chunk.Color),
			)
		}
	}

	g.screen.Show()
}
