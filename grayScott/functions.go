package main

//place your functions from the assignment here.


func InitializeBoard(numRows, numCols int) Board {

    var board Board
    board = make(Board, numRows)

    for r := range board {
        board[r] = make([]Cell, numCols)
    }
    return board
}


func SimulateGrayScott(initialBoard Board, numGens int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) []Board {
	    boards := make([]Board, numGens + 1)
		  boards[0] = initialBoard
	    var t Board
	    t = initialBoard
	    for i := 1; i <= numGens; i++{
			    boards[i] = UpdateBoard(t, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel)
	          	t = boards[i]
	    }
		return boards
	}

func UpdateBoard(currentBoard Board, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Board {
		r := len(currentBoard)
		c := len(currentBoard[0])
	    newBoard := make(Board, r)
	    for row:= 0; row< r; row++{
	        for col := 0; col< c; col++{
				newBoard[row] = append(newBoard[row], UpdateCell(currentBoard, row, col, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate, kernel))
	        }
	    }
		return newBoard
	}

func UpdateCell(currentBoard Board, row, col int, feedRate, killRate, preyDiffusionRate, predatorDiffusionRate float64, kernel [3][3]float64) Cell{
	    var currentCell Cell
	    currentCell = currentBoard[row][col]
	    var diffusionValues Cell
	    diffusionValues = ChangeDueToDiffusion(currentBoard, row, col, preyDiffusionRate, predatorDiffusionRate, kernel)
	    var reactionValues Cell
	    reactionValues = ChangeDueToReactions(currentCell, feedRate, killRate)
	    var t Cell
	    t = SumCells(currentCell, diffusionValues, reactionValues)
	    return t
	}

func ChangeDueToDiffusion(mat Board, i int, j int, prey float64, predator float64, k [3][3]float64) Cell {
		r := len(mat)
		c := len(mat[0])
		var mult Cell
		mult[0] = 0.0
		mult[1] = 0.0
		t := 0.0

		for a := -1; a <= 1; a++ {
			for b := -1; b <= 1; b++ {
				if i+a < 0 || i+a >= r || j+b < 0 || j+b >= c {
					t += 1
				} else {
					mult[0] += mat[i+a][j+b][0] * k[a+1][b+1] * prey
					mult[1] += mat[i+a][j+b][1] * k[a+1][b+1] * predator
				}
			}
		}
		return mult
	}

func ChangeDueToReactions(currentCell Cell, feedRate, killRate float64) Cell {
    var x Cell
    change1 := feedRate*(1-currentCell[0]) - currentCell[0] * currentCell[1]*currentCell[1]
	change2 := -killRate*currentCell[1] + currentCell[0] * currentCell[1]*currentCell[1]
	x[0] = change1
	x[1] = change2
	return x
	}

func SumCells(cells ...Cell) Cell {
	var y Cell
	sum := 0.0
	sum2 := 0.0
	for i := range cells {
		//fmt.Println(cells[i][0])
		sum += cells[i][0]
        sum2 += cells[i][1]
	}
    y[0] = sum
    y[1]= sum2
	return y
	}
