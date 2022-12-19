package main

import (
	t "fraktale/turtle"
	"image"
	"image/color"
	"image/gif"
	"os"
)

const CANVAS_WIDTH = 2000
const CANVAS_HEIGHT = 2000

func main() {
	f, err := os.Create("my-image.gif")
	if err != nil {
		panic(err)
	}

	var images []*image.Paletted
	var delay []int

	var palette = []color.Color{
		color.RGBA{0xff, 0xff, 0xff, 0xff}, //Hintergrundfarbe des Bilds
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0xff, 0xff},
	}
	rect := image.Rect(0, 0, CANVAS_WIDTH, CANVAS_HEIGHT)
	img := image.NewPaletted(rect, palette)

	t.Color("blue")
	t.Transform(img, 200, 200)
	
	t.Transform(img, 200, 200)
	Koch(img, 60, 3, 729, 4)

	t.Transform(img, 1200, 200)
	Hilbert(img, 90, 2, 600, 8)

	t.Transform(img, 200, 1200)
	LevyC(img, 45, 1.41421356, 600, 16)

	t.Transform(img, 1200, 1200)
	Drachenkurve(img, 45, 1.41421356, 600, 14)

	images = append(images, img)
	delay = append(delay, 500)

	println(len(images))
	println(len(delay))

	// Programmcode für das Zeichnen
	anim := gif.GIF{Delay: delay, Image: images}
	gif.EncodeAll(f, &anim)

	f.Close()
}

func Koch(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	KochFraktal(img, angle, ratio, distance, iterations)
	t.Turn(60)
	t.Turn(60)
	KochFraktal(img, angle, ratio, distance, iterations)
	t.Turn(60)
	t.Turn(60)
	KochFraktal(img, angle, ratio, distance, iterations)
}

func KochFraktal(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}

	KochFraktal(img, angle, ratio, int(float64(distance)/ratio), iterations-1)
	t.Turn(-angle)
	KochFraktal(img, angle, ratio, int(float64(distance)/ratio), iterations-1)
	t.Turn(angle)
	t.Turn(angle)
	KochFraktal(img, angle, ratio, int(float64(distance)/ratio), iterations-1)
	t.Turn(-angle)
	KochFraktal(img, angle, ratio, int(float64(distance)/ratio), iterations-1)
}

func Hilbert(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	A(img, angle, ratio, distance, iterations)
}

func A(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}

	D(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Up()
	A(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Right()
	A(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Down()
	B(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
}

func B(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}

	C(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Left()
	B(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Down()
	B(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Right()
	A(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
}

func C(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}

	B(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Down()
	C(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Left()
	C(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Up()
	D(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
}

func D(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}

	A(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Right()
	D(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Up()
	D(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
	t.Left()
	C(img, angle, ratio, int(float64(distance)*1/ratio), iterations-1)
}

func LevyC(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	LevyF(img, 45, 1.41421356, distance, iterations)
}

func Drachenkurve(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	LevyL(img, 45, 1.41421356, distance, iterations)
}

func LevyF(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}
	t.Turn(-angle)
	LevyC(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(angle)
	t.Turn(angle)
	LevyC(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(-angle)
}

func LevyL(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}
	t.Turn(-angle)
	LevyR(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(angle)
	t.Turn(angle)
	LevyL(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(-angle)
}

func LevyR(img *image.Paletted, angle, ratio float64, distance, iterations int) {
	if iterations == 0 {
		t.Move(img, distance)
		return
	}
	t.Turn(angle)
	LevyR(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(-angle)
	t.Turn(-angle)
	LevyL(img, 45, 1.41421356, int(float64(distance)*1/ratio), iterations-1)
	t.Turn(angle)
}
