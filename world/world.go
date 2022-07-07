package world

var width float64
var height float64

func GetWidth() float64 {
	return width
}

func GetHeight() float64 {
	return height
}

func GetWidthInt() int {
	return int(width)
}

func GetHeightInt() int {
	return int(height)
}

func SetSize(w, h float64) {
	width = w
	height = h
}
