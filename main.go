package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 1080
	screenHeight = 720
)

type Game struct {
	simulation  *Simulation
	input       string
	speedSlider *Slider
}

type Slider struct {
	x, y, width, height int
	value               float64
}

func NewSlider(x, y, width, height int) *Slider {
	return &Slider{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		value:  0.5, // initial value
	}
}

func (s *Slider) Update() {
	// Check if the left mouse button is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Get the mouse position
		mx, my := ebiten.CursorPosition()

		// Check if the mouse is over the slider
		if mx >= s.x && mx <= s.x+s.width && my >= s.y && my <= s.y+s.height {
			// Update the slider value based on the mouse position
			s.value = float64(mx-s.x) / float64(s.width)
		}
	}
}

func (s *Slider) Draw(screen *ebiten.Image) {
	// Draw the slider bar
	bar := ebiten.NewImage(s.width, s.height)
	bar.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(bar, op)

	// Draw the slider handle
	handleX := s.x + int(s.value*float64(s.width))
	handle := ebiten.NewImage(10, s.height)
	handle.Fill(color.Gray{Y: 128})
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(handleX), float64(s.y))
	screen.DrawImage(handle, op)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Parse the input when the Enter key is pressed
		numParticles, err := strconv.Atoi(g.input)
		if err != nil {
			// handle error
		}

		// Use numParticles and speed to start your simulation
		g.simulation = NewSimulation(numParticles, g.speedSlider.value)
		g.input = "" // Clear the input
	} else {
		// Append the input characters to the input string
		g.input += string(ebiten.InputChars())
	}
	g.simulation.Update()
	g.speedSlider.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	text.Draw(screen, "Number of particles: "+g.input, basicfont.Face7x13, 10, 20, color.White)
	text.Draw(screen, "Speed: "+strconv.FormatFloat(g.speedSlider.value, 'f', 2, 64), basicfont.Face7x13, 10, 40, color.White)
	g.simulation.Draw(screen)
	g.speedSlider.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particle Simulator")
	game := &Game{
		simulation:  NewSimulation(1000, 0.25),  // default values
		speedSlider: NewSlider(10, 60, 200, 20), // slider position and size
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
