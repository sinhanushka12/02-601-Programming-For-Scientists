package main

import (
	"canvas"
	"os"
	"log"
	"bufio"
	"image/png"
	"fmt"
	
)


func DrawBoard(b *Board, cellWidth int) canvas.Canvas {

	height := b.height
	width := b.length

	c := canvas.CreateNewCanvas(height, width)

	for i := range b.cells {
		for j := range b.cells[i] {

			if b.cells[i][j].numOfCoins == 0 {
				c.SetFillColor(canvas.MakeColor(0,0,0))
				x := i * cellWidth
				y := j * cellWidth
				c.ClearRect(x, y, x+cellWidth, y+cellWidth)
				c.Fill()
			} else if b.cells[i][j].numOfCoins == 1 { //lightGray
				c.SetFillColor(canvas.MakeColor(85,85,85))
				x := i * cellWidth
				y := j * cellWidth
				c.ClearRect(x, y, x+cellWidth, y+cellWidth)
				c.Fill()
			} else if b.cells[i][j].numOfCoins == 2 {
				c.SetFillColor(canvas.MakeColor(170,170,170))
				x := i * cellWidth
				y := j * cellWidth
				c.ClearRect(x, y, x+cellWidth, y+cellWidth)
				c.Fill()
			} else if b.cells[i][j].numOfCoins == 3 {
				c.SetFillColor(canvas.MakeColor(255,255,255))
				x := i * cellWidth
				y := j * cellWidth
				c.ClearRect(x, y, x+cellWidth, y+cellWidth)
				c.Fill()
			} 
		
		}
	}

	return c
}

func (c *Canvas) SaveToPNG(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, c.img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}