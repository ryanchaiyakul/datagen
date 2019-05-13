package datagen

// GenOneDimensional sends one dimensional arrays that exhaustively goes through all the possible permutations into results channel
func GenOneDimensional(axisOneLength int, validValues []int, results chan []int) {
	temp := []int{}
	// initialize array with the first element as that is the first possible variation
	for i := 0; i < axisOneLength; i++ {
		temp = append(temp, validValues[0])
	}
	results <- append(temp[:0:0], temp...)

	for {
		temp = incrementOneLeft(temp, validValues, 0)
		if temp == nil {
			break
		}
		results <- append(temp[:0:0], temp...)
	}
	close(results)
}

// GenOneDimensionalRet returns a 2D array that contains all valid 1 dimensional permutations
func GenOneDimensionalRet(axisOneLength int, validValues []int) [][]int {
	ret := [][]int{}
	temp := []int{}

	for i := 0; i < axisOneLength; i++ {
		temp = append(temp, validValues[0])
	}
	ret = append(ret, append([]int(nil), temp...))

	for {
		temp = incrementOneLeft(temp, validValues, 0)
		if temp == nil {
			break
		}
		ret = append(ret, append([]int(nil), temp...))
	}
	return ret
}

func incrementOneLeft(intList []int, validValues []int, startingval int) []int {
	if startingval == len(intList) {
		return nil
	}
	currentIndex := 1 + sliceIndex(len(validValues), func(sliceI int) bool { return validValues[sliceI] == intList[startingval] })
	if currentIndex < len(validValues) {
		intList[startingval] = validValues[currentIndex]
		return intList
	}
	intList[startingval] = validValues[0]
	return incrementOneLeft(intList, validValues, startingval+1)
}

func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
