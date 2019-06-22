package helperlib

import "errors"

// GetPermutationString returns the number of permutations of a string
func GetPermutationString(length int, stringValues [][]string) int {
	stringValuesInt := [][]int{}
	for _, v := range stringValues {
		tempSlice := []int{}
		for range v {
			tempSlice = append(tempSlice, 0)
		}
		stringValuesInt = append(stringValuesInt, tempSlice)
	}
	return GetPermutationIntSlice([]int{length}, stringValuesInt)
}

// ReshapeStringSlice tries to reshape the single int slice into a multidimensional slice
func ReshapeStringSlice(dimensions []int, flatSlice []string) (interface{}, error) {
	var ret interface{}
	switch len(dimensions) {
	case 1:
		ret = flatSlice
	case 2:
		temp := [][]string{}
		for i := 0; i < dimensions[0]; i++ {
			temp = append(temp, flatSlice[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]string{}
		for i := 0; i < dimensions[0]; i++ {
			doubleTemp := [][]string{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, flatSlice[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]string{}
		for i := 0; i < dimensions[0]; i++ {
			tripleTemp := [][][]string{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp := [][]string{}
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
