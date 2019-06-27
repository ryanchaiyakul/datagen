package genlib

import (
	"errors"
	"fmt"
)

// ComplexSliceParams is for complex slice generation
type ComplexSliceParams struct {
	Dimensions   []int
	ValidValues  [][]complex128
	Permutations []int
	Results      [][]complex128
}

// Gen returns a 2D slice of complex numbers that can be reshaped into the complex matrix requested
func (config *ComplexSliceParams) Gen() error {
	if len(config.ValidValues) == 0 {
		return errors.New("GenComplexSlice : missing validValues")
	}

	tempValues := [][]int{}
	for _, v := range config.ValidValues {
		tempSlice := []int{}
		for i := 0; i < len(v); i++ {
			tempSlice = append(tempSlice, i)
		}
		tempValues = append(tempValues, tempSlice)
	}
	sliceConfig := &IntSliceParams{Dimensions: config.Dimensions, ValidValues: tempValues, Permutations: config.Permutations}
	intChan, err := sliceConfig.GenChan()
	if err == nil {
		for intSlice := range intChan {
			tempRet := []complex128{}
			for k, v := range intSlice {
				tempRet = append(tempRet, config.ValidValues[k][v])
			}
			config.Results = append(config.Results, tempRet)
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations that exist for a certain 1D slice
func (config *ComplexSliceParams) PermutationCount() int {
	ret := 1
	for i := 0; i < len(config.ValidValues); i++ {
		ret *= len(config.ValidValues[i])
	}
	return ret
}

// Extract tries to reshape the single complex slice into a multidimensional slice
func (config *ComplexSliceParams) Extract(permutation int) (interface{}, error) {
	if permutation >= len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}

	var ret interface{}
	switch len(config.Dimensions) {
	case 1:
		ret = config.Results[permutation]
	case 2:
		temp := [][]complex128{}
		for i := 0; i < config.Dimensions[0]; i++ {
			temp = append(temp, config.Results[permutation][i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]complex128{}
		for i := 0; i < config.Dimensions[0]; i++ {
			doubleTemp := [][]complex128{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, config.Results[permutation][i*config.Dimensions[2]:(i+1)*config.Dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]complex128{}
		for i := 0; i < config.Dimensions[0]; i++ {
			tripleTemp := [][][]complex128{}
			for i := 0; i < config.Dimensions[0]*config.Dimensions[1]; i++ {
				doubleTemp := [][]complex128{}
				for i := 0; i < config.Dimensions[0]*config.Dimensions[1]*config.Dimensions[2]; i++ {
					doubleTemp = append(doubleTemp, config.Results[permutation][i*config.Dimensions[3]:(i+1)*config.Dimensions[3]])
				}
				tripleTemp = append(tripleTemp, doubleTemp[i*config.Dimensions[2]:(i+1)*config.Dimensions[2]])
			}
			temp = append(temp, tripleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	default:
		return nil, errors.New("Can not reshape to the number of dimensions requested")
	}
	return ret, nil
}

//SetPermutation set the permutation of the ComplexSliceParams
func (config *ComplexSliceParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
