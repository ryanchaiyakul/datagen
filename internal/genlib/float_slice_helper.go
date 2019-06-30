package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

// GenChan streams the resulting permutations of FloatSliceParams
func (config *FloatSliceParams) GenChan() (chan map[int][]float64, error) {
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

	permutationChan := make(chan map[int][]float64, bufferCount)

	base := []float64{}
	for i := 0; i < len(config.ValidValues); i++ {
		base = append(base, config.ValidValues[i][0])
	}

	go genFloatSliceHelper(&base, &config.ValidValues, &config.Permutations, permutationChan, routineCount)
	return permutationChan, nil
}

func genFloatSliceHelper(base *[]float64, validValues *[][]float64, permutations *[]int, results chan map[int][]float64, routineCount int) {
	sem := make(chan struct{})
	var wg sync.WaitGroup

	for _, v := range *permutations {
		wg.Add(1)
		sem <- struct{}{}
		go incrementSingleFloat(base, validValues, v, results, &wg, sem)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func incrementSingleFloat(base *[]float64, validValues *[][]float64, addend int, results chan map[int][]float64, wg *sync.WaitGroup, sem chan struct{}) {
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
			results <- map[int][]float64{addendCopy: temp}
			return
		}
	}
	results <- nil
}
