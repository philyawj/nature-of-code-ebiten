// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.4 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 240
	mean         = 320
	stddev       = 60
)

type Game struct {
	trailImage *ebiten.Image
}

func NewGame() *Game {
	trailImage := ebiten.NewImage(screenWidth, screenHeight)
	trailImage.Fill(color.White)
	return &Game{trailImage: trailImage}
}

func (g *Game) Update() error {
	x := float32(randomGaussian(mean, stddev))
	vector.DrawFilledCircle(g.trailImage, x, 120, 8, color.RGBA{0, 0, 0, 10}, true)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.trailImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func randomGaussian(mean, stddev float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z0 := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0*stddev + mean
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("A Gaussian Distribution")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
