package datagen

import (
	"math"
	"sync"
)

// GenArray returns a multidimensional slice that is one dimensions large than the len(dimensions)
func GenArray(dimensions []int, validValues []int) interface{} {
	var ret interface{}
	switch len(dimensions) {
	case 1:
		ret = gen1DArray(dimensions, validValues)
	case 2:
		ret = gen2DArray(dimensions, validValues)
	case 3:
		ret = gen3DArray(dimensions, validValues)
	case 4:
		ret = gen4DArray(dimensions, validValues)
	default:
		ret = nil
	}
	return ret
}

func gen1DArray(dimensions []int, validValues []int) [][]int {
	permutationchan := make(chan []int)
	permutationList := [][]int{}

	length := dimensions[0] * dimensions[1]

	go listPermutations(length, validValues, permutationchan)
	for permutation := range permutationchan {
		permutationList = append(permutationList, permutation)
	}
	return permutationList
}

func gen2DArray(dimensions []int, validValues []int) [][][]int {
	permutationchan := make(chan []int)
	permutationList := [][][]int{}

	length := dimensions[0] * dimensions[1]

	go listPermutations(length, validValues, permutationchan)
	for permutation := range permutationchan {
		permutationInArrayOne := [][]int{}
		for i := 0; i < dimensions[0]; i++ {
			permutationInArrayOne = append(permutationInArrayOne, permutation[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		permutationList = append(permutationList, permutationInArrayOne)
	}
	return permutationList
}

func gen3DArray(dimensions []int, validValues []int) [][][][]int {
	permutationchan := make(chan []int)
	permutationList := [][][][]int{}

	length := dimensions[0] * dimensions[1] * dimensions[2]

	go listPermutations(length, validValues, permutationchan)
	for permutation := range permutationchan {
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
	return permutationList
}

func gen4DArray(dimensions []int, validValues []int) [][][][][]int {
	permutationchan := make(chan []int)
	permutationList := [][][][][]int{}

	length := dimensions[0] * dimensions[1] * dimensions[2]

	go listPermutations(length, validValues, permutationchan)
	for permutation := range permutationchan {
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
	return permutationList
}

func listPermutations(length int, validValues []int, results chan []int) {
	base := []int{}
	for i := 0; i < length; i++ {
		base = append(base, validValues[0])
	}

	permutationcount := int(math.Pow(float64(length), float64(len(validValues))))
	if length == 1 {
		permutationcount = length * len(validValues)
	}

	// can this be a switch statement ?
	if permutationcount < 100 {
		go listPermutationsHelper(base, validValues, permutationcount, results, 10)
	} else if permutationcount < 1000 {
		go listPermutationsHelper(base, validValues, permutationcount, results, 100)
	} else {
		go listPermutationsHelperRange(base, validValues, permutationcount, results, 1000)
	}
}

func listPermutationsHelper(base []int, validValues []int, permutationcount int, results chan []int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup

	for i := 0; i < permutationcount; i++ {
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

func listPermutationsHelperRange(base []int, validValues []int, permutationcount int, results chan []int, routineCount int) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup

	for i := 0; i < routineCount-1; i++ {
		wg.Add(1)
		select {
		case sem <- struct{}{}:
			go incrementRangeWGSem(base, validValues, [2]int{(permutationcount / routineCount) * i, (permutationcount / routineCount) * (i + 1)}, results, &wg, sem)
		default:
			incrementRangeWGSem(base, validValues, [2]int{(permutationcount / routineCount) * i, (permutationcount / routineCount) * (i + 1)}, results, &wg, sem)
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
