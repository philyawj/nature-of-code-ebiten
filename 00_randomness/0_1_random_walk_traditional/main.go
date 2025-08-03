// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.1 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 240
)

type walker struct {
	x int
	y int
}

func newWalker() *walker {
	return &walker{
		x: screenWidth / 2,
		y: screenHeight / 2,
	}
}

func (w *walker) step() {
	choice := rand.Intn(4)
	switch choice {
	case 0:
		w.x++
	case 1:
		w.x--
	case 2:
		w.y++
	case 3:
		w.y--
	}
}

type game struct {
	walker     *walker
	trailImage *ebiten.Image
}

func newGame() *game {
	trailImage := ebiten.NewImage(screenWidth, screenHeight)
	trailImage.Fill(color.White)
	return &game{
		walker:     newWalker(),
		trailImage: trailImage,
	}
}

func (g *game) Update() error {
	g.walker.step()
	g.trailImage.Set(int(g.walker.x), int(g.walker.y), color.Black)
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
	ebiten.SetWindowTitle("A Traditional Random Walk")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
