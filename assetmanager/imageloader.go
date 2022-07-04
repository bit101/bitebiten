package assetmanager

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images map[string]*ebiten.Image

func LoadImage(name string, bs []byte) {
	if images == nil {
		images = make(map[string]*ebiten.Image)
	}
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}
	images[name] = img
}

func GetImage(name string) *ebiten.Image {
	img, ok := images[name]
	if !ok {
		log.Fatalf("Could not find image: %s", name)
	}
	return img

}
