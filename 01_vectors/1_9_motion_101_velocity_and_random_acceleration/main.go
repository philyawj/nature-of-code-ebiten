// Copyright 2025 Jordan Philyaw
//
// Port of Example 1.9 from "The Nature of Code" by Daniel Shiffman,
// published by No Starch Press® Inc., licensed under CC BY-NC-SA 4.0.
// Rewritten in Go using the Ebitengine game engine.

package main

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/philyawj/nature-of-code-ebiten/p5math"
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
	m.acceleration = p5math.Random2D()
	m.acceleration.Mult(rand.Float32() * 2)

	m.velocity.Add(m.acceleration)
	m.velocity.Limit(m.topSpeed)
	m.position.Add(m.velocity)
}

func (m *mover) checkEdges() {
	if m.position.X > screenWidth {
		m.position.X = 0
	} else if m.position.X < 0 {
		m.position.X = screenWidth
	}
	if m.position.Y > screenHeight {
		m.position.Y = 0
	} else if m.position.Y < 0 {
		m.position.Y = screenHeight
	}
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
	g.mover.checkEdges()
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
	ebiten.SetWindowTitle("Motion 101 (Velocity and Random Acceleration)")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
