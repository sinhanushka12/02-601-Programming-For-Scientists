package main

import (
	"math"
)

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	timePoints[0] = initialUniverse
	// Your code goes here. Use subroutines! :)
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time, theta)
	}
	return timePoints
}

func UpdateUniverse(currentUniverse *Universe, time float64, theta float64) *Universe {
	newUniverse := CopyUniverse(currentUniverse)
	newNodes := InitializeIntoNodes(newUniverse)

	for i := range newUniverse.stars {
		newUniverse.stars[i].acceleration = UpdateAcceleration(newNodes, newUniverse.stars[i], theta)
		newUniverse.stars[i].velocity = UpdateVelocity(newUniverse.stars[i], time)
		newUniverse.stars[i].position = UpdatePosition(newUniverse.stars[i], time)
	}
	return newUniverse
}

func CopyUniverse(currentUniverse *Universe) *Universe {
	var newUniverse Universe
	newUniverse.width = currentUniverse.width

	numStars := len(currentUniverse.stars)
	newUniverse.stars = make([]*Star, numStars)

	for i := range newUniverse.stars {
		var newStar Star
		newUniverse.stars[i] = &newStar
	}

	//now, copy all of the stars fields into our new stars
	for i := range currentUniverse.stars {
		//var stari *Star
		//newUniverse.stars = append(newUniverse.stars , stari)
		newUniverse.stars[i].position.x = currentUniverse.stars[i].position.x
		newUniverse.stars[i].position.y = currentUniverse.stars[i].position.y
		newUniverse.stars[i].velocity.x = currentUniverse.stars[i].velocity.x
		newUniverse.stars[i].velocity.y = currentUniverse.stars[i].velocity.y
		newUniverse.stars[i].acceleration.x = currentUniverse.stars[i].acceleration.x
		newUniverse.stars[i].acceleration.y = currentUniverse.stars[i].acceleration.y
		newUniverse.stars[i].mass = currentUniverse.stars[i].mass
		newUniverse.stars[i].radius = currentUniverse.stars[i].radius
		newUniverse.stars[i].red = currentUniverse.stars[i].red
		newUniverse.stars[i].blue = currentUniverse.stars[i].blue
		newUniverse.stars[i].green = currentUniverse.stars[i].green
	}

	return &newUniverse
}

func UpdateAcceleration(listOfNodes []*Node, s1 *Star, theta float64) OrderedPair {
	var currentTree QuadTree
	currentTree = ConstructQuadTree(listOfNodes)
	var newAccel OrderedPair

	force := ComputeNetForce(currentTree, s1, theta)

	newAccel.x = force.x / s1.mass
	newAccel.y = force.y / s1.mass

	return newAccel
}

func UpdateVelocity(b *Star, time float64) OrderedPair {
	var vel OrderedPair
	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*time
	vel.y = b.velocity.y + b.acceleration.y*time

	return vel
}

func UpdatePosition(b *Star, time float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	pos.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

	return pos

}

