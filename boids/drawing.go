package main

import (
	"canvas"
	"image"
)

//place your drawing code here.

// let's place our drawing functions here.

//AnimateSystem takes a slice of Sky objects along with a Skywidth
//parameter and a frequency parameter. It generates a slice of images corresponding to drawing each Sky whose index is divisible by the frequency parameter
//on a skyWidth x skyWidth canvas

func AnimateSystem(timePoints []Sky, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0)

	for i := range timePoints {
		if i%imageFrequency == 0 { //only draw if current index of Sky
			//is divisible by some parameter frequency
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Sky
//object's boids on a square canvas that is skyWidth pixels x skyWidth pixels
func DrawToCanvas(u Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the boids and draw them.
	for _, b := range u.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		// we want to return an image!
		// r := (1 / u.width) * float64(canvasWidth)
		c.Circle(cx, cy, 5.0)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
