package engine

import (
	"bufio"
	"os"
	"slices"

	"github.com/gdamore/tcell/v3/color"
	"github.com/rs/zerolog/log"

	dt "github.com/ereminiu/final-cut/internal/game/engine/datastruct"
)

const (
	screenHeight = 10
	screenWidth  = 15
)

var walls = []rune{
	'║', '═', '╔', '╗', '╚', '╝', '╠', '╣', '╦', '╩', '╬',
}

var bisquit = []rune{
	'*',
}

func ParseLevel(file *os.File) dt.Frame {
	res := make([][]dt.Chunk, screenHeight)

	reader := bufio.NewReader(file)
	for i := range screenHeight {
		res[i] = make([]dt.Chunk, screenWidth)
		byteline, _, err := reader.ReadLine()

		if err != nil {
			log.Debug().Err(err).Msg("error during readline")
		}

		strline := string(byteline)

		log.Debug().
			Str("byteline", string(byteline)).
			Int("run len", len([]rune(strline))).
			Int("str len", len(strline)).
			Msg("here")

		for j, char := range []rune(strline) {
			res[i][j] = dt.Chunk{
				Char:  char,
				Color: getColor(char),
			}
		}
	}

	return res
}

func getColor(char rune) color.Color {
	if slices.Contains(walls, char) {
		return color.Blue
	}

	if slices.Contains(bisquit, char) {
		return color.White
	}

	return color.Pink
}
