// Copyright 2025 Jordan Philyaw
//
// Port of Example 2.1 from "The Nature of Code" by Daniel Shiffman,
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

var (
	gravity = p5math.NewVector(0, 0.1)
	wind    = p5math.NewVector(0.1, 0)
)

type mover struct {
	mass         float32
	position     *p5math.Vector
	velocity     *p5math.Vector
	acceleration *p5math.Vector
}

func newMover() *mover {
	return &mover{
		mass:         1,
		position:     p5math.NewVector(screenWidth/2, 30),
		velocity:     p5math.NewVector(0, 0),
		acceleration: p5math.NewVector(0, 0),
	}
}

func (m *mover) applyForce(force *p5math.Vector) {
	f := p5math.DivVectors(force, m.mass)
	m.acceleration.Add(f)
}

func (m *mover) update() {
	m.velocity.Add(m.acceleration)
	m.position.Add(m.velocity)
	m.acceleration.Mult(0)
}

func (m *mover) checkEdges() {
	if m.position.X > screenWidth {
		m.position.X = screenWidth
		m.velocity.X *= -1
	} else if m.position.X < 0 {
		m.velocity.X *= -1
		m.position.X = 0
	}
	if m.position.Y > screenHeight {
		m.velocity.Y *= -1
		m.position.Y = screenHeight
	}
}

func (m *mover) show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(m.position.X), float32(m.position.Y), 24, color.NRGBA{127, 127, 127, 127}, true)
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
	g.mover.applyForce(gravity)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.mover.applyForce(wind)
	}

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
	ebiten.SetWindowTitle("Forces")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
