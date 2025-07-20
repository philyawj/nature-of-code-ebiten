// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.3 from "The Nature of Code" by Daniel Shiffman,
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

type Walker struct {
	x int
	y int
}

func NewWalker() *Walker {
	return &Walker{
		x: screenWidth / 2,
		y: screenHeight / 2,
	}
}

func (w *Walker) step() {
	r := rand.Float64()
	if r < 0.4 {
		w.x++
	} else if r < 0.6 {
		w.x--
	} else if r < 0.8 {
		w.y++
	} else {
		w.y--
	}
	w.x = constrain(w.x, 0, screenWidth-1)
	w.y = constrain(w.y, 0, screenHeight-1)
}

func constrain(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

type Game struct {
	walker     *Walker
	trailImage *ebiten.Image
}

func NewGame() *Game {
	trailImage := ebiten.NewImage(screenWidth, screenHeight)
	trailImage.Fill(color.White)
	return &Game{
		walker:     NewWalker(),
		trailImage: trailImage,
	}
}

func (g *Game) Update() error {
	g.walker.step()
	g.trailImage.Set(int(g.walker.x), int(g.walker.y), color.Black)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.trailImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("A Walker That Tends to Move to the Right")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
