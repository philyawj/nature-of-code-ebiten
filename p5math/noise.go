// Copyright 2025 Jordan Philyaw
//
// This file is part of a project licensed under CC BY-NC-SA 4.0.
// See LICENSE in the repository root for details.

package p5math

import (
	"math"
	"math/rand"
)

// Noise is a simplified port of the p5.js `noise` function.
// It returns 1D Perlin noise with multiple octaves. The output is normalized to [0, 1].
// MapRange can then be used to scale the result to a desired range.
func Noise(x float64) float64 {
	octaves := 4
	falloff := 0.5
	amp := 1.0
	maxAmp := 0.0
	freq := 1.0
	sum := 0.0
	for i := 0; i < octaves; i++ {
		sum += perlinNoise1D(x*freq) * amp
		maxAmp += amp
		amp *= falloff
		freq *= 2
	}
	return sum / maxAmp
}

// perlinNoise1D is a simplified Perlin noise function.
// It is 1D and single-octave.
func perlinNoise1D(input float64) float64 {
	intPart := int(math.Floor(input)) & 255
	fractionalPart := input - math.Floor(input)
	fadeValue := fade(fractionalPart)
	gradientA := permutationTable[intPart]
	gradientB := permutationTable[intPart+1]
	return linearInterpolate(gradient(gradientA, fractionalPart), gradient(gradientB, fractionalPart-1), fadeValue)*0.5 + 0.5
}

var permutationTable = [512]int{}

func init() {
	perm := [256]int{}
	for i := range perm {
		perm[i] = i
	}
	for i := 255; i > 0; i-- {
		j := rand.Intn(i + 1)
		perm[i], perm[j] = perm[j], perm[i]
	}
	for i := range permutationTable {
		permutationTable[i] = perm[i&255]
	}
}

func fade(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}

func linearInterpolate(a, b, t float64) float64 {
	return a + t*(b-a)
}

func gradient(hash int, x float64) float64 {
	if hash&1 == 0 {
		return x
	}
	return -x
}
