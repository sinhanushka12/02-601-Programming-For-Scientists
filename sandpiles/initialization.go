package main 

import (
	"math"
	"math/rand"
)

//Input: takes size, pile, and placement 
//Output: initializes the board with the provided size. also initializes the cells with the provided placement and pile. 
func InitializeBoard(size int, pile int, placement string) *Board {
    var board Board
	board.height = size
	board.length = size
	board.cells = InitializeCells(size, pile, placement)
    return &board
}

//Input: takes in size, pile, and placement 
//Output: initializes cells for the board and its other fields using different parameters provided. 
func InitializeCells (size int, pile int, placement string) [][]*Cell {
	cells := make([]([]Cell),size) 
	for r := range cells {
		cells[r] = make([]Cell, size)
		for c := range cells[r]{
			cells[r][c].numOfCoins = 0
			cells[r][c].neighbors = make([]*Cell, 0)
		}
	}
	for r:= range cells{
		for c := range cells[r]{
			if r-1 >= 0 {
				cells[r][c].neighbors = append(cells[r][c].neighbors, &cells[r-1][c])
			} 
			if r+1 < size {
				cells[r][c].neighbors = append(cells[r][c].neighbors, &cells[r+1][c])
			}
			if c-1 >= 0 {
				cells[r][c].neighbors = append(cells[r][c].neighbors, &cells[r][c-1])
			}
			if c+1 < size {
				cells[r][c].neighbors = append(cells[r][c].neighbors, &cells[r][c+1])
			}
		}
	}
	if placement == "central" {
		cells[int(math.Floor(float64(size)/2))][int(math.Floor(float64(size)/2))].numOfCoins = pile	
	} else {		
		numberOfCoinsLeft := pile		
		for i := 0; i < 99; i++ {
			x := rand.Intn(size)
			y := rand.Intn(size)
			cells[x][y].numOfCoins += rand.Intn(numberOfCoinsLeft)
			numberOfCoinsLeft -=  cells[x][y].numOfCoins
		}		
		x := rand.Intn(size)
		y := rand.Intn(size)
		cells[x][y].numOfCoins += numberOfCoinsLeft		
	}
	cells1 := make([][]*Cell , size)
	for r:= range cells1{
		cells1[r] = make([]*Cell , size)
		for c:= range cells1[r]{
			cells1[r][c] = &cells[r][c]
		}
	}
	return cells1
}