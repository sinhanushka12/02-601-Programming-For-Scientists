package main

import (
	"fmt"
	"gifhelper"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if numBoids < 0 {
		panic("Negative number of generations given.")
	}

	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic(err2)
	}
	if skyWidth < 0 {
		panic("Negative paramter provided for skyWidth.")
	}

	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic(err3)
	}
	if initialSpeed < 0 {
		panic("Negative paramter provided for initialSpeed.")
	}

	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic(err4)
	}
	if maxBoidSpeed < 0 {
		panic("Negative paramter provided for maxBoidSpeed.")
	}

	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic(err5)
	}
	if numGens < 0 {
		panic("Negative number of generations given.")
	}

	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic(err6)
	}
	if proximity < 0 {
		panic("Negative paramter provided for proximity.")
	}

	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic(err7)
	}
	if separationFactor < 0 {
		panic("Negative paramter provided for separationFactor.")
	}

	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic(err8)
	}
	if alignmentFactor < 0 {
		panic("Negative paramter provided for alignmentFactor.")
	}

	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9 != nil {
		panic(err9)
	}
	if cohesionFactor < 0 {
		panic("Negative paramter provided for cohesionFactor.")
	}

	timeStep, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10 != nil {
		panic(err10)
	}

	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil {
		panic(err11)
	}
	if canvasWidth < 0 {
		panic("Negative paramter provided for canvasWidth.")
	}

	imageFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12 != nil {
		panic(err12)
	}

	fmt.Println("Command line arguments read successfully.")

	initialSky := InitializeSky(numBoids, skyWidth, maxBoidSpeed, proximity, separationFactor, cohesionFactor, alignmentFactor, initialSpeed)

	fmt.Println("Simulating system.")

	timePoints := SimulateBoids(initialSky, numGens, timeStep)

	fmt.Println("Boids have been simulated!")
	fmt.Println("Ready to draw images.")

	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	gifhelper.ImagesToGIF(images, "Boids_out6")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}

func InitializeSky(numBoids int, skyWidth, initialSpeed, maxBoidSpeed, proximity, separationFactor, alignmentFactor, cohesionFactor float64) Sky {
	var initialSky Sky

	initialSky.width = skyWidth
	initialSky.maxBoidSpeed = maxBoidSpeed
	initialSky.proximity = proximity
	initialSky.separationFactor = separationFactor
	initialSky.cohesionFactor = cohesionFactor
	initialSky.alignmentFactor = alignmentFactor

	initialSky.boids = make([]Boid, numBoids)

	for i := range initialSky.boids {
		c := 2 * math.Pi * rand.Float64()
		//initialSky.boids[i].velocity.x = ((rand.Float64() * 2) - 1) * initialSpeed
		initialSky.boids[i].velocity.x = math.Sin(c) * initialSpeed
		initialSky.boids[i].velocity.y = math.Sqrt(initialSpeed*initialSpeed - initialSky.boids[i].velocity.x*initialSky.boids[i].velocity.x)
		//initialSky.boids[i].velocity.y = math.Cos(c) * initialSpeed
		initialSky.boids[i].position.x = rand.Float64() * skyWidth
		initialSky.boids[i].position.y = rand.Float64() * skyWidth
		initialSky.boids[i].acceleration.x = 0.0
		initialSky.boids[i].acceleration.y = 0.0
	}
	return initialSky
}

//./boids 200 2000 1.0 2.0 8000 200 1.5 1.0 0.02 1.0 2000 20
