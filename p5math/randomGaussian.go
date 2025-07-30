package p5math

import (
	"math"
	"math/rand"
)

// RandomGaussian is a simplified port of the p5.js `randomGaussian` function.
// It generates a random number from a Gaussian distribution with the specified mean and standard deviation
// using the Box-Muller transform.
func RandomGaussian(mean, stddev float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z0 := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0*stddev + mean
}
