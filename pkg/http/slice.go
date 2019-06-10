package httplib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// SliceConfigUnanimous contains parameters to input to GenSliceUnanimous (Both int and float64)
// ValidValues is []int and applies for all values of the returning slice
type SliceConfigUnanimous struct {
	Dimensions       []int
	ValidValues      []int
	PermutationRange [2]int
	APIURL           string
}

// SliceConfigUnique contains parameters to input to GenSliceUnique (Both int and float64)
// using ValidValues and ValidValuesIndex, the validValues for each value in the slice can differ
type SliceConfigUnique struct {
	Dimensions       []int
	ValidValues      []int
	ValidValuesIndex []int
	PermutationRange [2]int
	APIURL           string
}

// GenIntSliceUnanimous returns a 2D int slice that contains int slices that can be reshaped into the multidimensional slice described by dimensions
func GenIntSliceUnanimous(config SliceConfigUnanimous) ([][]int, error) {
	stringDimensions := strconv.Itoa(config.Dimensions[0])
	for _, dimensionVal := range config.Dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(config.ValidValues[0])
	for _, validValueVal := range config.ValidValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values_unanimous=%v&permutation_range=%v,%v", config.APIURL, stringDimensions, stringValidValues, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}

	ret := [][]int{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		temp := []int{}
		for _, valueInterface := range permutationInterface.([]interface{}) {
			temp = append(temp, int(valueInterface.(float64)))
		}
		ret = append(ret, temp)
	}
	return ret, nil
}

// GenIntSliceUnique returns a 2D int slice that contains int slices that can be reshaped into the multidimensional slice described by dimensions
func GenIntSliceUnique(config SliceConfigUnique) ([][]int, error) {
	stringDimensions := strconv.Itoa(config.Dimensions[0])
	for _, dimensionVal := range config.Dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(config.ValidValues[0])
	for _, validValueVal := range config.ValidValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	stringValidValuesIndex := strconv.Itoa(config.ValidValuesIndex[0])
	for _, validValueVal := range config.ValidValuesIndex[1:] {
		stringValidValuesIndex = fmt.Sprintf("%v,%v", stringValidValuesIndex, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values=%v&valid_values_index=%v&permutation_range=%v,%v", config.APIURL, stringDimensions, stringValidValues, stringValidValuesIndex, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}
	fmt.Println(string(body))
	ret := [][]int{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		temp := []int{}
		for _, valueInterface := range permutationInterface.([]interface{}) {
			temp = append(temp, int(valueInterface.(float64)))
		}
		ret = append(ret, temp)
	}
	return ret, nil
}

// GenFloatSliceUnanimous returns a 2D int slice that contains int slices that can be reshaped into the multidimensional slice described by dimensions
func GenFloatSliceUnanimous(config SliceConfigUnanimous) ([][]float64, error) {
	stringDimensions := strconv.Itoa(config.Dimensions[0])
	for _, dimensionVal := range config.Dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(config.ValidValues[0])
	for _, validValueVal := range config.ValidValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values_unanimous=%v&permutation_range=%v,%v", config.APIURL, stringDimensions, stringValidValues, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}

	ret := [][]float64{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		temp := []float64{}
		for _, valueInterface := range permutationInterface.([]interface{}) {
			temp = append(temp, valueInterface.(float64))
		}
		ret = append(ret, temp)
	}
	return ret, nil
}

// GenFloatSliceUnique returns a 2D int slice that contains int slices that can be reshaped into the multidimensional slice described by dimensions
func GenFloatSliceUnique(config SliceConfigUnique) ([][]float64, error) {
	stringDimensions := strconv.Itoa(config.Dimensions[0])
	for _, dimensionVal := range config.Dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(config.ValidValues[0])
	for _, validValueVal := range config.ValidValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	stringValidValuesIndex := strconv.Itoa(config.ValidValuesIndex[0])
	for _, validValueVal := range config.ValidValuesIndex[1:] {
		stringValidValuesIndex = fmt.Sprintf("%v,%v", stringValidValuesIndex, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values=%v&valid_values_index=%v&permutation_range=%v,%v", config.APIURL, stringDimensions, stringValidValues, stringValidValuesIndex, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}
	fmt.Println(string(body))
	ret := [][]float64{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		temp := []float64{}
		for _, valueInterface := range permutationInterface.([]interface{}) {
			temp = append(temp, valueInterface.(float64))
		}
		ret = append(ret, temp)
	}
	return ret, nil
}

// ReshapeIntSlice Tries to reshape the single int slice into a multidimensional slice
func ReshapeIntSlice(dimensions []int, flatSlice []int) (interface{}, error) {
	var ret interface{}
	switch len(dimensions) {
	case 1:
		ret = flatSlice
	case 2:
		temp := [][]int{}
		for i := 0; i < dimensions[0]; i++ {
			temp = append(temp, flatSlice[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 3:
		temp := [][][]int{}
		for i := 0; i < dimensions[0]; i++ {
			doubleTemp := [][]int{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp = append(doubleTemp, flatSlice[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, doubleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	case 4:
		temp := [][][][]int{}
		for i := 0; i < dimensions[0]; i++ {
			tripleTemp := [][][]int{}
			for i := 0; i < dimensions[0]*dimensions[1]; i++ {
				doubleTemp := [][]int{}
				for i := 0; i < dimensions[0]*dimensions[1]*dimensions[2]; i++ {
					doubleTemp = append(doubleTemp, flatSlice[i*dimensions[3]:(i+1)*dimensions[3]])
				}
				tripleTemp = append(tripleTemp, doubleTemp[i*dimensions[2]:(i+1)*dimensions[2]])
			}
			temp = append(temp, tripleTemp[i*dimensions[1]:(i+1)*dimensions[1]])
		}
		ret = temp
	default:
		return nil, errors.New("Can not reshape to the number of dimensions requested")
	}
	return ret, nil
}
