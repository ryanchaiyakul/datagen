package cmdlib

import (
	"encoding/json"
	"errors"
	"fmt"
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
	stringValues     [][]string
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
			ret = getSlice(r)
		case "string":
			ret = getString(r)
		default:
			ret = errors.New("Handler : unknown funcid")
		}
	} else {
		ret = errors.New("Handler : funcid not passed")
	}

	switch ret.(type) {
	case error:
		w.Write([]byte(fmt.Sprint(ret)))
		return
	}
	if retEncoded, err := json.Marshal(ret); err == nil {
		w.Write(retEncoded)
	}
}

func getSlice(r *http.Request) interface{} {
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
				return errors.New("GetSlice : unknown parameter")
			}
		}
	}

	if curParams.dimensions == nil {
		return errors.New("GetSlice : missing dimensions")
	}

	// validValue configuration
	if len(validValuesUnanimous) != 0 {
		for i := 0; i < helperlib.FlatSliceLength(curParams.dimensions); i++ {
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
	} else if validValuesRange[0] != 0 || validValuesRange[1] != 0 {
		tempSlice := []int{}
		for i := validValuesRange[0]; i < validValuesRange[1]+1; i++ {
			tempSlice = append(tempSlice, i)
		}
		for i := 0; i < helperlib.FlatSliceLength(curParams.dimensions); i++ {
			curParams.validValues = append(curParams.validValues, tempSlice)
		}
	} else {
		return errors.New("GetSlice : missing validValues")
	}

	// gets slice(s) and returns it back to Handler
	sliceRet, err := genlib.GenArray(curParams.dimensions, curParams.validValues, curParams.permutationRange)
	if err == nil {
		return sliceRet
	}
	return err
}

func getString(r *http.Request) interface{} {
	// initialize config variables
	curParams := stringParams{0, [][]string{}, [2]int{}}

	// different stringValues Configuration
	stringValuesUnanimous := []string{}
	stringValues := []string{}
	stringValuesIndex := []int{}

	for k, v := range r.URL.Query() {
		switch k {
		case "length":
			intLength, err := strconv.Atoi(string(v[0][0]))
			if err != nil {
				return errors.New("GetString : strconv failed when called on length")
			}
			curParams.length = intLength

		// Different ways to configure stringValues

		// Unique configuration
		case "string_values":
			stringValues = strings.Split(v[0], ",")

		case "string_values_index":
			strList := strings.Split(v[0], ",")
			for _, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					stringValuesIndex = append(stringValuesIndex, intV)
				}
			}

		// Config duplicated to all characters in the string
		case "string_values_unanimous":
			stringValuesUnanimous = strings.Split(v[0], ",")

		case "permutation_range":
			strList := strings.Split(v[0], ",")
			if len(strList) == 2 {
				for k, v := range strList {
					if intV, err := strconv.Atoi(v); err == nil {
						curParams.permutationRange[k] = intV
					}
				}
			}
		default:
			if k != "funcid" {
				return errors.New("GetString : unknown paramter")
			}
		}
	}
	if curParams.length == 0 {
		return errors.New("GetString : missing length")
	}

	// stringValues configuration
	if len(stringValuesUnanimous) != 0 {
		for i := 0; i < helperlib.FlatSliceLength([]int{curParams.length}); i++ {
			curParams.stringValues = append(curParams.stringValues, stringValuesUnanimous)
		}
	} else if len(stringValues) != 0 && len(stringValuesIndex) != 0 {
		tempVal := 0
		for i := 0; i < len(stringValuesIndex); i++ {
			tempSlice := []string{}
			for j := 0; j < stringValuesIndex[i]; j++ {
				tempSlice = append(tempSlice, stringValues[tempVal])
				tempVal++
			}
			curParams.stringValues = append(curParams.stringValues, tempSlice)
		}
	} else {
		return errors.New("GetString : missing stringValues")
	}

	// gets string(s) and returns them to Handler
	stringRet, err := genlib.GenString(curParams.length, curParams.stringValues, curParams.permutationRange)
	if err == nil {
		return stringRet
	}
	return err
}
