package main

import (
	fr "turtle/fraktale"
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

	//write your Turtle commands here
	
	//fractals drawn with the turtle
	fr.Initialise(img)

	images = append(images, img)
	delay = append(delay, 500)

	println(len(images))
	println(len(delay))

	// Programmcode für das Zeichnen
	anim := gif.GIF{Delay: delay, Image: images}
	gif.EncodeAll(f, &anim)

	f.Close()
}
