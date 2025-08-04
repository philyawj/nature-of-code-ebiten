// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.10 from "The Nature of Code" by Daniel Shiffman,
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

type mover struct {
	position     *p5math.Vector
	velocity     *p5math.Vector
	acceleration *p5math.Vector
	topSpeed     float32
}

func newMover() *mover {
	return &mover{
		position:     p5math.NewVector(screenWidth/2, screenHeight/2),
		velocity:     p5math.NewVector(0, 0),
		acceleration: p5math.NewVector(0, 0),
		topSpeed:     5,
	}
}

func (m *mover) update() {
	mouseX, mouseY := ebiten.CursorPosition()
	boundedX := util.ConstrainToScreen(mouseX, screenWidth)
	boundedY := util.ConstrainToScreen(mouseY, screenHeight)
	mouse := p5math.NewVector(float32(boundedX), float32(boundedY))

	dir := p5math.SubVectors(mouse, m.position)
	dir.Normalize()
	dir.Mult(0.2)

	m.acceleration = dir

	m.velocity.Add(m.acceleration)
	m.velocity.Limit(m.topSpeed)
	m.position.Add(m.velocity)
}

func (m *mover) show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(m.position.X), float32(m.position.Y), 24, color.RGBA{127, 127, 127, 255}, true)
	vector.StrokeCircle(screen, float32(m.position.X), float32(m.position.Y), 24, 2, color.Black, true)
}

type game struct {
	mover *mover
}

func newGame() *game {
	return &game{
		mover: newMover(),
	}
}

func (g *game) Update() error {
	g.mover.update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.mover.show(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Accelerating Toward the Mouse")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
