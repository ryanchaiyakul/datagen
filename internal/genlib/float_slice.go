package genlib

import (
	"errors"
	"fmt"
)

// FloatSliceParams is for complex slice generation
type FloatSliceParams struct {
	Dimensions   []int
	ValidValues  [][]float64
	Permutations []int
	Results      [][]float64
}

// Gen returns a 2D slice that contains all the permutations mapped to a single slice
// if permutationRange is blank, returns all permutations
func (config *FloatSliceParams) Gen() error {
	permutationChan, err := config.GenChan()
	if err == nil {
		config.Results = make([][]float64, len(config.Permutations))
		for permutation := range permutationChan {
			for k, v := range permutation {
				index := SliceIndex(len(config.Permutations), func(i int) bool { return config.Permutations[i] == k })
				if index == -1 {
					return fmt.Errorf("Gen : unknown permutation : %v generated", k)
				}
				config.Results[index] = v
			}
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations that exist for a certain 1D slice
func (config *FloatSliceParams) PermutationCount() int {
	ret := 1
	for i := 0; i < len(config.ValidValues); i++ {
		ret *= len(config.ValidValues[i])
	}
	return ret
}

// Extract tries to reshape the single int slice into a multidimensional slice
func (config *FloatSliceParams) Extract(permutation int) (interface{}, error) {
	if permutation >= len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	var ret interface{}
	switch len(config.Dimensions) {
	case 1:
		ret = config.Results[permutation]
	case 2:
		temp := [][]float64{}
		for i := 0; i < config.Dimensions[0]; i++ {
			temp = append(temp, config.Results[permutation][i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]float64{}
		for i := 0; i < config.Dimensions[0]; i++ {
			doubleTemp := [][]float64{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, config.Results[permutation][i*config.Dimensions[2]:(i+1)*config.Dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]float64{}
		for i := 0; i < config.Dimensions[0]; i++ {
			tripleTemp := [][][]float64{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp := [][]float64{}
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
func (config *FloatSliceParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
