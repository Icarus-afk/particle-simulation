package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	quadtreeCapacity = 4
)

type Simulation struct {
	Particles []*Particle
	Quadtree  *Quadtree
}

func (s *Simulation) Update() {
	mouseX, mouseY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.AttractParticles(float64(mouseX), float64(mouseY))
	}

	// Create the quadtree and insert particles
	s.Quadtree = NewQuadtree(Rect{screenWidth / 2, screenHeight / 2, screenWidth / 2, screenHeight / 2}, quadtreeCapacity)
	for _, p := range s.Particles {
		p.Update()
		s.Quadtree.insert(p)
	}

	// Check for collisions using the quadtree
	for _, p := range s.Particles {
		rangeRect := Rect{p.X, p.Y, p.Radius * 2, p.Radius * 2}
		found := []*Particle{}
		s.Quadtree.query(rangeRect, &found)
		for _, other := range found {
			if p != other && p.CollidesWith(other) {
				p.ResolveCollision(other)
			}
		}
	}
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	for _, p := range s.Particles {
		for y := -int(p.Radius); y <= int(p.Radius); y++ {
			for x := -int(p.Radius); x <= int(p.Radius); x++ {
				if x*x+y*y <= int(p.Radius*p.Radius) {
					screen.Set(int(p.X)+x, int(p.Y)+y, color.White)
				}
			}
		}
	}
}

func (s *Simulation) AttractParticles(x, y float64) {
	for _, p := range s.Particles {
		dx := x - p.X
		dy := y - p.Y
		distanceSquared := dx*dx + dy*dy
		if distanceSquared < 1 {
			distanceSquared = 1
		}
		force := 100 / distanceSquared
		angle := math.Atan2(dy, dx)
		fx := force * math.Cos(angle)
		fy := force * math.Sin(angle)
		p.ApplyForce(fx, fy)
	}
}

func NewSimulation(numParticles int, speed float64) *Simulation {
	particles := make([]*Particle, numParticles)
	for i := range particles {
		particles[i] = &Particle{
			X:      float64(rand.Intn(screenWidth)),
			Y:      float64(rand.Intn(screenHeight)),
			VX:     (rand.Float64()*2 - 1) * speed,
			VY:     (rand.Float64()*2 - 1) * speed,
			Radius: 1,
		}
	}
	return &Simulation{
		Particles: particles,
		Quadtree:  NewQuadtree(Rect{screenWidth / 2, screenHeight / 2, screenWidth / 2, screenHeight / 2}, quadtreeCapacity),
	}
}
