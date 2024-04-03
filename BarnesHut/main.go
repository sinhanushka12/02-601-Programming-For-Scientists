package main

import (
	"os"
	"gifhelper"
	"fmt"
)

func main() {
	
		whichDataset := os.Args[1]
		numGens := 10 //500000
		time := 2e14
		theta := 0.5
		width := 1.0e23
		scalingFactor := 1e11

		var initialUniverse *Universe

		if whichDataset == "jupiter" {
			time = 1.0
			scalingFactor = 10.0
			numGens = 10
			initialUniverse = InitializeJupiter()
			//timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		} else if whichDataset == "galaxy" {
			g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
			galaxies := []Galaxy{g0}
			initialUniverse = InitializeUniverse(galaxies, width)
			//timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		} else {
			g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
			g1 := InitializeGalaxy(500, 4e21, 3e22, 7e22)
			galaxies := []Galaxy{g0, g1}
			initialUniverse = InitializeUniverse(galaxies, width)
			
		}
		timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		
		fmt.Println("Simulation run. Now drawing images.")
		canvasWidth := 1000
		frequency := 1000
		// a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)

		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "galaxy")
		fmt.Println("GIF drawn.")
	
}
