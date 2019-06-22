package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"

	helperlib "github.com/ryanchaiyakul/datagen/internal/helper"
)

//IntParams are the inputs to GenInt
type IntParams struct {
	ValidValues  []int
	Permutations []int
}

// IntSliceParams are the inputs to GenIntSlice
type IntSliceParams struct {
	Dimensions       []int
	ValidValues      [][]int
	ValidValuesRaw   []int
	ValidValuesIndex []int
	Permutations     []int
}

//GenInt returns a 1D slice that contains the permutations of an integer
func GenInt(config IntParams) ([]int, error) {
	if len(config.Permutations) == 0 {
		return nil, errors.New("GenInt : missing permutations")
	}
	if len(config.ValidValues) == 0 {
		return nil, errors.New("GenInt : missing validValues")
	}

	ret := []int{}
	for _, v := range config.Permutations {
		if v < len(config.ValidValues) {
			ret = append(ret, config.ValidValues[v])
		} else {
			return nil, fmt.Errorf("GenInt : permutation : %v out of range", v)
		}
	}
	return ret, nil
}

// GenIntSlice returns a 2D slice that contains all the permutations mapped to a single slice
// if permutationRange is blank, returns all permutations
func GenIntSlice(sliceConfig IntSliceParams) ([][]int, error) {
	// parameter checking
	if len(sliceConfig.Dimensions) == 0 {
		return nil, errors.New("GenSlice : missing dimensions")
	}
	if len(sliceConfig.ValidValues) == 0 {
		return nil, errors.New("GenSlice : missing ValidValues")
	}

	length := helperlib.FlatIntSliceLength(sliceConfig.Dimensions)
	if length != len(sliceConfig.ValidValues) {
		return nil, fmt.Errorf("GenSlice : mismatched validValues : %v and dimensions : %v", len(sliceConfig.ValidValues), length)
	}

	permutationCount := helperlib.GetPermutationIntSlice(sliceConfig.Dimensions, sliceConfig.ValidValues)
	if len(sliceConfig.Permutations) == 0 {
		for i := 0; i < permutationCount; i++ {
			sliceConfig.Permutations = append(sliceConfig.Permutations, i)
		}
	} else {
		for _, v := range sliceConfig.Permutations {
			if v < 0 {
				return nil, fmt.Errorf("GenSlice : permutation : %v out of lower bound", v)
			}
			if v > permutationCount+1 {
				return nil, fmt.Errorf("GenSlice : permutation : %v out of higher bound", v)
			}
		}
	}

	// initialize variables
	permutationList := [][]int{}
	permutationChan := make(chan []int)

	go listPermutations(length, sliceConfig.ValidValues, permutationChan, sliceConfig.Permutations)
	for permutation := range permutationChan {
		permutationList = append(permutationList, permutation)
	}
	return permutationList, nil
}

func listPermutations(length int, validValues [][]int, results chan []int, permutations []int) {
	// generate base slice
	base := []int{}
	for i := 0; i < len(validValues); i++ {
		base = append(base, validValues[i][0])
	}

	// call helper functions
	switch permutationCount := len(permutations); {
	case permutationCount < 100:
		go listPermutationsHelper(&base, &validValues, &permutations, results, 10)
	case permutationCount < 1000:
		go listPermutationsHelper(&base, &validValues, &permutations, results, 100)
	default:
		go listPermutationsHelper(&base, &validValues, &permutations, results, 500)
	}
}

func listPermutationsHelper(base *[]int, validValues *[][]int, permutations *[]int, results chan []int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup

	for _, v := range *permutations {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go incrementSingleWGSem(base, validValues, v, results, &wg, sem)
		default:
			incrementSingle(base, validValues, v, results)
			wg.Done()
		}
	}
	// to close the channel after all the workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
}

func incrementSingle(base *[]int, validValues *[][]int, addend int, results chan []int) {
	results <- incrementSingleRet(base, validValues, addend)
}

func incrementSingleWGSem(base *[]int, validValues *[][]int, addend int, results chan []int, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()
	incrementSingle(base, validValues, addend, results)
}

func incrementSingleRet(base *[]int, validValues *[][]int, addend int) []int {
	// a copy is used because otherwise you would mutilate the base list
	temp := append((*base)[:0:0], (*base)...)
	for i := 0; i < len(temp); i++ {
		temp[i] = (*validValues)[i][int(math.Mod(float64(addend), float64(len((*validValues)[i]))))]

		// addend is used as carry over
		addend = addend / len((*validValues)[i])

		if addend == 0 {
			return temp
		}
	}
	return nil
}
