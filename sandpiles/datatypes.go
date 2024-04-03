package main

import (
	"image"
	"github.com/llgcode/draw2d/draw2dimg"
)

type Canvas struct {
	gc     *draw2dimg.GraphicContext
	img    image.Image
	width  int // both width and height are in pixels
	height int
}

type Board struct {
	height int
	length int
	cells [][]*Cell
}

type Cell struct {
	numOfCoins int //number of coins
	neighbors []*Cell
}



