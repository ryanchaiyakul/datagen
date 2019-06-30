package genlib

import (
	"errors"
	"fmt"
)

//FloatParams are the inputs to GenInt
type FloatParams struct {
	ValidValues  []float64
	Permutations []int
	Results      []float64
}

//Gen returns a 1D slice that contains the permutations of an integer
func (config *FloatParams) Gen() error {
	if len(config.Permutations) == 0 {
		return errors.New("GenInt : missing permutations")
	}
	if len(config.ValidValues) == 0 {
		return errors.New("GenInt : missing validValues")
	}

	for _, v := range config.Permutations {
		if v < len(config.ValidValues) {
			config.Results = append(config.Results, config.ValidValues[v])
		} else {
			return fmt.Errorf("GenInt : permutation : %v out of range", v)
		}
	}
	return nil
}

// PermutationCount returns the number of permutations that exist for a certain integer
func (config *FloatParams) PermutationCount() int {
	return len(config.ValidValues)
}

//Extract returns the value of IntParams
func (config *FloatParams) Extract(permutation int) (interface{}, error) {
	if permutation > len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	return config.Results[permutation], nil
}

//SetPermutation set the permutation of the IntParams
func (config *FloatParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
