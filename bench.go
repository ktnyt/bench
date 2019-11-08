package bench

import "testing"

// F represents a benchmark function.
type F func(*testing.B)

// C creates a benchmark case for the given name and benchmark function.
func C(name string, bfs ...F) F {
	return func(b *testing.B) {
		b.Run(name, All(bfs...))
	}
}

// Apply the given testing.B object to the given benchmark functions.
func Apply(b *testing.B, bfs ...F) {
	for _, bf := range bfs {
		bf(b)
	}
}

// All combines the given benchmark functions into a single benchmark function.
func All(bfs ...F) F {
	return func(b *testing.B) {
		for _, bf := range bfs {
			bf(b)
		}
	}
}
