package genlib

// GenString returns a list of string permutations that are valid to the asciiRange
// if asciiRange blank, uses all lowercase and uppercase letters
func GenString(length int, asciiValues [][]int, permutationRange [2]int) ([]string, error) {
	retList := []string{}
	intList, err := GenArray([]int{length}, asciiValues, permutationRange)
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
