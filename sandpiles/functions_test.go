package main

import (
	"fmt"
	"testing"
)


func TestTopplingOperation(t *testing.T) {

	type Test struct {
		testCell *Cell
	}

	
	var correct Test 
	
	correct.testCell.numOfCoins = 4
	correct.testCell.neighbors = make([]*Cell, 4)
	correct.testCell.neighbors[0].numOfCoins = 1
	correct.testCell.neighbors[1].numOfCoins = 1
	correct.testCell.neighbors[2].numOfCoins = 1
	correct.testCell.neighbors[3].numOfCoins = 1

	output := TopplingOperation(correct.testCell)

	if correct.testCell != output {
		t.Errorf("This is incorrect")
	}
	fmt.Println("This is correct")


}