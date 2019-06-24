package genlib

import (
	"errors"
)

// ComplexParams is the input to GenComplex
type ComplexParams struct {
	ValidValues     []complex128
	RealValues      []int
	ImaginaryValues []int
	Permutations    []int
}

// ComplexSliceParams is the input to GenComplexSlice
type ComplexSliceParams struct {
	Dimensions       []int
	ValidValues      [][]complex128
	RealValues       []int
	ImaginaryValues  []int
	ValidValuesIndex []int
	Permutations     []int
}

// GenComplex returns a 1D slice of complex numbers
func GenComplex(config ComplexParams) ([]complex128, error) {
	if len(config.ValidValues) == 0 {
		return nil, errors.New("GenComplexSlice : missing validValues")
	}

	tempValues := []int{}
	for i := 0; i < int(len(config.ValidValues)); i++ {
		tempValues = append(tempValues, i)
	}

	intList, err := GenInt(IntParams{ValidValues: tempValues, Permutations: config.Permutations})
	if err != nil {
		return nil, err
	}

	ret := []complex128{}
	for _, v := range intList {
		ret = append(ret, config.ValidValues[v])
	}
	return ret, nil
}

// GenComplexSlice returns a 2D slice of complex numbers that can be reshaped into the complex matrix requested
func GenComplexSlice(config ComplexSliceParams) ([][]complex128, error) {
	if len(config.ValidValues) == 0 {
		return nil, errors.New("GenComplexSlice : missing validValues")
	}

	tempValues := [][]int{}
	for _, v := range config.ValidValues {
		tempSlice := []int{}
		for i := 0; i < len(v); i++ {
			tempSlice = append(tempSlice, i)
		}
		tempValues = append(tempValues, tempSlice)
	}

	intList, err := GenIntSlice(IntSliceParams{Dimensions: config.Dimensions, ValidValues: tempValues, Permutations: config.Permutations})
	if err != nil {
		return nil, err
	}

	ret := [][]complex128{}
	for _, intSlice := range intList {
		tempRet := []complex128{}
		for k, v := range intSlice {
			tempRet = append(tempRet, config.ValidValues[k][v])
		}
		ret = append(ret, tempRet)
	}
	return ret, nil
}
