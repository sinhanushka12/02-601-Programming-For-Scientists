package main

import (
	"math"
)

//place your non-drawing functions here.

//Input: an initial Sky, a number of generations, and a time parameter (in seconds).
//Output: a slice of Sky objects corresponding to simulating the force of gravity over the number of generations time points.
func SimulateBoids(initialSky Sky, numGens int, timeStep float64) []Sky {
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialSky

	//now range over the number of generations and update the universe each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], timeStep)
	}

	return timePoints
}

//Input: a currentSky
//Output: updated current sky using updated acceleration, velocity, and position

func UpdateSky(currentSky Sky, timeStep float64) Sky {
	newSky := CopySky(currentSky)

	//range over all boids in the universe and update their acceleration,
	//velocity, and position
	for i := range newSky.boids {
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, currentSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(newSky, newSky.boids[i], timeStep)
		newSky.boids[i].position = UpdatePosition(newSky, newSky.boids[i], timeStep)
	}
	return newSky
}

//Input: Sky object and a body B
//Output: The net acceleration on B due to gravity calculated by every body in the Universe
func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(currentSky, currentSky.boids, b)

	//now, calculate acceleration (F = ma)! In this case, F = a since m = 1
	accel.x = force.x
	accel.y = force.y

	return accel
}

//Input: a slice of boid an one boid object. proximity, cohesionFactor, separationFactor, alignmentFactor (float64)
//Output: boids when present within a certain proximity, use their net force caused by cohesion, separation, and alignment to calculate the netForce in both x and y directions.
func ComputeNetForce(currentSky Sky, boids []Boid, b Boid) OrderedPair {
	var netForce OrderedPair
	var totalForce OrderedPair
	var allSeparation OrderedPair
	var allCohesion OrderedPair
	var allAlignment OrderedPair

	count := 0.0
	for i := range boids {
		if (boids[i] != b) && (Distance(b, boids[i]) <= currentSky.proximity) && (Distance(b, boids[i]) > 0.0) {
			count += 1.0
			//if boids == []
			cohesionForcee := ComputeCohesionForce(currentSky, b, boids[i])
			separationForcee := ComputeSeparationForce(currentSky, b, boids[i])
			alignmentForcee := ComputeAlignmentForce(currentSky, b, boids[i])

			//totalForce.x += ComputeCohesionForce(currentSky, b, boids[i]).x + ComputeSeparationForce(currentSky, b, boids[i]).x + ComputeAlignmentForce(currentSky, b, boids[i]).x
			//totalForce.y += ComputeCohesionForce(currentSky, b, boids[i]).y + ComputeSeparationForce(currentSky, b, boids[i]).y + ComputeAlignmentForce(currentSky, b, boids[i]).y
			allCohesion.x += cohesionForcee.x
			allCohesion.y += cohesionForcee.y
			allSeparation.x += separationForcee.x
			allSeparation.y += separationForcee.y
			allAlignment.x += alignmentForcee.x
			allAlignment.y += alignmentForcee.y
		}
	}

	if count != 0.0 {
		totalForce.x = allCohesion.x + allSeparation.x + allAlignment.x
		totalForce.y = allCohesion.y + allSeparation.y + allAlignment.y
		netForce.x = totalForce.x / count
		netForce.y = totalForce.y / count

	}

	return netForce
}

//Input: a slice of boids, one boid object, alignmentFactor (float64) proximity (float64)
//Output: if a boid is in the proximity, calculate the alignment forces
func ComputeAlignmentForce(currentSky Sky, b1 Boid, b2 Boid) OrderedPair {
	var alignmentForce OrderedPair
	alignmentForce.x = (currentSky.alignmentFactor * (b2.velocity.x)) / Distance(b2, b1)
	alignmentForce.y = (currentSky.alignmentFactor * (b2.velocity.y)) / Distance(b2, b1)
	return alignmentForce
}

