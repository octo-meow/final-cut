package datastruct

import "github.com/gdamore/tcell/v3/color"

type Chunk struct {
	Char  rune
	Color color.Color
}

type Frame [][]Chunk
