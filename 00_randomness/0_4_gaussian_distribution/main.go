// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.4 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/philyawj/nature-of-code-ebiten/p5math"
)

const (
	screenWidth  = 640
	screenHeight = 240
	mean         = 320
	stddev       = 60
)

type game struct {
	trailImage *ebiten.Image
}

func newGame() *game {
	trailImage := ebiten.NewImage(screenWidth, screenHeight)
	trailImage.Fill(color.White)
	return &game{trailImage: trailImage}
}

func (g *game) Update() error {
	x := float32(p5math.RandomGaussian(mean, stddev))
	vector.DrawFilledCircle(g.trailImage, x, 120, 8, color.RGBA{0, 0, 0, 10}, true)
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.trailImage, nil)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("A Gaussian Distribution")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
