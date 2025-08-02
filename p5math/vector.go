// Copyright 2025 Jordan Philyaw
//
// This file is part of a project licensed under CC BY-NC-SA 4.0.
// See LICENSE in the repository root for details.

package p5math

type Vector struct {
	X, Y float32
}

// NewVector is a port of the p5.js `createVector()` function.
// It creates a new Vector with the given x and y coordinates.
func NewVector(x, y float32) *Vector {
	return &Vector{X: x, Y: y}
}

// Add is a port of the p5.js `vector.add()` function.
// It adds two Vectors together and modifies the original Vector.
func (v *Vector) Add(v2 *Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
