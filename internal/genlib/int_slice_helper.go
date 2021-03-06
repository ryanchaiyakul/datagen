package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

// GenChan streams the resulting permutations of IntSliceParams
func (config *IntSliceParams) GenChan() (chan map[int][]int, error) {
	if len(config.Dimensions) == 0 {
		return nil, errors.New("GenSlice : missing dimensions")
	}
	if len(config.ValidValues) == 0 {
		return nil, errors.New("GenSlice : missing ValidValues")
	}
	if DimensionsLength(config.Dimensions) != len(config.ValidValues) {
		return nil, fmt.Errorf("GenSlice : mismatched validValues : %v and dimensions : %v", len(config.ValidValues), DimensionsLength(config.Dimensions))
	}
	if len(config.Permutations) == 0 {
		return nil, errors.New("GenSlice : missing permutations")
	}
	permutationCount := config.PermutationCount()

	for _, v := range config.Permutations {
		if v < 0 {
			return nil, fmt.Errorf("GenSlice : permutation : %v out of lower bound", v)
		}
		if v > permutationCount {
			return nil, fmt.Errorf("GenSlice : permutation : %v out of higher bound", v)
		}
	}

	routineCount := 0
	bufferCount := 0
	switch {
	case permutationCount < 100:
		routineCount = 10
		bufferCount = 5
	case permutationCount < 100:
		routineCount = 50
		bufferCount = 10
	default:
		routineCount = 100
		bufferCount = 20
	}

	permutationChan := make(chan map[int][]int, bufferCount)

	base := []int{}
	for i := 0; i < len(config.ValidValues); i++ {
		base = append(base, config.ValidValues[i][0])
	}

	go genIntSliceHelper(&base, &config.ValidValues, &config.Permutations, permutationChan, routineCount)
	return permutationChan, nil
}

func genIntSliceHelper(base *[]int, validValues *[][]int, permutations *[]int, results chan map[int][]int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup
	wg.Add(len(*permutations))

	for _, v := range *permutations {
		sem <- struct{}{}
		go incrementSingleWGSem(base, validValues, v, results, &wg, sem)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func incrementSingleWGSem(base *[]int, validValues *[][]int, addend int, results chan map[int][]int, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

	temp := append((*base)[:0:0], (*base)...)
	addendCopy := addend
	for i := 0; i < len(temp); i++ {
		temp[i] = (*validValues)[i][int(math.Mod(float64(addend), float64(len((*validValues)[i]))))]

		addend = addend / len((*validValues)[i])

		if addend == 0 {
			results <- map[int][]int{addendCopy: temp}
			return
		}
	}
	results <- nil
}

// SliceIndex returns the location of the value in the slice
// returns -1 if it does not exist
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// DimensionsLength returns the length of the requested slice when flattened
func DimensionsLength(dimensions []int) int {
	ret := dimensions[0]
	if len(dimensions) > 1 {
		for _, v := range dimensions[1:] {
			ret *= v
		}
	}
	return ret
}
