package main

import (
	"os"
	"fmt"
	"strconv"
	"time"
	"log"
)

func main () {

	//start := time.Now()

	size, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if size < 0 {
		panic("Negative number of size given.")
	}

	pile, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	if pile < 0 {
		panic("Negative number of pile given.")
	}

	placement := os.Args[3]
	if (placement != "random") && (placement != "central") {
		panic("Incorrect placement provided")
	}
	
	//Initializing Board
	cellWidth := 1
	initialBoard := InitializeBoard(size, pile, placement)
	
	//Parallel
	
	fmt.Println("Simulation running using a parrallel program")
	start := time.Now()
	newBoard := SandpileMultiprocs(initialBoard, true)
	elapsed := time.Since(start)
	log.Printf("Simulating Sandpiles in parallel took %s", elapsed)
	


	//Serial
	/*
	fmt.Println("Simulation running using a serial program")
	start2 := time.Now()
	newBoard := SandpileMultiprocs(initialBoard, false)
	elapsed2 := time.Since(start2)
	log.Printf("Simulating Sandpiles serially took %s", elapsed2)
	*/
	
	Board := DrawBoard(newBoard, cellWidth)
	name := "Board" + strconv.Itoa(size) + strconv.Itoa(pile) + placement + ".png" 
	Board.SaveToPNG(name)
	

	// elapsed := time.Since(start)
	// log.Printf("Time taken %s" , elapsed)
}
