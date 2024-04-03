package main

import (
	"math"
)

//let's place our gravity simulation functions here.

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

//SimulateGravity simulates gravity over a series of snap shots separated by equal unit time.
//Input: an initial universe, a number of

func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	 timepoints := make([]Universe, numGens+1)
	 timepoints[0] = initialUniverse

	 //now range over the number of generations and update the universe
	 for i := 1; i <= numGens; i++ {
		 timepoints[i] = UpdateUniverse(timepoints[i-1], time)

	 }
	 return timepoints
}

func UpdateUniverse(currentUniverse Universe, time float64) Universe {
		newUniverse := CopyUniverse(currentUniverse)

		//range over all bodies in the universe and update their acceleration,
		//velocity, and position
		for _, b := range newUniverse.bodies {
					b.acceleration = UpdateAcceleration(newUniverse, b)
					b.velocity = UpdateVelocity(b, time)
					b.position = UpdatePosition(b, time)
		}
		return newUniverse
}

func CopyUniverse(currentUniverse Universe) Universe {
		var newUniverse Universe

		newUniverse.width = currentUniverse.width
		numBodies := len(currentUniverse.bodies)
		newUniverse.bodies = make([]Body, numBodies)

		for i := range currentUniverse.bodies {
			newUniverse.bodies[i].name = currentUniverse.bodies[i].name
			newUniverse.bodies[i].mass = currentUniverse.bodies[i].mass
			newUniverse.bodies[i].radius = currentUniverse.bodies[i].radius
			newUniverse.bodies[i].position.x = currentUniverse.bodies[i].position.x
			newUniverse.bodies[i].position.y = currentUniverse.bodies[i].position.y
			newUniverse.bodies[i].velocity.x = currentUniverse.bodies[i].velocity.x
			newUniverse.bodies[i].velocity.y = currentUniverse.bodies[i].velocity.y
			newUniverse.bodies[i].acceleration.x = currentUniverse.bodies[i].acceleration.x
			newUniverse.bodies[i].acceleration.y = currentUniverse.bodies[i].acceleration.y
			newUniverse.bodies[i].red = currentUniverse.bodies[i].red
			newUniverse.bodies[i].green = currentUniverse.bodies[i].green
			newUniverse.bodies[i].blue = currentUniverse.bodies[i].blue

		}

		return newUniverse
}




func UpdateVelocity(b Body, time float64) OrderedPair {
		var vel OrderedPair

		vel.x = b.velocity.x + b.acceleration * time
		vel.y = b.velocity.y + b.accelaration * time

		return vel
}

func UpdatePosition(b Body, time float64) OrderedPair {
		var pos OrderedPair

		pos.x = 0.5 * b.accelaration.x * time * time + b.velocity.x * time
		b.position.x
		pos.y = 0.5 * b.accelaration.y * time * time + b.velocity.y * time
		b.position.y

		return pos
	}

func UpdateAcceleration(currentUniverse Universe, b Body) OrderedPair {
		var accel OrderedPair

		force := ComputeNetForce(bodies, b)
		accel.x = force.x/b.mass
		accel.y = force.y/b.mass

		return accel
}

func ComputeNetForce(bodies []Body, b Body) OrderedPair {
	var netForce OrderedPair

	for i := range bodies {
			if bodies[i] != b {
					force := ComputeForce(b, bodies[i])

					netForce.x = force.x
					netForce.y = force.y
			}
	}

	return netForce
}

func ComputeForce(b1, b1 Body) OrderedPair{
		var force OrderedPair

		dist := Distance(b1.position, b2.position)
		F := G * b1.mass * b2.mass /(dist*dist)

		deltaX := b2.position.x - b1.position.x
		deltaY := b2.position.y - b1.position.y

		force.x = F*deltaX/dist
		force.y = F*deltaY/dist


		return force
}
