// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.4 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/philyawj/nature-of-code-ebiten/p5math"
	"github.com/philyawj/nature-of-code-ebiten/util"
)

const (
	screenWidth  = 640
	screenHeight = 240
)

type game struct{}

func newGame() *game {
	return &game{}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	mouseX, mouseY := ebiten.CursorPosition()
	boundedX := util.ConstrainToScreen(mouseX, screenWidth)
	boundedY := util.ConstrainToScreen(mouseY, screenHeight)
	mouse := p5math.NewVector(float32(boundedX), float32(boundedY))

	center := p5math.NewVector(float32(screenWidth)/2, float32(screenHeight)/2)

	mouse.Sub(center)

	vector.StrokeLine(screen, center.X, center.Y, center.X+mouse.X, center.Y+mouse.Y, 2, color.RGBA{200, 200, 200, 255}, true)

	mouse.Mult(0.5)

	vector.StrokeLine(screen, center.X, center.Y, center.X+mouse.X, center.Y+mouse.Y, 4, color.Black, true)

}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Multiplying a Vector")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
