package helperlib

import "errors"

// GetPermutationIntSlice returns the number of permutations that exist for a certain 1D slice
// validValues is [][]int
func GetPermutationIntSlice(validValues [][]int) int {
	ret := 1
	for i := 0; i < len(validValues); i++ {
		ret *= len(validValues[i])
	}
	return ret
}

// FlatIntSliceLength returns the multidimensional slice flattened into a single dimensional slice
func FlatIntSliceLength(dimensions []int) int {
	ret := dimensions[0]
	if len(dimensions) > 1 {
		for _, v := range dimensions[1:] {
			ret *= v
		}
	}
	return ret
}

// ReshapeIntSlice tries to reshape the single int slice into a multidimensional slice
func ReshapeIntSlice(dimensions []int, flatSlice []int) (interface{}, error) {
	var ret interface{}
	switch len(dimensions) {
	case 1:
		ret = flatSlice
	case 2:
		temp := [][]int{}
		for i := 0; i < dimensions[0]; i++ {
			temp = append(temp, flatSlice[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]int{}
		for i := 0; i < dimensions[0]; i++ {
			doubleTemp := [][]int{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, flatSlice[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]int{}
		for i := 0; i < dimensions[0]; i++ {
			tripleTemp := [][][]int{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp := [][]int{}
				for i := 0; i < dimensions[0]*dimensions[1]*dimensions[2]; i++ {
					doubleTemp = append(doubleTemp, flatSlice[i*dimensions[3]:(i+1)*dimensions[3]])
				}
				tripleTemp = append(tripleTemp, doubleTemp[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, tripleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	default:
		return nil, errors.New("Can not reshape to the number of dimensions requested")
	}
	return ret, nil
}
