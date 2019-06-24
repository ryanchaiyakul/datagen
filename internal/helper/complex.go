package helperlib

import "errors"

// GetPermutationComplexSlice returns the number of permutations that exist for a certain 1D slice
// validValues is [][]complex128
func GetPermutationComplexSlice(validValues [][]complex128) int {
	ret := 1
	for i := 0; i < len(validValues); i++ {
		ret *= len(validValues[i])
	}
	return ret
}

// ReshapeComplexSlice tries to reshape the single int slice into a multidimensional slice
func ReshapeComplexSlice(dimensions []int, flatSlice []complex128) (interface{}, error) {
	var ret interface{}
	switch len(dimensions) {
	case 1:
		ret = flatSlice
	case 2:
		temp := [][]complex128{}
		for i := 0; i < dimensions[0]; i++ {
			temp = append(temp, flatSlice[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]complex128{}
		for i := 0; i < dimensions[0]; i++ {
			doubleTemp := [][]complex128{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, flatSlice[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]complex128{}
		for i := 0; i < dimensions[0]; i++ {
			tripleTemp := [][][]complex128{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp := [][]complex128{}
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
