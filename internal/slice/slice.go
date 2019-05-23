package slicelib

import (
	"math"
	"sync"
)

// GenArray returns a multidimensional slice that is one dimensions large than the len(dimensions)
func GenArray(dimensions []int, validValues []int, permutationRange [2]int) interface{} {
	var ret interface{}
	permutationChan := make(chan []int)
	length := dimensions[0]
	if len(dimensions) > 1 {
		for _, v := range dimensions[1:] {
			length *= v
		}
	}

	if permutationRange[0] == 0 && permutationRange[1] == 0 {
		if dimensionsLength := len(dimensions); dimensionsLength == 1 {
			permutationRange = [2]int{0, int(math.Pow(float64(len(validValues)), float64(dimensions[0])))}
		} else {
			permutationRange[1] = 1
			for i := 0; i < dimensionsLength; i++ {
				permutationRange = [2]int{0, permutationRange[1] * dimensions[i]}
			}
			permutationRange = [2]int{0, int(math.Pow(float64(len(validValues)), float64(permutationRange[1])))}
		}
	}

	switch len(dimensions) {
	case 1:
		permutationList := [][]int{}
		go listPermutations(length, validValues, permutationChan, permutationRange)
		for permutation := range permutationChan {
			permutationList = append(permutationList, permutation)
		}
		ret = permutationList
	case 2:
		permutationList := [][][]int{}
		go listPermutations(length, validValues, permutationChan, permutationRange)
		for permutation := range permutationChan {
			permutationInArrayOne := [][]int{}
			for i := 0; i < dimensions[0]; i++ {
				permutationInArrayOne = append(permutationInArrayOne, permutation[i*dimensions[1]:(i+1)*dimensions[1]])
			}
			permutationList = append(permutationList, permutationInArrayOne)
		}
		ret = permutationList
	case 3:
		permutationList := [][][][]int{}
		go listPermutations(length, validValues, permutationChan, permutationRange)
		for permutation := range permutationChan {
			permutationInArrayOne := [][][]int{}
			for i := 0; i < dimensions[0]; i++ {
				permutationInArrayTwo := [][]int{}
				for j := 0; j < dimensions[1]*dimensions[0]; j++ {
					permutationInArrayTwo = append(permutationInArrayTwo, permutation[j*dimensions[2]:(j+1)*dimensions[2]])
				}
				permutationInArrayOne = append(permutationInArrayOne, permutationInArrayTwo[i*dimensions[1]:(i+1)*dimensions[1]])
			}
			permutationList = append(permutationList, permutationInArrayOne)
		}
		ret = permutationList
	case 4:
		permutationList := [][][][][]int{}
		go listPermutations(length, validValues, permutationChan, permutationRange)
		for permutation := range permutationChan {
			permutationInArrayOne := [][][][]int{}
			for i := 0; i < dimensions[0]; i++ {
				permutationInArrayTwo := [][][]int{}
				for j := 0; j < dimensions[1]*dimensions[0]; j++ {
					permutationInArrayThree := [][]int{}
					for k := 0; k < dimensions[2]*dimensions[1]*dimensions[0]; k++ {
						permutationInArrayThree = append(permutationInArrayThree, permutation[j*dimensions[3]:(j+1)*dimensions[3]])
					}
					permutationInArrayTwo = append(permutationInArrayTwo, permutationInArrayThree[j*dimensions[2]:(j+1)*dimensions[2]])
				}
				permutationInArrayOne = append(permutationInArrayOne, permutationInArrayTwo[i*dimensions[1]:(i+1)*dimensions[1]])
			}
			permutationList = append(permutationList, permutationInArrayOne)
		}
		ret = permutationList
	default:
		ret = nil
	}
	return ret
}

func listPermutations(length int, validValues []int, results chan []int, permutationRange [2]int) {
	base := []int{}
	for i := 0; i < length; i++ {
		base = append(base, validValues[0])
	}
	switch permutationCount := permutationRange[1] - permutationRange[0]; {
	case permutationCount < 1000:
		go listPermutationsHelper(base, validValues, permutationRange, results, 10)
	case permutationCount < 10000:
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
		currentIndex := sliceIndex(len(validValues), func(sliceI int) bool { return validValues[sliceI] == temp[i] })
		currentIndex += addend

		// addend is used as carry over
		addend = currentIndex / len(validValues)

		temp[i] = validValues[int(math.Mod(float64(currentIndex), float64(len(validValues))))]
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
		currentIndex := sliceIndex(len(validValues), func(sliceI int) bool { return validValues[sliceI] == temp[i] })
		currentIndex += addend

		// addend is used as carry over
		addend = currentIndex / len(validValues)

		temp[i] = validValues[int(math.Mod(float64(currentIndex), float64(len(validValues))))]
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
		currentIndex := sliceIndex(len(validValues), func(sliceI int) bool { return validValues[sliceI] == temp[i] })
		currentIndex += addend

		// addend is used as carry over
		addend = currentIndex / len(validValues)

		temp[i] = validValues[int(math.Mod(float64(currentIndex), float64(len(validValues))))]
		if addend == 0 {
			return temp
		}
	}
	return nil
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
