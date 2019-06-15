package genlib

import (
	"errors"
)

// GenString returns a list of string permutations that are valid to the asciiRange
// if asciiRange blank, uses all lowercase and uppercase letters
func GenString(length int, stringValues [][]string, permutationRange [2]int) ([]string, error) {
	retList := []string{}
	asciiValues := [][]int{}
	for i := 0; i < len(stringValues); i++ {
		temp := []int{}
		for _, v := range stringValues[i] {
			if len(v) != 1 {
				return nil, errors.New("GenString : Invalid stringValues")
			}
			temp = append(temp, int([]byte(v)[0]))
		}
		asciiValues = append(asciiValues, temp)
	}

	intList, err := GenArray([]int{length}, asciiValues, permutationRange)
	if err == nil {
		for k := range intList.([][]int) {
			ret := ""
			for _, v := range intList.([][]int)[k] {
				ret += string(v)
			}
			retList = append(retList, ret)
		}
		return retList, err
	}
	return nil, err
}
