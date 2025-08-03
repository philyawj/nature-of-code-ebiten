// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.1 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 240
)

type game struct {
	x, y, xspeed, yspeed float32
}

func newGame() *game {
	return &game{
		x:      100,
		y:      100,
		xspeed: 2.5,
		yspeed: 2,
	}
}

func (g *game) Update() error {
	g.x += g.xspeed
	g.y += g.yspeed

	if g.x > screenWidth || g.x < 0 {
		g.xspeed *= -1
	}
	if g.y > screenHeight || g.y < 0 {
		g.yspeed *= -1
	}

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	vector.DrawFilledCircle(screen, float32(g.x), float32(g.y), 24, color.RGBA{127, 127, 127, 255}, true)
	vector.StrokeCircle(screen, float32(g.x), float32(g.y), 24, 2, color.Black, true)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bouncing Ball with No Vectors")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
