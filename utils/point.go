package utils

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func CursorPositionPoint() *Point {
	x, y := ebiten.CursorPosition()
	return &Point{float64(x), float64(y)}
}

func (p *Point) Xint() int {
	return int(p.X)
}

func (p *Point) Yint() int {
	return int(p.Y)
}

func (p *Point) Distance(p2 *Point) float64 {
	return math.Hypot(p2.X-p.X, p2.Y-p.Y)
}
