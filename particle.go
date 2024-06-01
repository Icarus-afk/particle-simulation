package main

import (
	"math"
)

type Particle struct {
	X, Y   float64
	VX, VY float64
	Radius float64
}

func (p *Particle) Update() {
	p.X += p.VX
	p.Y += p.VY

	// Reverse velocity component if particle collides with screen borders
	if p.X < 0 || p.X > screenWidth {
		// If particle goes out of left or right boundary, reverse its X velocity
		p.VX = -p.VX
	}
	if p.Y < 0 || p.Y > screenHeight {
		// If particle goes out of top or bottom boundary, reverse its Y velocity
		p.VY = -p.VY
	}
}

func (p *Particle) ApplyForce(fx, fy float64) {
	p.VX += fx
	p.VY += fy
}

func (p *Particle) CollidesWith(other *Particle) bool {
	dx := p.X - other.X
	dy := p.Y - other.Y
	distance := dx*dx + dy*dy
	return distance <= (p.Radius+other.Radius)*(p.Radius+other.Radius)
}

func (p *Particle) ResolveCollision(other *Particle) {
	dx := other.X - p.X
	dy := other.Y - p.Y
	distance := dx*dx + dy*dy
	if distance == 0 {
		return
	}

	// Normalize the collision direction
	nx := dx / math.Sqrt(distance)
	ny := dy / math.Sqrt(distance)

	// Calculate the relative velocity
	rx := p.VX - other.VX
	ry := p.VY - other.VY

	// Calculate the relative velocity along the collision direction
	dotProduct := rx*nx + ry*ny

	// If particles are moving away from each other, do not resolve collision
	if dotProduct > 0 {
		return
	}

	// Calculate the impulse
	j := (2 * dotProduct) / (p.Radius + other.Radius)

	// Update the velocity of the particles
	p.VX -= j * other.Radius * nx
	p.VY -= j * other.Radius * ny
	other.VX += j * p.Radius * nx
	other.VY += j * p.Radius * ny
}
