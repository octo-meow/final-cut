package engine

import (
	"bufio"
	"os"
	"slices"

	"github.com/gdamore/tcell/v3/color"

	dt "github.com/ereminiu/final-cut/game/engine/datastruct"
)

const (
	screenHeight = 5
	screenWidth  = 5
)

var walls = []rune{
	'║', '═', '╔', '╗', '╚', '╝', '╠', '╣', '╦', '╩', '╬',
}

var bisquit = []rune{
	'*',
}

func ParseLevel(file *os.File) [][]dt.Chunk {
	res := make([][]dt.Chunk, screenHeight)

	reader := bufio.NewReader(file)
	for i := range screenHeight {
		res[i] = make([]dt.Chunk, screenWidth)
		byteline, _, _ := reader.ReadLine()
		for j, char := range string(byteline) {
			res[i][j] = dt.Chunk{
				Char:  char,
				Color: getColor(char),
			}
		}
	}
	_ = '▛'

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
