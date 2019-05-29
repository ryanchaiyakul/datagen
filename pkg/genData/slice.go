package genlib

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GenArray(dimensions []int, validValues []int, permutation int, apiURL string) (interface{}, error) {
	stringDimensions := strconv.Itoa(dimensions[0])
	for _, dimensionVal := range dimensions[1:] {
		stringDimensions = fmt.Sprintf("%v, %v", stringDimensions, strconv.Itoa(dimensionVal))
	}
	stringValidValues := strconv.Itoa(validValues[0])
	for _, validValueVal := range validValues[1:] {
		stringValidValues = fmt.Sprintf("%v, %v", stringValidValues, strconv.Itoa(validValueVal))
	}
	resp, err := http.Get(fmt.Sprintf("%v?dimensions=%v&valid_values=%v&permutation_range=%v, %v", apiURL, stringDimensions, stringValidValues, permutation, permutation+1))
	if err != nil {
		return nil, errors.New("genlib : invalid URL")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if body == nil {
		return nil, errors.New("genlib : invalid dimension paramter or permutation out of range")
	}
}
