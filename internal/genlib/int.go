package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"golang.org/x/sync/semaphore"

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

	permutationCount := helperlib.GetPermutationIntSlice(sliceConfig.ValidValues)
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

	permutationList := [][]int{}
	permutationChan := make(chan []int, bufferCount)

	base := []int{}
	for i := 0; i < len(sliceConfig.ValidValues); i++ {
		base = append(base, sliceConfig.ValidValues[i][0])
	}

	go genIntSliceHelper(&base, &sliceConfig.ValidValues, &sliceConfig.Permutations, permutationChan, routineCount)
	for permutation := range permutationChan {
		permutationList = append(permutationList, permutation)
	}
	return permutationList, nil
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
