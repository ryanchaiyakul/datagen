package genlib

import (
	"errors"
	"fmt"
)

//BoolParams are the inputs to GenInt
type BoolParams struct {
	Permutations []int
	Results      []bool
}

//Gen returns a 1D slice that contains the permutations of an integer
func (config *BoolParams) Gen() error {
	if len(config.Permutations) == 0 {
		return errors.New("GenInt : missing permutations")
	}

	for _, v := range config.Permutations {
		if v < 2 {
			boolVal := false
			if v == 1 {
				boolVal = true
			}
			config.Results = append(config.Results, boolVal)
		} else {
			return fmt.Errorf("GenInt : permutation : %v out of range", v)
		}
	}
	return nil
}

// PermutationCount returns the number of permutations that exist for a certain integer
func (config *BoolParams) PermutationCount() int {
	return 2
}

//Extract returns the value of IntParams
func (config *BoolParams) Extract(permutation int) (interface{}, error) {
	if permutation > len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	return config.Results[permutation], nil
}

//SetPermutation set the permutation of the IntParams
func (config *BoolParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
