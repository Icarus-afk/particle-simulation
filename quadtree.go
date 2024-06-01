package main

type Quadtree struct {
	bounds    Rect
	particles []*Particle
	capacity  int
	divided   bool
	northwest *Quadtree
	northeast *Quadtree
	southwest *Quadtree
	southeast *Quadtree
}

type Rect struct {
	x, y, w, h float64
}

func (r Rect) contains(p *Particle) bool {
	return p.X >= r.x-r.w &&
		p.X < r.x+r.w &&
		p.Y >= r.y-r.h &&
		p.Y < r.y+r.h
}

func (r Rect) intersects(other Rect) bool {
	return !(other.x-other.w > r.x+r.w ||
		other.x+other.w < r.x-r.w ||
		other.y-other.h > r.y+r.h ||
		other.y+other.h < r.y-r.h)
}

func NewQuadtree(bounds Rect, capacity int) *Quadtree {
	return &Quadtree{
		bounds:    bounds,
		capacity:  capacity,
		particles: []*Particle{},
	}
}

func (qt *Quadtree) subdivide() {
	x := qt.bounds.x
	y := qt.bounds.y
	w := qt.bounds.w / 2
	h := qt.bounds.h / 2

	qt.northwest = NewQuadtree(Rect{x - w, y - h, w, h}, qt.capacity)
	qt.northeast = NewQuadtree(Rect{x + w, y - h, w, h}, qt.capacity)
	qt.southwest = NewQuadtree(Rect{x - w, y + h, w, h}, qt.capacity)
	qt.southeast = NewQuadtree(Rect{x + w, y + h, w, h}, qt.capacity)
	qt.divided = true
}

func (qt *Quadtree) insert(p *Particle) bool {
	if !qt.bounds.contains(p) {
		return false
	}

	if len(qt.particles) < qt.capacity {
		qt.particles = append(qt.particles, p)
		return true
	}

	if !qt.divided {
		qt.subdivide()
	}

	if qt.northwest.insert(p) || qt.northeast.insert(p) ||
		qt.southwest.insert(p) || qt.southeast.insert(p) {
		return true
	}

	return false
}

func (qt *Quadtree) query(rangeRect Rect, found *[]*Particle) {
	if !qt.bounds.intersects(rangeRect) {
		return
	}

	for _, p := range qt.particles {
		if rangeRect.contains(p) {
			*found = append(*found, p)
		}
	}

	if qt.divided {
		qt.northwest.query(rangeRect, found)
		qt.northeast.query(rangeRect, found)
		qt.southwest.query(rangeRect, found)
		qt.southeast.query(rangeRect, found)
	}
}
