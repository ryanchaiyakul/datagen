package genlib

import (
	"errors"
	"fmt"
)

// StringParams are the inputs to GenString
type StringParams struct {
	Length       int
	StringValues [][]string
	Permutations []int
	Results      []string
}

// Gen returns a list of string permutations that are valid to the asciiRange
// if asciiRange blank, uses all lowercase and uppercase letters
func (config *StringParams) Gen() error {
	// param checking
	if config.Length == 0 {
		return errors.New("GenString : missing length")
	}
	if len(config.StringValues) == 0 {
		return errors.New("GenString : missing stringValues")
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

	sliceConfig := &IntSliceParams{Dimensions: []int{config.Length}, ValidValues: asciiValues, Permutations: config.Permutations}
	intChan, err := sliceConfig.GenChan()

	if err == nil {
		config.Results = make([]string, len(config.Permutations))
		for permutation := range intChan {
			temp := ""
			for k, intSlice := range permutation {
				for _, v := range intSlice {
					temp += string(v)
				}
				index := SliceIndex(len(config.Permutations), func(i int) bool { return config.Permutations[i] == k })
				if index == -1 {
					return fmt.Errorf("Gen : unknown permutation : %v generated", k)
				}
				config.Results = append(config.Results, temp)
			}
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations of a string
func (config *StringParams) PermutationCount() int {
	ret := 1
	for i := 0; i < len(config.StringValues); i++ {
		ret *= len(config.StringValues[i])
	}
	return ret
}

//Extract returns the value of StringParams
func (config *StringParams) Extract(permutation int) (interface{}, error) {
	if permutation > len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	return config.Results[permutation], nil
}

//SetPermutation set the permutation of the StringParams
func (config *StringParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
