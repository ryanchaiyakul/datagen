package genlib

import (
	"errors"
	"fmt"
)

// ComplexParams is for complex number generation
type ComplexParams struct {
	ValidValues  []complex128
	Permutations []int
	Results      []complex128
}

// Gen returns a 1D slice of complex numbers
func (config *ComplexParams) Gen() error {
	if len(config.ValidValues) == 0 {
		return errors.New("GenComplexSlice : missing validValues")
	}

	tempValues := []int{}
	for i := 0; i < int(len(config.ValidValues)); i++ {
		tempValues = append(tempValues, i)
	}

	intConfig := &IntParams{ValidValues: tempValues, Permutations: config.Permutations}
	err := intConfig.Gen()
	if err == nil {
		for _, v := range intConfig.Results {
			config.Results = append(config.Results, config.ValidValues[v])
		}
		return nil
	}
	return err
}

// PermutationCount returns the number of permutations that exist for a certain 1D slice
func (config *ComplexParams) PermutationCount() int {
	return len(config.ValidValues)
}

// Extract returns the value of ComplexParams
func (config *ComplexParams) Extract(permutation int) (interface{}, error) {
	if permutation < len(config.Results) || permutation < 0 {
		return nil, fmt.Errorf("Extract : permutation : %v is out of range", permutation)
	}
	return config.Results[permutation], nil
}

//SetPermutation set the permutation of the ComplexParams
func (config *ComplexParams) SetPermutation(permutations []int) {
	config.Permutations = permutations
}
