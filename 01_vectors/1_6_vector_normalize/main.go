// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.6 from "The Nature of Code" by Daniel Shiffman,
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
	mouseVec, centerVec, subVec, normVec *p5math.Vector
}

func newGame() *game {
	return &game{
		mouseVec:  p5math.NewVector(0, 0),
		centerVec: p5math.NewVector(float32(screenWidth)/2, float32(screenHeight)/2),
		subVec:    p5math.NewVector(0, 0),
		normVec:   p5math.NewVector(0, 0),
	}
}

func (g *game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	boundedX := util.ConstrainToScreen(mouseX, screenWidth)
	boundedY := util.ConstrainToScreen(mouseY, screenHeight)
	g.mouseVec = p5math.NewVector(float32(boundedX), float32(boundedY))

	g.subVec = p5math.SubVectors(g.mouseVec, g.centerVec)

	g.normVec = p5math.NormalizeVector(g.subVec)
	g.normVec.Mult(50)

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	vector.StrokeLine(screen, g.centerVec.X, g.centerVec.Y, g.centerVec.X+g.subVec.X, g.centerVec.Y+g.subVec.Y, 2, color.RGBA{200, 200, 200, 255}, true)
	vector.StrokeLine(screen, g.centerVec.X, g.centerVec.Y, g.centerVec.X+g.normVec.X, g.centerVec.Y+g.normVec.Y, 8, color.Black, true)

}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Normalizing a Vector")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
