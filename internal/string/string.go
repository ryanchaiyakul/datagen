package stringlib

import (
	slicelib "github.com/ryanchaiyakul/datagen/internal/slice"
)

func GenString(length int, asciiRange [2]int, permutationRange [2]int) []string {
	validValues := []int{}
	if asciiRange[0] == 0 && asciiRange[1] == 0 {
		for i := 65; i < 91; i++ {
			validValues = append(validValues, i)
		}
		for i := 97; i < 123; i++ {
			validValues = append(validValues, i)
		}
	} else {
		for i := asciiRange[0]; i < asciiRange[1]+1; i++ {
			validValues = append(validValues, i)
		}
	}

	intList := slicelib.GenArray([]int{length}, validValues, permutationRange).([][]int)
	retList := []string{}
	for k := range intList {
		ret := ""
		for _, v := range intList[k] {
			ret += string(v)
		}
		retList = append(retList, ret)
	}
	return retList
}
