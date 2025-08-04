// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.3 from "The Nature of Code" by Daniel Shiffman,
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

type game struct {
	mouse, center, sub *p5math.Vector
}

func newGame() *game {
	return &game{
		mouse:  p5math.NewVector(0, 0),
		center: p5math.NewVector(float32(screenWidth)/2, float32(screenHeight)/2),
		sub:    p5math.NewVector(0, 0),
	}
}

func (g *game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	boundedX := util.ConstrainToScreen(mouseX, screenWidth)
	boundedY := util.ConstrainToScreen(mouseY, screenHeight)
	g.mouse = p5math.NewVector(float32(boundedX), float32(boundedY))

	g.sub = p5math.SubVectors(g.mouse, g.center)

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	vector.StrokeLine(screen, 0, 0, g.mouse.X, g.mouse.Y, 4, color.RGBA{200, 200, 200, 255}, true)
	vector.StrokeLine(screen, 0, 0, g.center.X, g.center.Y, 4, color.RGBA{200, 200, 200, 255}, true)
	vector.StrokeLine(screen, g.center.X, g.center.Y, g.center.X+g.sub.X, g.center.Y+g.sub.Y, 4, color.Black, true)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Vector Subtraction")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
