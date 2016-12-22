package product

func product(s []int) []int {

	// create a new array
	// initialize new array with 1
	// for each number in old array
	// for each element in new array
	// multiply old[i] * new[j], except for ith
	// store product in new[j]
	// O(n^2)

	// create a new array
	// var leftNumbers
	// for each number in old array
	// if i is 0, =leftNumbers, continue
	// new array[i] = left numbers
	// left-numbers *= i

	output := make([]int, len(s))
	var leftNumbers int

	for i := 0; i < len(s); i++ {
		if i == 0 {
			leftNumbers = s[i]
			output[i] = 1
			continue
		}
		output[i] = leftNumbers
		leftNumbers *= s[i]
	}

	// for each number in old array, starting at len(array)-1
	// right-numbers *= i
	// if i is len(array)-1, continue
	// new array[i] *= right numbers
	var rightNumbers int
	for i := len(s) - 1; i > -1; i-- {
		if i == len(s)-1 {
			rightNumbers = s[i]
			continue
		}
		output[i] *= rightNumbers
		rightNumbers *= s[i]
	}
	return output
}
