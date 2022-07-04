package component

import (
	"github.com/bit101/bitebiten/assetmanager"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageView struct {
	Actor
	image *ebiten.Image
}

func NewImageView(name string, x, y float64) *ImageView {
	img := &ImageView{}
	img.SetPos(-x, -y) // sets the registration point by moving origin in opposite direction
	img.SetImage(name)
	return img
}

func (c *ImageView) SetImage(name string) {
	c.image = assetmanager.GetImage(name)
}

func (c *ImageView) GetImage() *ebiten.Image {
	return c.image
}

func (c *ImageView) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Rotate(c.GetRotation())
	opts.GeoM.Translate(c.x, c.y)

	if c.image != nil {
		screen.DrawImage(c.image, opts)
	}
}
