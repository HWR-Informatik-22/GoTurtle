package turtle

import (
"image"
"math"
)

type Point image.Point

// orientation stores the current orientation of the turtle (in degrees)
var orientation float64 = 0.0

// center stores the current location of the turtle as a Point struct
var center = image.Point{0, 0}

// color stores the color of the turtle as a uint8 value
var color uint8 = 1

// Transform sets the location of the turtle to the coordinates (x1, y1)
func Transform(img *image.Paletted, x1, y1 int) {
	center.X = x1
	center.Y = y1
}

// Color sets the color of the turtle to one of the given colors (black, white, red, green, blue, yellow, magenta, cyan).
// If an invalid color is passed, the color is set to black.
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

// Up sets the orientation of the turtle to 90 degrees upwards
func Up() {
	orientation = 90
}

// Down sets the orientation of the turtle to 90 degrees downwards
func Down() {
	orientation = -90
}

// Left sets the orientation of the turtle to 180 degrees to the left
func Left() {
	orientation = 180
}

// Right sets the orientation of the turtle to 0 degrees to the right
func Right() {
	orientation = 0
}

// Move moves the turtle a certain distance in its current orientation and draws a line from its current location to its new location.
// The new location is then stored in "center".
func Move(img *image.Paletted, distance int) {
	x2 := center.X + int(math.Cos(orientation*(2*math.Pi/360))*float64(distance))
	y2 := center.Y + int(math.Sin(orientation*(2*math.Pi/360))*float64(distance))
	DrawLine(img, center.X, center.Y, x2, y2, color)
	Transform(img, x2, y2)
}

// Turn rotates the turtle a certain angle (in degrees) clockwise
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
