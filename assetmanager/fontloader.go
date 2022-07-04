package assetmanager

import (
	"log"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func LoadFont(byteSlice []byte) font.Face {
	fontFace, err := truetype.Parse(byteSlice)
	if err != nil {
		log.Fatal("Could not parse font")
	}

	font := truetype.NewFace(fontFace, &truetype.Options{
		Size:              16,
		DPI:               72,
		Hinting:           font.HintingFull,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	})
	return font
}
