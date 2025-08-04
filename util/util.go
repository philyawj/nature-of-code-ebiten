package util

// ConstrainToScreen constrains an integer to be within the bounds of the screen dimensions.
// max is typically screenWidth or screenHeight.
// It ensures that the value does not go below 0 or above max-1.
func ConstrainToScreen(val, max int) int {
	if val < 0 {
		return 0
	}
	if val > max-1 {
		return max - 1
	}
	return val
}
