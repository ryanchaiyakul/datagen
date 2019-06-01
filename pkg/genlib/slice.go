package genlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GenIntSlice returns a single int slice that can be reshaped into the one described by dimensions
func GenIntSlice(dimensions []int, validValues []int, permutation int, apiURL string) ([]int, error) {
	stringDimensions := strconv.Itoa(dimensions[0])
	for _, dimensionVal := range dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(validValues[0])
	for _, validValueVal := range validValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values=%v&permutation_range=%v,%v", apiURL, stringDimensions, stringValidValues, permutation, permutation+1))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}

	ret := []int{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		for _, valueInterface := range permutationInterface.([]interface{}) {
			ret = append(ret, int(valueInterface.(float64)))
		}
	}
	return ret, nil
}

// GenFloatSlice returns a single float slice that can be reshaped into the one described by dimensions
func GenFloatSlice(dimensions []int, validValues []int, permutation int, apiURL string) ([]float64, error) {
	stringDimensions := strconv.Itoa(dimensions[0])
	for _, dimensionVal := range dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v,%v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(validValues[0])
	for _, validValueVal := range validValues[1:] {
		stringValidValues = fmt.Sprintf("%v,%v", stringValidValues, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?funcid=slice&dimensions=%v&valid_values=%v&permutation_range=%v,%v", apiURL, stringDimensions, stringValidValues, permutation, permutation+1))
	if err != nil {
		return nil, errors.New("genlib : invalid URL")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}

	ret := []float64{}
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		for _, valueInterface := range permutationInterface.([]interface{}) {
			ret = append(ret, valueInterface.(float64))
		}
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
