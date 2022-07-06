package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHight = 640, 360
	boidCount                = 500
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHight
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHight*2)
	ebiten.SetWindowTitle("Boids in a Box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
