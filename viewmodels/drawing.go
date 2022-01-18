package viewmodels

type Line struct {
	Points      []Point
	BrushRadius int
}

type Point struct {
	X float64
	Y float64
}

type Drawing struct {
	Lines  []Line
	Width  int
	Height int
}

type Recognition struct {
	Matches []rune
}
