package component

import (
	"image"

	"github.com/bit101/bitebiten/assetmanager"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimView struct {
	Actor
	regX, regY  float64
	image       *ebiten.Image
	tick        int
	frame       int
	frameWidth  int
	frameHeight int
	delayFrames int
	frameCount  int
}

func NewAnimView(name string, fw, fh, fc, delay int) *AnimView {
	img := &AnimView{
		frameWidth:  fw,
		frameHeight: fh,
		frameCount:  fc,
		delayFrames: delay,
	}
	img.SetImage(name)
	return img
}

func (i *AnimView) SetImage(name string) {
	i.image = assetmanager.GetImage(name)
}

func (i *AnimView) GetImage() *ebiten.Image {
	return i.image
}

func (i *AnimView) SetRegistration(x, y float64) {
	i.regX = -x
	i.regY = -y
}

func (i *AnimView) Update() {
	i.tick++
	if i.tick%i.delayFrames == 0 {
		i.frame++
	}
	i.frame %= i.frameCount
}

func (i *AnimView) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(i.regX, i.regY)
	opts.GeoM.Rotate(i.Rotation)
	opts.GeoM.Translate(i.X, i.Y)

	cols := i.image.Bounds().Dx() / i.frameWidth
	col := i.frame % cols
	row := i.frame / cols

	subImage := i.image.SubImage(image.Rect(col*i.frameWidth, row*i.frameHeight, (col+1)*i.frameWidth, (row+1)*i.frameHeight)).(*ebiten.Image)

	if i.image != nil {
		screen.DrawImage(subImage, opts)
	}
}
