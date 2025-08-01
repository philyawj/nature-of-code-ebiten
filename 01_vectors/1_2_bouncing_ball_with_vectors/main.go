// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.2 from "The Nature of Code" by Daniel Shiffman,
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
)

type Game struct {
	position, velocity *p5math.Vector
}

func NewGame() *Game {
	return &Game{
		position: p5math.NewVector(100, 100),
		velocity: p5math.NewVector(2.5, 2),
	}
}

func (g *Game) Update() error {
	g.position.Add(g.velocity)

	if g.position.X > screenWidth || g.position.X < 0 {
		g.velocity.X *= -1
	}
	if g.position.Y > screenHeight || g.position.Y < 0 {
		g.velocity.Y *= -1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	vector.DrawFilledCircle(screen, float32(g.position.X), float32(g.position.Y), 24, color.RGBA{127, 127, 127, 255}, true)
	vector.StrokeCircle(screen, float32(g.position.X), float32(g.position.Y), 24, 2, color.Black, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bouncing Ball with Vectors!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
