package component

import (
	"math"

	"github.com/bit101/bitebiten/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// implements IActor
type Actor struct {
	x, y, rotation float64
}

func NewActor() *Actor {
	return &Actor{}
}

func (c *Actor) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Rotate(c.rotation)
	opts.GeoM.Translate(c.x, c.y)
}

func (c *Actor) MoveBy(x, y float64) {
	c.x += x
	c.y += y
}

func (c *Actor) SetPos(x, y float64) {
	c.x, c.y = x, y
}

func (c *Actor) GetPos() *utils.Point {
	return utils.NewPoint(c.x, c.y)
}

func (c *Actor) RotateBy(r float64) {
	c.rotation += r
}

func (c *Actor) SetRotation(r float64) {
	c.rotation = r
}

func (c *Actor) GetRotation() float64 {
	return c.rotation
}

func (c *Actor) Distance(x, y float64) float64 {
	dx := x - c.x
	dy := y - c.y
	return math.Sqrt(dx*dx + dy*dy)
}
