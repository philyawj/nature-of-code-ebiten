package p5math

// MapRange is a port of the p5.js `map` function.
// It maps a value from one range to another.
func MapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)/(inMax-inMin)*(outMax-outMin) + outMin
}
