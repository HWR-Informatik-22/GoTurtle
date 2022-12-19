package fraktale

import (
	t "turtle/turtle"
	"image"
)

func Initialise(img *image.Paletted) {
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
