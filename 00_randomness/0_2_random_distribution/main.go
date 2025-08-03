// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.2 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 240
	numBars      = 20
)

type game struct {
	barCounts [numBars]int
}

func newGame() *game {
	return &game{}
}

func (g *game) Update() error {
	index := rand.Intn(numBars)
	g.barCounts[index]++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	barWidth := float32(screenWidth) / float32(numBars)
	for x := range g.barCounts {
		barHeight := float32(g.barCounts[x])
		rectX := float32(x) * barWidth
		rectY := float32(screenHeight) - barHeight
		rectWidth := barWidth
		rectHeight := barHeight

		vector.DrawFilledRect(screen, rectX, rectY, rectWidth, rectHeight, color.RGBA{127, 127, 127, 255}, false)
		vector.StrokeRect(screen, rectX, rectY, rectWidth, rectHeight, 2, color.Black, false)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("A Random-Number Distribution")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
