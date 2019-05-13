package datagen

import "testing"

// CheckOneDimensionalIntChantoList raises a testing error if the list is not found within the list
func CheckOneDimensionalIntChantoList(inputListChan chan []int, checkList [][]int, t *testing.T) {
	for output := range inputListChan {
		equal, checkID := checkOneDimensionalInttoList(output, checkList)
		if equal {
			if checkID >= len(checkList) {
				checkList = checkList[:len(checkList)-2]
			} else {
				checkList = append(checkList[:checkID], checkList[checkID+1:]...)
			}
		} else {
			t.Errorf("[]int %d not found in [][]int, %d", output, checkList)
		}
	}
	if len(checkList) != 0 {
		t.Errorf("Not all values found in %d", checkList)
	}
}

// CheckOneDimensionalInttoList raises a testing error if the list is not found within the list
func CheckOneDimensionalInttoList(listOne []int, checkList [][]int, t *testing.T) {
	equal, _ := checkOneDimensionalInttoList(listOne, checkList)
	if !equal {
		t.Errorf("[]int %d not found in [][]int, %d", listOne, checkList)
	}
}

func checkOneDimensionalInttoList(listOne []int, listCheck [][]int) (bool, int) {
	for i := 0; i < len(listCheck); i++ {
		equal := checkOneDimensionalInt(listOne, listCheck[i])
		if equal {
			return true, i
		}
	}
	return false, 0
}

func checkOneDimensionalInt(listOne []int, listTwo []int) bool {
	for i := 0; i < len(listOne); i++ {
		if listOne[i] != listTwo[i] {
			return false
		}
	}
	return true
}