//This function uses width/2 to divide the universe into subsquares (quadrants)
//Input: universe (list of star)
//Output: return 4 quadrants
func InitializeIntoNodes(currentUniverse *Universe) []*Node {

	//have a list of node pointers that has elements corresponding to each star in the universe
	Nodes1 := make([]*Node, len(currentUniverse.stars))

  for i := range Nodes1 {
		var star1 Star
		Nodes1[i].star = &star1

    for i := range currentUniverse.stars {
        Nodes[i].star = currentUniverse.stars[i]
    }

	return Nodes1
}

//makes a map that matches any quadrant to nodes inside of it
func RecordQuadrantsOfStars(listOfNodes []*Node) map[Quadrant]([]*Node) {
	listChildren := make(map[Quadrant]([]*Node))

	for i := range listOfNodes { //range over the array of node pointers
		_, isPresent := listChildren[listOfNodes[i].sector]
		if isPresent { //if the quadrant that the current node belongs to has already been observed, add current node to that quadrant
			listChildren[listOfNodes[i].sector] = append(listChildren[listOfNodes[i].sector], listOfNodes[i])
		} else { //register the quadrant into the map and add node to quadrant
			l := make([]*Node, 0)
			listChildren[listOfNodes[i].sector] = append(l, listOfNodes[i])
		}
	}
	return listChildren
}

func ConstructQuadTree(Nodes []*Node) QuadTree {
	//declare the empty tree
	var barnesTree QuadTree
	var x *Node
	barnesTree.root = x           //declare the root to be of type Node
	x.children = make([]*Node, 4) // declare the four children to be nil

	//don't forget about updating the center of gravity and mass
	var listChildren map[Quadrant]([]*Node)
	listChildren = RecordQuadrantsOfStars(Nodes)

	for _, nodes := range listChildren {
		if len(nodes) == 1 { //if there is only star in the quadrant, set it to the root's children
			barnesTree.root.children = append(barnesTree.root.children, nodes[0])
		} else { //convert the quadrant with more than one star into a universe and run ConstructQuadTree
			updatedNodes := UpdateQuadrant(nodes)
			quadTree := ConstructQuadTree(updatedNodes) // make another quad tree for that quadrant
			barnesTree.root.children = append(barnesTree.root.children, quadTree.root)
		}
	}

	//Let's calculate the mass for the nodes
	var star *Star //making another variable so i don't have to type root.Node.star everytime
	x.star = star
	star.mass = 0.0
	star.position.x = 0.0
	star.position.y = 0.0

	starsInNodes := make([]*Star, 0)
	for i := range Nodes {
		starsInNodes = append(starsInNodes, Nodes[i].star)
	}

	for i := range starsInNodes {
		star.mass += starsInNodes[i].mass
		star.position.x += starsInNodes[i].position.x * starsInNodes[i].mass
		star.position.y += starsInNodes[i].position.y * starsInNodes[i].mass
	}

	star.position.x = star.position.x / star.mass
	star.position.y = star.position.y / star.mass

	return barnesTree
}

func UpdateQuadrant(listOfNodes []*Node) []*Node {
	nqw := listOfNodes[0].sector.width / 2 //new quadrant width

	blx := listOfNodes[0].sector.x //bottom left x coordinate
	bry := listOfNodes[0].sector.y //bottom right y coordinate

	var NW Quadrant
	NW.x = blx
	NW.y = bry - nqw
	NW.width = nqw

	var NE Quadrant
	NE.x = blx + nqw
	NE.y = bry - nqw
	NE.width = nqw

	var SW Quadrant
	SW.x = blx
	SW.y = bry
	SW.width = nqw

	var SE Quadrant
	SE.x = blx + nqw
	SE.y = bry
	SE.width = nqw

	updatedNodes := make([]*Node, 0)

	for i := range listOfNodes {
		//check for the star x and y coordinates and assign them to appropriate quadrant using sector
		if (listOfNodes[i].star.position.x <= blx+nqw) && (listOfNodes[i].star.position.y <= bry-nqw) {
			listOfNodes[i].sector = NW
		} else if (listOfNodes[i].star.position.x >= blx+nqw) && (listOfNodes[i].star.position.y <= bry-nqw) {
			listOfNodes[i].sector = NE
		} else if (listOfNodes[i].star.position.x <= blx+nqw) && (listOfNodes[i].star.position.y >= bry-nqw) {
			listOfNodes[i].sector = SW
		} else if (listOfNodes[i].star.position.x >= blx+nqw) && (listOfNodes[i].star.position.y >= bry-nqw) {
			listOfNodes[i].sector = SE
		}
		updatedNodes = append(updatedNodes, listOfNodes[i])
	}

	return updatedNodes
}

func ComputeNetForce(anyTree QuadTree, body *Star, theta float64) OrderedPair {
	var netForce OrderedPair

	var s float64
	var d float64

	//To determine if a node is sufficiently far away, compute the quotient s/d s is the width of the region represented by the internal node d is the distance between the body and the node's center of mass
	//Then compare this ratio against a threshold value of theta
	//if s/d <= theta, then treat the internal node as a single body whose mass is equal to the sum of bodies in the subtree beneath it and whose position is the center of gravity of these bodies. Compute the force that this dummy body exerts on b and add this force to the net force acting on b
	//if s/d > theta, then iterate on the internal node's children

	//first if statement: when the index has a leaf --> use ComputeForce
	//second if statement: when index has an internal node and s/d > theta
	//thrid if statemnt : when index has an internal node and s/d <= theta

	for i := range anyTree.root.children {
		//if the star is not the body inputted and if the leaf does have a star assigned and if it is not an internal node
		//does anyTree.root.children[i] != nil mean it is a leaf that is not an internal node?
		s = anyTree.root.children[i].sector.width
		d = Distance(anyTree.root.children[i].star.position, body.position)
		if (anyTree.root.children[i].star != body) && (anyTree.root.children[i] != nil) && (anyTree.root.children[i].children == nil) { //&& (anyTree.root.children[i] != )
			netForce.x += ComputeForce(anyTree.root.children[i], body).x
			netForce.y += ComputeForce(anyTree.root.children[i], body).y
		} else {
			if s/d > theta {
				var newTree QuadTree
				newTree.root = anyTree.root.children[i]
				netForce.x += ComputeNetForce(newTree, body, theta).x
				netForce.y += ComputeNetForce(newTree, body, theta).y
			} else if anyTree.root.children[i] != nil {
				netForce.x += ComputeForce(anyTree.root.children[i], body).x
				netForce.y += ComputeForce(anyTree.root.children[i], body).y
			}
		}
	}
	return netForce
}

func ComputeForce(body1 *Node, body2 *Star) OrderedPair {
	var force OrderedPair

	dist := Distance(body1.star.position, body2.position)

	gravitationalForce := (G * body1.star.mass * body2.mass) / (dist * dist)
	distanceX := body2.position.x - body1.star.position.x
	distanceY := body2.position.y - body1.star.position.y

	force.x = gravitationalForce * distanceX / dist
	force.y = gravitationalForce * distanceY / dist

	return force
}

func Distance(a1, a2 OrderedPair) float64 {

	deltaX := a2.x - a1.x
	deltaY := a2.y - a1.y

	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))

}
