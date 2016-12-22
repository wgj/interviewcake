package product

import "testing"

func TestProduct(t *testing.T) {
	input := []int{1, 7, 3, 4}
	output := product(input)
	t.Error("expected: [84, 12, 28, 21]")
	t.Errorf("output: %v", output)
}
