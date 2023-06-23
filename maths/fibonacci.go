package maths

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	current := 0
	next := 0
	prev := 1
	return func() int {
		current = next
		next = current + prev
		prev = current
		return current
	}

}
