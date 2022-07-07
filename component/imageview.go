package component

import (
	"github.com/bit101/bitebiten/assetmanager"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageView struct {
	Actor
	regX, regY float64
	image      *ebiten.Image
}

func NewImageView(name string) *ImageView {
	img := &ImageView{}
	img.SetImage(name)
	return img
}

func (i *ImageView) SetImage(name string) {
	i.image = assetmanager.GetImage(name)
}

func (i *ImageView) GetImage() *ebiten.Image {
	return i.image
}

func (i *ImageView) SetRegistration(x, y float64) {
	i.regX = -x
	i.regY = -y
}

func (i *ImageView) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(i.regX, i.regY)
	opts.GeoM.Rotate(i.Rotation)
	opts.GeoM.Translate(i.X, i.Y)

	if i.image != nil {
		screen.DrawImage(i.image, opts)
	}
}
