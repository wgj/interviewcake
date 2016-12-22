package product

// Can we use a number more than once?

// naive solution
// for every number in list, i
//   for every number in list, j
//     for every number in list, k
// multiply i*j*k
// compare with the largest we've seen so far, and store it if it's larger

O(n3)

// Don't think we can get away from n3, since we have to compare every triple
// But we can cut down on the amount of operations, by memoizing products in a map
// multiply is way faster than map access


func highest_product([]int) int {

}