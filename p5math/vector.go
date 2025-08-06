// Copyright 2025 Jordan Philyaw
//
// This file is part of a project licensed under CC BY-NC-SA 4.0.
// See LICENSE in the repository root for details.

package p5math

import (
	"math"
	"math/rand"
)

// Vector is a port of the p5.js `p5.Vector` class.
// It represents a 2D vector with x and y coordinates.
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

// Sub is a port of the p5.js `vector.sub()` function.
// It subtracts two Vectors and modifies the original Vector.
func (v *Vector) Sub(v2 *Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

// SubVectors is a port of the p5.js `p5.Vector.sub()` function.
// It returns a new Vector that is the result of subtracting v2 from v1.
// This does not modify the original Vector.
func SubVectors(v1, v2 *Vector) *Vector {
	return &Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

// Mult is a port of the p5.js `vector.mult()` function.
// It multiplies the Vector by a scalar value and modifies the original Vector.
func (v *Vector) Mult(n float32) {
	v.X *= n
	v.Y *= n
}

// MultVectors is a port of the p5.js `p5.Vector.mult()` function.
// It returns a new Vector that is the result of multiplying v by n.
// This does not modify the original Vector.
func MultVectors(v *Vector, n float32) *Vector {
	return &Vector{
		X: v.X * n,
		Y: v.Y * n,
	}
}

// Div is a port of the p5.js `vector.div()` function.
// It divides the Vector by a scalar value and modifies the original Vector.
func (v *Vector) Div(n float32) {
	v.X /= n
	v.Y /= n
}

// DivVectors is a port of the p5.js `p5.Vector.div()` function.
// It returns a new Vector that is the result of dividing v by n.
// This does not modify the original Vector.
func DivVectors(v *Vector, n float32) *Vector {
	return &Vector{
		X: v.X / n,
		Y: v.Y / n,
	}
}

// Mag is a port of the p5.js `vector.mag()` function.
// It returns the magnitude (length) of the Vector.
func (v *Vector) Mag() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Normalize is a port of the p5.js `vector.normalize()` function.
// It normalizes the Vector to a unit Vector (length of 1) and modifies the original Vector.
func (v *Vector) Normalize() {
	m := v.Mag()
	if m > 0 {
		v.Div(m)
	}
}

// NormalizeVector is a port of the p5.js `p5.Vector.normalize()` function.
// It returns a new unit Vector (length of 1) version of v.
// This does not modify the original Vector.
func NormalizeVector(v *Vector) *Vector {
	m := v.Mag()
	if m > 0 {
		return &Vector{
			X: v.X / m,
			Y: v.Y / m,
		}
	}
	return &Vector{X: 0, Y: 0}
}

// Limit is a simplified port of the p5.js `vector.limit()` function.
// It limits the magnitude of the Vector to a maximum value.
func (v *Vector) Limit(max float32) {
	if v.Mag() > max {
		v.Normalize()
		v.Mult(max)
	}
}

// Random2D is a port of the p5.js `p5.Vector.random2D()` function.
// It returns a new unit Vector (length of 1) with a random direction.
func Random2D() *Vector {
	angle := rand.Float64() * 2 * math.Pi
	return &Vector{
		X: float32(math.Cos(angle)),
		Y: float32(math.Sin(angle)),
	}
}
