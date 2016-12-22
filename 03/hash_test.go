// Benchmark to test whether O(1) multiplcation is faster than O(1) map lookup.
// Used to see whether memoization with a map was a performant solution to repeatedly multiplying values.
package hash

import (
	"testing"
)

var m map[[3]int]int

func init() {
	m = map[[3]int]int{
		[3]int{1, 2, 3}: 6,
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := [3]int{1, 2, 3}
		i, ok := m[k]
		_, _ = i, ok
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := 1 * 2 * 3
		_ = j
	}
}
