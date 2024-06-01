package main

import (
	"image/color"
	"math/rand"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	simulation *Simulation
}

func (g *Game) Update() error {
	g.simulation.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.simulation.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particle Simulator")
	rand.Seed(1)
	game := &Game{
		simulation: NewSimulation(0, 0), // Initialize with zero particles and speed
	}

	go func() {
		if err := ebiten.RunGame(game); err != nil {
			panic(err)
		}
	}()

	err := ui.Main(func() {
		numParticlesEntry := ui.NewEntry()
		speedEntry := ui.NewEntry()

		startButton := ui.NewButton("Start Simulation")
		startButton.OnClicked(func(*ui.Button) {
			numParticles, err := strconv.Atoi(numParticlesEntry.Text())
			if err != nil {
				// handle error
			}

			speed, err := strconv.ParseFloat(speedEntry.Text(), 64)
			if err != nil {
				// handle error
			}

			// Use numParticles and speed to start your simulation
			// You might need to modify your NewSimulation function to accept these parameters
			game.simulation = NewSimulation(numParticles, speed)
		})

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Number of Particles:"), false)
		box.Append(numParticlesEntry, false)
		box.Append(ui.NewLabel("Speed:"), false)
		box.Append(speedEntry, false)
		box.Append(startButton, false)

		window := ui.NewWindow("Particle Simulator", 200, 100, false)
		window.SetChild(box)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})

	if err != nil {
		panic(err)
	}
}