//Input: two Boid objects (b1, b2 Boid), cohesionFactor (float64) and proximity (float64)
//Output: calculate and return the cohesion force as an OrderedPair
func ComputeCohesionForce(currentSky Sky, b1 Boid, b2 Boid) OrderedPair {
	var cohesionForce OrderedPair

	cohesionForce.x = (currentSky.cohesionFactor * (b2.position.x - b1.position.x)) / Distance(b2, b1)
	cohesionForce.y = (currentSky.cohesionFactor * (b2.position.y - b1.position.y)) / Distance(b2, b1)

	return cohesionForce

}

//Input: two Boid objects (b1, b2 Boid), separationFactor (float64) and proximity (float64)
//Output: calculate and return the separation force as an OrderedPair
func ComputeSeparationForce(currentSky Sky, b1 Boid, b2 Boid) OrderedPair {
	var separationForce OrderedPair

	separationForce.x = (currentSky.separationFactor * (b1.position.x - b2.position.x)) / (Distance(b2, b1) * Distance(b2, b1))
	separationForce.y = (currentSky.separationFactor * (b1.position.y - b2.position.y)) / (Distance(b2, b1) * Distance(b2, b1))

	return separationForce
}

//Input: a Sky object
//Output: a new Sky object, all of whose fields are copied over into the new Sky's fields. (Deep copy)
func CopySky(currentSky Sky) Sky {
	var newSky Sky

	newSky.width = currentSky.width
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed
	newSky.proximity = currentSky.proximity
	newSky.separationFactor = currentSky.separationFactor
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor

	numBoids := len(currentSky.boids)
	newSky.boids = make([]Boid, numBoids)

	//now, copy all of the boids' fields into our new boids
	for i := range currentSky.boids {
		newSky.boids[i].position.x = currentSky.boids[i].position.x
		newSky.boids[i].position.y = currentSky.boids[i].position.y
		newSky.boids[i].velocity.x = currentSky.boids[i].velocity.x
		newSky.boids[i].velocity.y = currentSky.boids[i].velocity.y
		newSky.boids[i].acceleration.x = currentSky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = currentSky.boids[i].acceleration.y

	}

	return newSky
}

//Input: a Boid object and a time step (float64).
//Output: The orderedPair corresponding to the velocity of this object after a single time step, using the body's current acceleration. Return the allowed maximum speed with the updated direction when the updated velocity is greater than maxBoidSpeed
func UpdateVelocity(currentSky Sky, b Boid, timeStep float64) OrderedPair {
	var vel OrderedPair
	//new velocity is current velocity + acceleration * time
	vel.x = b.velocity.x + b.acceleration.x*timeStep
	vel.y = b.velocity.y + b.acceleration.y*timeStep
	magnitudeOfVel := math.Sqrt((vel.x * vel.x) + (vel.y * vel.y))

	if magnitudeOfVel > currentSky.maxBoidSpeed {
		ratio := currentSky.maxBoidSpeed / magnitudeOfVel
		vel.x = vel.x * ratio
		vel.y = vel.y * ratio
	}
	return vel
}

//How do I know it's off grid? - compare to this width
//How do I update (x,y)? - keep the direction, change the magnitude
//Input: a Boid b, time step (float64), and proximity.
//Output: The OrderedPair corresponding to the updated position of the body after a single time step, using the body's current acceleration and velocity. If the updated position goes beyond the width of the canvas, bring it back to the starting point while maintaining the updated direction.
func UpdatePosition(currentSky Sky, b Boid, timeStep float64) OrderedPair {
	var pos OrderedPair

	pos.x = 0.5*b.acceleration.x*timeStep*timeStep + b.velocity.x*timeStep + b.position.x
	pos.y = 0.5*b.acceleration.y*timeStep*timeStep + b.velocity.y*timeStep + b.position.y

	for pos.x > currentSky.width {
		pos.x = pos.x - currentSky.width
	}
	for pos.y > currentSky.width {
		pos.y = pos.y - currentSky.width
	}
	for pos.x < 0 {
		pos.x += currentSky.width
	}
	for pos.y < 0 {
		pos.y += currentSky.width
	}
	return pos
}

//Input: 2 boid objects
//Output: squared distance
func Distance(b1 Boid, b2 Boid) float64 {
	distX := b2.position.x - b1.position.x
	distY := b2.position.y - b1.position.y

	return math.Sqrt((distX * distX) + (distY * distY))

}
