package helperlib

import (
	"math"
)

// GetPermutationSliceUnique returns the number of permutations that exist for a certain 1D slice
// validValues is [][]int
func GetPermutationSliceUnique(dimensions []int, validValues [][]int) int {
	ret := 1
	for i := 0; i < len(validValues); i++ {
		ret *= len(validValues[i])
	}
	return ret
}

// GetPermutationSliceUnanimous returns the number of permutations that exist for a certain 1D slice
// validValues is []int
func GetPermutationSliceUnanimous(dimensions []int, validValues []int) int {
	ret := 1
	if dimensionsLength := len(dimensions); dimensionsLength == 1 {
		ret = int(math.Pow(float64(len(validValues)), float64(dimensions[0])))
	} else {
		for i := 0; i < dimensionsLength; i++ {
			ret = ret * dimensions[i] * len(validValues)
		}
	}
	return ret
}

// FlatSliceLength returns the multidimensional slice flattened into a single dimensional slice
func FlatSliceLength(dimensions []int) int {
	ret := dimensions[0]
	if len(dimensions) > 1 {
		for _, v := range dimensions[1:] {
			ret *= v
		}
	}
	return ret
}
