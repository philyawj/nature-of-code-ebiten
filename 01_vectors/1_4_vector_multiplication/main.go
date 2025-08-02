// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.4 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/philyawj/nature-of-code-ebiten/p5math"
)

const (
	screenWidth  = 640
	screenHeight = 240
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	mouseX, mouseY := ebiten.CursorPosition()
	clampedX := float32(math.Max(0, math.Min(float64(mouseX), float64(screenWidth-1))))
	clampedY := float32(math.Max(0, math.Min(float64(mouseY), float64(screenHeight-1))))
	mouse := p5math.NewVector(clampedX, clampedY)

	center := p5math.NewVector(float32(screenWidth)/2, float32(screenHeight)/2)

	mouse.Sub(center)

	vector.StrokeLine(screen, center.X, center.Y, center.X+mouse.X, center.Y+mouse.Y, 2, color.RGBA{200, 200, 200, 255}, true)

	mouse.Mult(0.5)

	vector.StrokeLine(screen, center.X, center.Y, center.X+mouse.X, center.Y+mouse.Y, 4, color.Black, true)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Multiplying a Vector")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
