package turtle

import (
	"image"
	"math"
)

type Point image.Point

var orientation float64 = 0.0

var center = image.Point{0, 0}

var color uint8 = 1

func Transform(img *image.Paletted, x1, y1 int) {
	center.X = x1
	center.Y = y1
}

func Color(c string) {
	switch c {
	case "black":
		color = 1
	case "white":
		color = 0
	case "red":
		color = 2
	case "green":
		color = 3
	case "blue":
		color = 4
	case "yellow":
		color = 5
	case "magenta":
		color = 6
	case "cyan":
		color = 7
	default:
		color = 1
	}
}

func Up() {
	orientation = 90
}
func Down() {
	orientation = -90
}
func Left() {
	orientation = 180
}
func Right() {
	orientation = 0
}

func Move(img *image.Paletted, distance int) {
	x2 := center.X + int(math.Cos(orientation*(2*math.Pi/360))*float64(distance))
	y2 := center.Y + int(math.Sin(orientation*(2*math.Pi/360))*float64(distance))
	DrawLine(img, center.X, center.Y, x2, y2, color)
	Transform(img, x2, y2)
}

func Turn(angle float64) {
	orientation += angle
}

func DrawLine(img *image.Paletted, x1, y1, x2, y2 int, col uint8) {
	Bresenham(img, x1, y1, x2, y2, col)
}

// Generalized with integer
func Bresenham(img *image.Paletted, x1, y1, x2, y2 int, col uint8) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.SetColorIndex(x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.SetColorIndex(x1, y1, col)
			x1++
		}
		img.SetColorIndex(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.SetColorIndex(x1, y1, col)
			y1++
		}
		img.SetColorIndex(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				y1--
			}
		}
		img.SetColorIndex(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.SetColorIndex(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.SetColorIndex(x2, y2, col)
	}
}
