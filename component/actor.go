package component

import (
	"math"

	"github.com/bit101/bitebiten/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Actor struct {
	X, Y, Rotation float64
}

func NewActor() *Actor {
	return &Actor{}
}

func (c *Actor) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Rotate(c.Rotation)
	opts.GeoM.Translate(c.X, c.Y)
}

func (c *Actor) MoveBy(x, y float64) {
	c.X += x
	c.Y += y
}

func (c *Actor) SetPos(x, y float64) {
	c.X, c.Y = x, y
}

func (c *Actor) GetPos() *utils.Point {
	return utils.NewPoint(c.X, c.Y)
}

func (c *Actor) Distance(x, y float64) float64 {
	dx := x - c.X
	dy := y - c.Y
	return math.Sqrt(dx*dx + dy*dy)
}
