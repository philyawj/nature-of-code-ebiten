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
// It adds two vectors together and modifies the original vector.
func (v *Vector) Add(v2 *Vector) {
	v.X += v2.X
	v.Y += v2.Y
}
