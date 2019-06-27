package genlib

import (
	"errors"
	"fmt"
)

// IntSliceParams are the inputs to GenIntSlice
type IntSliceParams struct {
	Dimensions   []int
	ValidValues  [][]int
	Permutations []int
	Results      [][]int
}

// Gen returns a 2D slice that contains all the permutations mapped to a single slice
// if permutationRange is blank, returns all permutations
func (config *IntSliceParams) Gen() error {
	permutationChan, err := config.GenChan()
	if err == nil {
		for permutation := range permutationChan {
			config.Results = append(config.Results, permutation)
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations that exist for a certain 1D slice
func (config *IntSliceParams) PermutationCount() int {
	ret := 1
	for i := 0; i < len(config.ValidValues); i++ {
		ret *= len(config.ValidValues[i])
	}
	return ret
}

// Extract tries to reshape the single int slice into a multidimensional slice
func (config *IntSliceParams) Extract(permutation int) (interface{}, error) {
	if permutation >= len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	var ret interface{}
	switch len(config.Dimensions) {
	case 1:
		ret = config.Results[permutation]
	case 2:
		temp := [][]int{}
		for i := 0; i < config.Dimensions[0]; i++ {
			temp = append(temp, config.Results[permutation][i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]int{}
		for i := 0; i < config.Dimensions[0]; i++ {
			doubleTemp := [][]int{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, config.Results[permutation][i*config.Dimensions[2]:(i+1)*config.Dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]int{}
		for i := 0; i < config.Dimensions[0]; i++ {
			tripleTemp := [][][]int{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp := [][]int{}
				for i := 0; i < config.Dimensions[0]*config.Dimensions[1]*config.Dimensions[2]; i++ {
					doubleTemp = append(doubleTemp, config.Results[permutation][i*config.Dimensions[3]:(i+1)*config.Dimensions[3]])
				}
				tripleTemp = append(tripleTemp, doubleTemp[i*config.Dimensions[2]:(i+1)*config.Dimensions[2]])
			}
			temp = append(temp, tripleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	default:
		return nil, errors.New("Extract : can not reshape to the number of dimensions requested")
	}
	return ret, nil
}

//SetPermutation set the permutation of the IntSliceParams
func (config *IntSliceParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
