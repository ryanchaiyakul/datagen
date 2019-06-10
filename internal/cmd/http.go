package cmdlib

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
	helperlib "github.com/ryanchaiyakul/datagen/internal/helper"
)

type arrayParams struct {
	dimensions       []int
	validValues      [][]int
	permutationRange [2]int
}

type stringParams struct {
	length           int
	asciiValues      [][]int
	permutationRange [2]int
}

// Handler is the HTTP Rest API function to be passed into net/http
func Handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	var ret interface{}

	if funcid, ok := r.URL.Query()["funcid"]; ok {
		switch funcid[0] {
		case "slice":
			ret = getSlice(w, r)
		case "string":
			ret = getString(w, r)
		default:
			http.Error(w, "400 incorrect funcid", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "400 funcid not passed", http.StatusBadRequest)
		return
	}

	if retEncoded, err := json.Marshal(ret); err == nil && retEncoded != nil {
		w.Write(retEncoded)
	} else {
		http.Error(w, "400 paramters are incorrect", http.StatusBadRequest)
	}
}

func getSlice(w http.ResponseWriter, r *http.Request) interface{} {
	// initialize config variables
	curParams := arrayParams{nil, nil, [2]int{0, 0}}

	// different validValue configurations
	validValuesUnanimous := []int{}
	validValues := []int{}
	validValuesIndex := []int{}
	validValuesRange := [2]int{}

	for k, v := range r.URL.Query() {
		switch k {
		case "dimensions":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					curParams.dimensions = append(curParams.dimensions, intV)
				}
			}

		// Different ways to configure validValues

		// Unique configuration
		case "valid_values":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					validValues = append(validValues, intV)
				}
			}
		case "valid_values_index":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					validValuesIndex = append(validValuesIndex, intV)
				}
			}

		// Config duplicated for all integers in the slice
		case "valid_values_unanimous":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					validValuesUnanimous = append(validValuesUnanimous, intV)
				}
			}

		// 2 long int array representing the bounds of the validValues
		case "valid_values_range":
			strList := strings.Split(v[0], ",")
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					validValuesRange[k] = intV
				}
			}

		case "permutation_range":
			strList := strings.Split(v[0], ",")
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					curParams.permutationRange[k] = intV
				}
			}
		default:
			if k != "funcid" {
				http.Error(w, "400 unknown parameter.", http.StatusBadRequest)
				return nil
			}
		}
	}

	if curParams.dimensions == nil {
		http.Error(w, "400 missing dimensions.", http.StatusBadRequest)
		return nil
	}

	// validValue configuration
	if len(validValuesUnanimous) != 0 {
		for i := 0; i < len(helperlib.FlatSlice(curParams.dimensions, 0)); i++ {
			curParams.validValues = append(curParams.validValues, validValuesUnanimous)
		}
	} else if len(validValues) != 0 && len(validValuesIndex) != 0 {
		tempVal := 0
		for i := 0; i < len(validValuesIndex); i++ {
			temp := []int{}
			for j := 0; j < validValuesIndex[i]; j++ {
				temp = append(temp, validValues[tempVal])
				tempVal++
			}
			curParams.validValues = append(curParams.validValues, temp)
		}
	} else if validValuesRange[0] != 0 && validValuesRange[1] != 0 {
		tempSlice := []int{}
		for i := validValuesRange[0]; i < validValuesRange[1]; i++ {
			tempSlice = append(tempSlice, i)
		}
		for i := 0; i < len(helperlib.FlatSlice(curParams.dimensions, 0)); i++ {
			curParams.validValues = append(curParams.validValues, tempSlice)
		}
	} else {
		http.Error(w, "400 validValues not properly passed", http.StatusBadRequest)
		return nil
	}

	// gets slice(s) and returns it back to Handler
	slicelibRet, err := genlib.GenArray(curParams.dimensions, curParams.validValues, curParams.permutationRange)
	if err == nil {
		return slicelibRet
	}
	return nil
}

func getString(w http.ResponseWriter, r *http.Request) interface{} {
	// initialize config variables
	curParams := stringParams{0, [][]int{}, [2]int{}}

	// different asciiValues Configuration
	asciiValuesUnanimous := []int{}
	asciiValues := []int{}
	asciiValuesIndex := []int{}
	asciiValuesRange := [2]int{}

	for k, v := range r.URL.Query() {
		switch k {
		case "length":
			intLength, err := strconv.Atoi(string(v[0][0]))
			if err != nil {
				http.Error(w, "400 can not convert length to a string.", http.StatusBadRequest)
				return nil
			}
			curParams.length = intLength

		// Different ways to configure asciiValues

		// Unique configuration
		case "ascii_values":
			strList := strings.Split(v[0], ",")
			intList := []int{}
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					intList = append(intList, intV)
				}
			}
			asciiValues = intList
		case "ascii_values_index":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					asciiValuesIndex = append(asciiValuesIndex, intV)
				}
			}

		// Config duplicated to all characters in the string
		case "ascii_values_unanimous":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					asciiValuesUnanimous = append(asciiValuesUnanimous, intV)
				}
			}

		// 2 long int array representing the bounds for asciiValues
		case "ascii_values_range":
			strList := strings.Split(v[0], ",")
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					asciiValuesRange[k] = intV
				}
			}

		case "permutation_range":
			strList := strings.Split(v[0], ",")
			intList := [2]int{}
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					intList[k] = intV
				}
			}
			curParams.permutationRange = intList
		default:
			if k != "funcid" {
				http.Error(w, "400 unknown parameter.", http.StatusBadRequest)
				return nil
			}
		}
	}
	if curParams.length == 0 {
		http.Error(w, "400 missing length.", http.StatusBadRequest)
		return nil
	}

	// asciiValues configuration
	if len(asciiValuesUnanimous) != 0 {
		for i := 0; i < len(helperlib.FlatSlice([]int{curParams.length}, 0)); i++ {
			curParams.asciiValues = append(curParams.asciiValues, asciiValuesUnanimous)
		}
	} else if len(asciiValues) != 0 && len(asciiValuesIndex) != 0 {
		tempVal := 0
		for i := 0; i < len(asciiValuesIndex); i++ {
			temp := []int{}
			for j := 0; j < asciiValuesIndex[i]; j++ {
				temp = append(temp, asciiValues[tempVal])
				tempVal++
			}
			curParams.asciiValues = append(curParams.asciiValues, temp)
		}
	} else if asciiValuesRange[0] != 0 && asciiValuesRange[1] != 0 {
		tempSlice := []int{}
		for i := asciiValuesRange[0]; i < asciiValuesRange[1]; i++ {
			tempSlice = append(tempSlice, i)
		}
		for i := 0; i < len(helperlib.FlatSlice([]int{curParams.length}, 0)); i++ {
			curParams.asciiValues = append(curParams.asciiValues, tempSlice)
		}
	} else {
		http.Error(w, "400 asciiValues not properly passed", http.StatusBadRequest)
		return nil
	}

	// gets string(s) and returns them to Handler
	stringlibRet, err := genlib.GenString(curParams.length, curParams.asciiValues, curParams.permutationRange)
	if err == nil {
		return stringlibRet
	}
	return nil
}
