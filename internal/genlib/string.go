package genlibinternal

// GenString returns a list of string permutations that are valid to the asciiRange
// if asciiRange blank, uses all lowercase and uppercase letters
func GenString(length int, asciiRange [2]int, permutationRange [2]int) ([]string, error) {
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

	retList := []string{}

	intList, err := GenArray([]int{length}, validValues, permutationRange)
	if err == nil {
		for k := range intList.([][]int) {
			ret := ""
			for _, v := range intList.([][]int)[k] {
				ret += string(v)
			}
			retList = append(retList, ret)
		}
		return retList, err
	}
	return nil, err
}
