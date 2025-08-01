// Copyright 2025 Jordan Philyaw
//
// Port of Example 0.6 from "The Nature of Code" by Daniel Shiffman,
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

type Walker struct {
	tx, ty float64
	x, y   float64
}

func NewWalker() *Walker {
	return &Walker{
		tx: 0,
		ty: 10000,
	}
}

func (w *Walker) step() {
	w.x = p5math.MapRange(p5math.Noise(w.tx), 0, 1, 0, screenWidth)
	w.y = p5math.MapRange(p5math.Noise(w.ty), 0, 1, 0, screenHeight)
	w.tx += 0.01
	w.ty += 0.01
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
	vector.DrawFilledCircle(g.trailImage, float32(g.walker.x), float32(g.walker.y), 24, color.RGBA{127, 127, 127, 255}, true)
	vector.StrokeCircle(g.trailImage, float32(g.walker.x), float32(g.walker.y), 24, 2, color.Black, true)
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
	ebiten.SetWindowTitle("Perlin Noise Walker")
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
