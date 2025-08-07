// Copyright 2025 Jordan Philyaw
//
// Port of Example 2.3 from "The Nature of Code" by Daniel Shiffman,
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

func (m *mover) checkEdges() {
	if m.position.X > screenWidth-m.radius {
		m.position.X = screenWidth - m.radius
		m.velocity.X *= -1
	} else if m.position.X < m.radius {
		m.position.X = m.radius
		m.velocity.X *= -1
	}
	if m.position.Y > screenHeight-m.radius {
		m.position.Y = screenHeight - m.radius
		m.velocity.Y *= -1
	}
}

func (m *mover) show(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(m.position.X), float32(m.position.Y), m.radius, color.NRGBA{127, 127, 127, 127}, true)
	vector.StrokeCircle(screen, float32(m.position.X), float32(m.position.Y), m.radius, 2, color.Black, true)
}

type game struct {
	moverA *mover
	moverB *mover
}

func newGame() *game {
	return &game{
		moverA: newMover(200, 30, 10),
		moverB: newMover(440, 30, 2),
	}
}

func (g *game) Update() error {
	gravityA := p5math.MultVectors(gravity, g.moverA.mass)
	g.moverA.applyForce(gravityA)

	gravityB := p5math.MultVectors(gravity, g.moverB.mass)
	g.moverB.applyForce(gravityB)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.moverA.applyForce(wind)
		g.moverB.applyForce(wind)
	}

	g.moverA.update()
	g.moverA.checkEdges()

	g.moverB.update()
	g.moverB.checkEdges()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	g.moverA.show(screen)
	g.moverB.show(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Gravity Scaled by Mass")
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
