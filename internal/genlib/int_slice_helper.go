package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"golang.org/x/sync/semaphore"
)

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

// GenChan streams the resulting permutations of IntSliceParams
func (config *IntSliceParams) GenChan() (chan []int, error) {
	if len(config.Dimensions) == 0 {
		return nil, errors.New("GenSlice : missing dimensions")
	}
	if len(config.ValidValues) == 0 {
		return nil, errors.New("GenSlice : missing ValidValues")
	}

	if DimensionsLength(config.Dimensions) != len(config.ValidValues) {
		return nil, fmt.Errorf("GenSlice : mismatched validValues : %v and dimensions : %v", len(config.ValidValues), DimensionsLength(config.Dimensions))
	}

	permutationCount := config.PermutationCount()
	if len(config.Permutations) == 0 {
		for i := 0; i < permutationCount; i++ {
			config.Permutations = append(config.Permutations, i)
		}
	} else {
		for _, v := range config.Permutations {
			if v < 0 {
				return nil, fmt.Errorf("GenSlice : permutation : %v out of lower bound", v)
			}
			if v > permutationCount+1 {
				return nil, fmt.Errorf("GenSlice : permutation : %v out of higher bound", v)
			}
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

	permutationChan := make(chan []int, bufferCount)

	base := []int{}
	for i := 0; i < len(config.ValidValues); i++ {
		base = append(base, config.ValidValues[i][0])
	}

	go genIntSliceHelper(&base, &config.ValidValues, &config.Permutations, permutationChan, routineCount)
	return permutationChan, nil
}

func genIntSliceHelper(base *[]int, validValues *[][]int, permutations *[]int, results chan []int, routineCount int) {
	sem := semaphore.NewWeighted(int64(routineCount))
	var wg sync.WaitGroup

	for _, v := range *permutations {
		wg.Add(1)
		if sem.TryAcquire(1) {
			go incrementSingleWGSem(base, validValues, v, results, &wg, sem)
		} else {
			incrementSingle(base, validValues, v, results)
			wg.Done()
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func incrementSingle(base *[]int, validValues *[][]int, addend int, results chan []int) {
	temp := append((*base)[:0:0], (*base)...)
	for i := 0; i < len(temp); i++ {
		temp[i] = (*validValues)[i][int(math.Mod(float64(addend), float64(len((*validValues)[i]))))]

		addend = addend / len((*validValues)[i])

		if addend == 0 {
			results <- temp
			return
		}
	}
	results <- nil
}

func incrementSingleWGSem(base *[]int, validValues *[][]int, addend int, results chan []int, wg *sync.WaitGroup, sem *semaphore.Weighted) {
	defer func() {
		sem.Release(1)
		wg.Done()
	}()
	incrementSingle(base, validValues, addend, results)
}
