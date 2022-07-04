package utils

import "image/color"

func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func Red() color.RGBA {
	return color.RGBA{255, 0, 0, 255}
}
func Green() color.RGBA {
	return color.RGBA{0, 255, 0, 255}
}
func Blue() color.RGBA {
	return color.RGBA{0, 0, 255, 255}
}
func Yellow() color.RGBA {
	return color.RGBA{255, 255, 0, 255}
}
func Cyan() color.RGBA {
	return color.RGBA{0, 255, 255, 255}
}
func Purple() color.RGBA {
	return color.RGBA{255, 0, 255, 255}
}
