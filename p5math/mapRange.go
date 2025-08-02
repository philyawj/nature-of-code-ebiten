// Copyright 2025 Jordan Philyaw
//
// This file is part of a project licensed under CC BY-NC-SA 4.0.
// See LICENSE in the repository root for details.

package p5math

// MapRange is a port of the p5.js `map` function.
// It maps a value from one range to another.
func MapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
}
