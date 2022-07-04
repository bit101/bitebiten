package utils

import (
	"math"
)

func Clamp(val, min, max float64) float64 {
	return math.Min(math.Max(val, min), max)
}

// https://stackoverflow.com/questions/849211/shortest-distance-between-a-point-and-a-line-segment
func DistanceSegmentToPoint(x0, y0, x1, y1, x2, y2 float64) float64 {
	a := x0 - x1
	b := y0 - y1
	c := x2 - x1
	d := y2 - y1

	dot := a*c + b*d
	len_sq := c*c + d*d
	param := -1.0
	if len_sq != 0 {
		//in case of 0 length line
		param = dot / len_sq
	}

	var xx, yy float64
	if param < 0 {
		xx = x1
		yy = y1
	} else if param > 1 {
		xx = x2
		yy = y2
	} else {
		xx = x1 + param*c
		yy = y1 + param*d
	}

	return math.Hypot(x0-xx, y0-yy)
}

func DistanceLineToPoint(p0, p1, p2 *Point) float64 {
	// p0 is point
	// p1 -> p2 is line
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	numerator := math.Abs(dx*(p1.Y-p0.Y) - (p1.X-p0.X)*dy)
	denominator := math.Sqrt(dx*dx + dy*dy)
	return numerator / denominator
}

func SegmentIntersect(p0, p1, p2, p3 *Point) (*Point, bool) {
	a1 := p1.Y - p0.Y
	b1 := p0.X - p1.X
	c1 := a1*p0.X + b1*p0.Y
	a2 := p3.Y - p2.Y
	b2 := p2.X - p3.X
	c2 := a2*p2.X + b2*p2.Y
	denominator := a1*b2 - a2*b1

	if denominator == 0.0 {
		return NewPoint(0, 0), false
	}
	intersectX := (b2*c1 - b1*c2) / denominator
	intersectY := (a1*c2 - a2*c1) / denominator
	rx0 := (intersectX - p0.X) / (p1.X - p0.X)
	ry0 := (intersectY - p0.Y) / (p1.Y - p0.Y)
	rx1 := (intersectX - p2.X) / (p3.X - p2.X)
	ry1 := (intersectY - p2.Y) / (p3.Y - p2.Y)

	if ((rx0 >= 0.0 && rx0 <= 1.0) || (ry0 >= 0.0 && ry0 <= 1.0)) &&
		((rx1 >= 0.0 && rx1 <= 1.0) || (ry1 >= 0.0 && ry1 <= 1.0)) {
		return NewPoint(intersectX, intersectY), true
	}
	return NewPoint(0, 0), false
}

func CircleLineIntersection(p0, p1, c *Point, r float64) (bool, *Point, *Point) {

	// https://mathworld.wolfram.com/Circle-LineIntersection.html
	x0 := p0.X - c.X
	y0 := p0.Y - c.Y
	x1 := p1.X - c.X
	y1 := p1.Y - c.Y

	dx := x1 - x0
	dy := y1 - y0
	dr := dx*dx + dy*dy
	r2 := r * r
	dot := x0*y1 - x1*y0
	dot2 := dot * dot
	sign := 1.0
	if dy < 0 {
		sign = -1.0
	}

	m := math.Sqrt(r2*dr - dot2)
	n := dot * dy
	p := -dot * dx
	q := sign * dx * m
	s := math.Abs(dy) * m

	xA := c.X + (n+q)/dr
	yA := c.Y + (p+s)/dr

	xB := c.X + (n-q)/dr
	yB := c.Y + (p-s)/dr

	hit := r2*dr-dot2 >= 0
	return hit, NewPoint(xA, yA), NewPoint(xB, yB)
}
