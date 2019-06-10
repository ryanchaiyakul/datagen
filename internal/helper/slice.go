package helperlib

import "math"

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

// FlatSlice returns the multidimensional slice flattened into a single dimensional slice
func FlatSlice(dimensions []int, value int) []int {
	ret := []int{}
	for i := 0; i < len(dimensions); i++ {
		for j := 0; j < dimensions[i]; j++ {
			ret = append(ret, value)
		}
	}
	return ret
}
