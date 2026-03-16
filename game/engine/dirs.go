package engine

type dirtype uint8

const (
	NONE dirtype = iota
	LEFT
	RIGHT
	DOWN
	UP
)

func GetOposite(dir dirtype) dirtype {
	switch dir {
	case LEFT:
		return RIGHT

	case RIGHT:
		return LEFT

	case DOWN:
		return UP

	case UP:
		return DOWN
	}

	return NONE
}
