// Copyright 2025 Jordan Philyaw
//
// Port of Example 2.4 from "The Nature of Code" by Daniel Shiffman,
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
	gravity = p5math.NewVector(0, 1)
	wind    = p5math.NewVector(0.5, 0)
)

type mover struct {
	mass         float32
	radius       float32
	position     *p5math.Vector
	velocity     *p5math.Vector
	acceleration *p5math.Vector
}

func newMover(x, y, m float32) *mover {
	return &mover{
		mass:         m,
		radius:       m * 8,
		position:     p5math.NewVector(x, y),
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

func (m *mover) contactEdge() bool {
	return m.position.Y > screenHeight-m.radius-1
}

func (m *mover) bounceEdges() {
	bounce := float32(-0.9)
	if m.position.X > screenWidth-m.radius {
		m.position.X = screenWidth - m.radius
		m.velocity.X *= bounce
	} else if m.position.X < m.radius {
		m.position.X = m.radius
		m.velocity.X *= bounce
	}
	if m.position.Y > screenHeight-m.radius {
		m.position.Y = screenHeight - m.radius
		m.velocity.Y *= bounce
	}
}

func (m *mover) show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(m.position.X), float32(m.position.Y), m.radius, color.NRGBA{127, 127, 127, 127}, true)
	vector.StrokeCircle(screen, float32(m.position.X), float32(m.position.Y), m.radius, 2, color.Black, true)
}

type game struct {
	mover *mover
}

func newGame() *game {
	return &game{
		mover: newMover(screenWidth/2, 30, 5),
	}
}

func (g *game) Update() error {
	g.mover.applyForce(gravity)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.mover.applyForce(wind)
	}

	if g.mover.contactEdge() {
		c := float32(0.1)
		friction := p5math.MultVectors(g.mover.velocity, -1)
		friction.SetMag(c)
		g.mover.applyForce(friction)
	}

	g.mover.bounceEdges()
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
	ebiten.SetWindowTitle("Including Friction")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
