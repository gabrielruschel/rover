package rover

const (
	North = 'N'
	South = 'S'
	East  = 'E'
	West  = 'W'

	Left  = 'L'
	Right = 'R'
)

var (
	Orientations   = [4]rune{East, South, West, North}
	OrientationIdx = map[rune]int{
		East:  0,
		South: 1,
		West:  2,
		North: 3,
	}
)
