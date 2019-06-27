package genlib

import (
	"errors"
	"fmt"
)

//StringSliceParams are the inputs to GenStringSlice
type StringSliceParams struct {
	Dimensions   []int
	Lengths      []int
	StringValues [][]string
	Permutations []int
	Results      [][]string
}

// Gen returns a 2D slice of complex numbers that can be reshaped into the complex matrix requested
func (config *StringSliceParams) Gen() error {
	if len(config.Dimensions) == 0 {
		return errors.New("GenStringSlice : missing dimensions")
	}
	if len(config.StringValues) == 0 {
		return errors.New("GenStringSlice : missing stringValues")
	}

	dimensionLength := DimensionsLength(config.Dimensions)
	if dimensionLength != len(config.Lengths) {
		return fmt.Errorf("GenStringSlice : mismatched dimensions : %v or lengths : %v", config.Dimensions, config.Lengths)
	}

	dimensions := []int{}
	index := 0
	for _, length := range config.Dimensions {
		tempVal := 0
		for i := 0; i < length; i++ {
			tempVal += config.Lengths[index]
			index++
		}
		dimensions = append(dimensions, tempVal)
	}

	// generate asciiValues from stringValues
	asciiValues := [][]int{}
	for i := 0; i < len(config.StringValues); i++ {
		temp := []int{}
		for _, v := range config.StringValues[i] {
			if len(v) != 1 {
				return fmt.Errorf("GenString : invalid stringValues : %v", config.StringValues)
			}
			temp = append(temp, int([]byte(v)[0]))
		}
		asciiValues = append(asciiValues, temp)
	}

	sliceConfig := &IntSliceParams{Dimensions: dimensions, ValidValues: asciiValues, Permutations: config.Permutations}
	intChan, err := sliceConfig.GenChan()
	if err == nil {
		for v := range intChan {
			index := 0
			tempSlice := []string{}
			for _, length := range config.Lengths {
				tempString := ""
				for i := 0; i < length; i++ {
					tempString += string(v[index])
					index++
				}
				tempSlice = append(tempSlice, tempString)
			}
			config.Results = append(config.Results, tempSlice)
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations of a string slice
func (config *StringSliceParams) PermutationCount() int {
	ret := 1
	for i := 0; i < len(config.StringValues); i++ {
		ret *= len(config.StringValues[i])
	}
	return ret
}

// Extract tries to reshape the single string slice into a multidimensional slice
func (config *StringSliceParams) Extract(permutation int) (interface{}, error) {
	if permutation >= len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}

	var ret interface{}
	switch len(config.Dimensions) {
	case 1:
		ret = config.Results[permutation]
	case 2:
		temp := [][]string{}
		for i := 0; i < config.Dimensions[0]; i++ {
			temp = append(temp, config.Results[permutation][i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]string{}
		for i := 0; i < config.Dimensions[0]; i++ {
			doubleTemp := [][]string{}
			for j := 0; j < config.Dimensions[0]*config.Dimensions[1]; j++ {
				doubleTemp = append(doubleTemp, config.Results[permutation][j*config.Dimensions[2]:(j+1)*config.Dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]string{}
		for i := 0; i < config.Dimensions[0]; i++ {
			tripleTemp := [][][]string{}
			for j := 0; j < config.Dimensions[0]*config.Dimensions[1]; j++ {
				doubleTemp := [][]string{}
				for k := 0; k < config.Dimensions[0]*config.Dimensions[1]*config.Dimensions[2]; k++ {
					doubleTemp = append(doubleTemp, config.Results[permutation][k*config.Dimensions[3]:(k+1)*config.Dimensions[3]])
				}
				tripleTemp = append(tripleTemp, doubleTemp[j*config.Dimensions[2]:(j+1)*config.Dimensions[2]])
			}
			temp = append(temp, tripleTemp[i*config.Dimensions[1]:(i+1)*config.Dimensions[1]])
		}
		ret = temp
	default:
		return nil, errors.New("Can not reshape to the number of dimensions requested")
	}
	return ret, nil
}

//SetPermutation set the permutation of the StringSliceParams
func (config *StringSliceParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
