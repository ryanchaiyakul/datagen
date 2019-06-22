package genlib

import (
	"errors"
	"fmt"

	helperlib "github.com/ryanchaiyakul/datagen/internal/helper"
)

// StringParams are the inputs to GenString
type StringParams struct {
	Length            int
	StringValues      [][]string
	StringValuesRaw   []string
	StringValuesIndex []int
	Permutations      []int
}

//StringSliceParams are the inputs to GenStringSlice
type StringSliceParams struct {
	Dimensions        []int
	Lengths           []int
	StringValues      [][]string
	StringValuesRaw   []string
	StringValuesIndex []int
	Permutations      []int
}

// GenString returns a list of string permutations that are valid to the asciiRange
// if asciiRange blank, uses all lowercase and uppercase letters
func GenString(config StringParams) ([]string, error) {
	// param checking
	if config.Length == 0 {
		return nil, errors.New("GenString : missing length")
	}
	if len(config.StringValues) == 0 {
		return nil, errors.New("GenString : missing stringValues")
	}

	// generate asciiValues from stringValues
	asciiValues := [][]int{}
	for i := 0; i < len(config.StringValues); i++ {
		temp := []int{}
		for _, v := range config.StringValues[i] {
			if len(v) != 1 {
				return nil, fmt.Errorf("GenString : invalid stringValues : %v", config.StringValues)
			}
			temp = append(temp, int([]byte(v)[0]))
		}
		asciiValues = append(asciiValues, temp)
	}

	intList, err := GenIntSlice(IntSliceParams{Dimensions: []int{config.Length}, ValidValues: asciiValues, Permutations: config.Permutations})

	retList := []string{}
	if err == nil {
		for k := range intList {
			ret := ""
			for _, v := range intList[k] {
				ret += string(v)
			}
			retList = append(retList, ret)
		}
		return retList, nil
	}
	return nil, err
}

// GenStringSlice returns a 2D slice of complex numbers that can be reshaped into the complex matrix requested
func GenStringSlice(config StringSliceParams) ([][]string, error) {
	if len(config.Dimensions) == 0 {
		return nil, errors.New("GenStringSlice : missing dimensions")
	}
	if len(config.StringValues) == 0 {
		return nil, errors.New("GenStringSlice : missing stringValues")
	}

	dimensionLength := helperlib.FlatIntSliceLength(config.Dimensions)
	if dimensionLength != len(config.Lengths) {
		return nil, fmt.Errorf("GenStringSlice : mismatched dimensions : %v or lengths : %v", config.Dimensions, config.Lengths)
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
				return nil, fmt.Errorf("GenString : invalid stringValues : %v", config.StringValues)
			}
			temp = append(temp, int([]byte(v)[0]))
		}
		asciiValues = append(asciiValues, temp)
	}

	intList, err := GenIntSlice(IntSliceParams{Dimensions: dimensions, ValidValues: asciiValues, Permutations: config.Permutations})
	if err == nil {
		retList := [][]string{}
		for _, v := range intList {
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
			retList = append(retList, tempSlice)
		}
		return retList, nil
	}
	return nil, err
}
