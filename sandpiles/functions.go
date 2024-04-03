package main 

import(
	"math"
	"runtime"
)

//Input: initialBoard and isParallel
//Output: an updated board
func SandpileMultiprocs(initialBoard *Board, isParallel bool) *Board {
	
	for MaxCoinsInCell(initialBoard) >= 4 {
		initialBoard = UpdateBoard(initialBoard, isParallel)
	}
	return initialBoard
}

//Input: an initialBoard and isParallel
//Output: applies CopyBoard and PostToppling to return a newBoard
func UpdateBoard(initialBoard *Board, isParallel bool) *Board {	
	newBoard := CopyBoard(initialBoard)
	newBoard = PostToppling(newBoard, isParallel)

	return newBoard
	
}

//Input: takes a board
//Output: copies a board and its other fields
func CopyBoard(b *Board) *Board {
	var newBoard Board

 	newBoard.height = b.height
	newBoard.length = b.length
 	newBoard.cells = make([][]*Cell, b.height)
 	for r := range newBoard.cells {
 		newBoard.cells[r] = make([]*Cell, b.length)
	}

 	for i := range b.cells {
		for j := range b.cells[i] {
 			newBoard.cells[i][j] = b.cells[i][j]
		}
 	}
	return &newBoard
}

//Input: takes a board and isParallel
//Output: divides the work between isParallel being true and false. returns a board after applying toppling operation on the cells
func PostToppling(b *Board, isParallel bool) *Board {
	
	if isParallel {
		numProcs := runtime.NumCPU()
		TopplingParallel(b , numProcs)
	} else {
		for i := range b.cells {
			for j := range b.cells[i]{
				for b.cells[i][j].numOfCoins >= 4 {
					b.cells[i][j] = TopplingOperation(b.cells[i][j])
				}
			}
		}
	}
	return b
}


//Input: takes a board and numProcs
//Output: divides the board into subslices to be used by different number of procs. 
func TopplingParallel(b *Board, numProcs int) {
	numRows := len(b.cells)
	finished := make(chan bool, numProcs)

	for i := 0; i < numProcs; i++ {
		startIndex := i * int(math.Floor(float64(numRows)/float64(numProcs)))
		var endIndex int
		if i < numProcs-1 {
			endIndex = (i+1)* int(math.Floor(float64(numRows)/float64(numProcs)))
		} else {
			endIndex = numRows
		}
		var b1 Board
		b1.height = b.height
		b1.length = endIndex - startIndex 
		b1.cells = b.cells[startIndex:endIndex]
		var b2 *Board
		b2 = &b1
		go TopplingOneProc(b2 , finished)
	}
	
	for i := 0; i < numProcs; i++ {
		<-finished
	}
	
}

//Input: takes a board and channel
//Output: runs the Toppling operation on the cells of the provided sublice of the board and return the message true to the channels when the job is done. 
func TopplingOneProc(b *Board, finished chan bool) {
	for i := range b.cells {
		for j := range b.cells[i] {
			for b.cells[i][j].numOfCoins >= 4 {
				b.cells[i][j] = TopplingOperation(b.cells[i][j])
			}			
		}
	}
	finished <- true
} 

//Input: a cell
//Output: returns a cell after appling the toppling operation
func TopplingOperation(cell *Cell) *Cell {
	for i:= range cell.neighbors{
		cell.neighbors[i].numOfCoins += 1
	}
		cell.numOfCoins += -4
	return cell
}

//Input: takes a board
//Output: returns the number of max coins
func MaxCoinsInCell(board *Board) int{
	maxCoins := 0
	for r := range board.cells {
		for c := range board.cells[r] {
			if board.cells[r][c].numOfCoins > maxCoins {
				maxCoins = board.cells[r][c].numOfCoins
			}
		}
	}
	return maxCoins
}



