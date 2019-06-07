package genlibinternal

import (
	"errors"
	"math"
	"sync"
)

// GenArray returns a 2D slice that contains all the permutations mapped to a single slice
// if permutationRange is blank, returns all permutations
func GenArray(dimensions []int, validValues []int, permutationRange [2]int) (interface{}, error) {
	permutationChan := make(chan []int)

	length := dimensions[0]
	if len(dimensions) > 1 {
		for _, v := range dimensions[1:] {
			length *= v
		}
	}

	if permutationRange[0] < 0 {
		return nil, errors.New("Slicelib : permutation lower limit out of range")
	}
	if permutationCount := maxPermutations(dimensions, validValues); permutationRange[1] > permutationCount-1 {
		return nil, errors.New("Slicelib : permutation lower limit out of range")
	} else if permutationRange[0] == 0 && permutationRange[1] == 0 {
		permutationRange[1] = permutationCount
	}

	permutationList := [][]int{}
	go listPermutations(length, validValues, permutationChan, permutationRange)
	for permutation := range permutationChan {
		permutationList = append(permutationList, permutation)
	}
	return permutationList, nil
}

func maxPermutations(dimensions []int, validValues []int) int {
	ret := 0
	if dimensionsLength := len(dimensions); dimensionsLength == 1 {
		ret = int(math.Pow(float64(len(validValues)), float64(dimensions[0])))
	} else {
		ret = 1
		for i := 0; i < dimensionsLength; i++ {
			ret = ret * dimensions[i]
		}
		ret = int(math.Pow(float64(len(validValues)), float64(ret)))
	}
	return ret
}

func listPermutations(length int, validValues []int, results chan []int, permutationRange [2]int) {
	base := []int{}
	for i := 0; i < length; i++ {
		base = append(base, validValues[0])
	}
	switch permutationCount := permutationRange[1] - permutationRange[0]; {
	case permutationCount < 100:
		go listPermutationsHelper(base, validValues, permutationRange, results, 10)
	case permutationCount < 1000:
		go listPermutationsHelper(base, validValues, permutationRange, results, 100)
	default:
		go listPermutationsHelperRange(base, validValues, permutationRange, results, 1000)
	}
}

func listPermutationsHelper(base []int, validValues []int, permutationRange [2]int, results chan []int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup

	for i := permutationRange[0]; i < permutationRange[1]; i++ {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go incrementSingleWGSem(base, validValues, i, results, &wg, sem)
		default:
			incrementSingle(base, validValues, i, results)
			wg.Done()
		}
	}
	// to close the channel after all the workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
}

func listPermutationsHelperRange(base []int, validValues []int, permutationRange [2]int, results chan []int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup

	for i := 0; i < routineCount-1; i++ {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go incrementRangeWGSem(base, validValues, [2]int{(permutationRange[0] / routineCount) * i, (permutationRange[1] / routineCount) * (i + 1)}, results, &wg, sem)
		default:
			incrementRange(base, validValues, [2]int{(permutationRange[0] / routineCount) * i, (permutationRange[1] / routineCount) * (i + 1)}, results)
			wg.Done()
		}
	}
	// to close the channel after all the workers are done
	go func() {
		wg.Wait()
		close(results)
	}()
}

func incrementRangeWGSem(base []int, validValues []int, addendRange [2]int, results chan []int, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

	for i := addendRange[0]; i < addendRange[1]; i++ {
		results <- incrementSingleRet(base, validValues, i)
	}
}

func incrementRange(base []int, validValues []int, addendRange [2]int, results chan []int) {
	for i := addendRange[0]; i < addendRange[1]; i++ {
		results <- incrementSingleRet(base, validValues, i)
	}
}

func incrementSingle(base []int, validValues []int, addend int, results chan []int) {
	// a copy is used because otherwise you would mutilate the base list
	temp := append(base[:0:0], base...)
	for i := 0; i < len(temp); i++ {
		temp[i] = validValues[int(math.Mod(float64(addend), float64(len(validValues))))]

		// addend is used as carry over
		addend = addend / len(validValues)

		if addend == 0 {
			results <- temp
			return
		}
	}
	results <- nil
}

func incrementSingleWGSem(base []int, validValues []int, addend int, results chan []int, wg *sync.WaitGroup, sem chan struct{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

	// a copy is used because otherwise you would mutilate the base list
	temp := append(base[:0:0], base...)
	for i := 0; i < len(temp); i++ {
		temp[i] = validValues[int(math.Mod(float64(addend), float64(len(validValues))))]

		// addend is used as carry over
		addend = addend / len(validValues)

		if addend == 0 {
			results <- temp
			return
		}
	}
	results <- nil
}

func incrementSingleRet(base []int, validValues []int, addend int) []int {
	// a copy is used because otherwise you would mutilate the base list
	temp := append(base[:0:0], base...)
	for i := 0; i < len(temp); i++ {
		temp[i] = validValues[int(math.Mod(float64(addend), float64(len(validValues))))]

		// addend is used as carry over
		addend = addend / len(validValues)

		if addend == 0 {
			return temp
		}
	}
	return nil
}
