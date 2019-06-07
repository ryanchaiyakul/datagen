package genlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GenString returns a string permutation that is equal to the length passed in
// if asciiRange is {0,0}, all upper and lowercase letters will be used
func GenString(length int, asciiRange [2]int, permutation int, apiURL string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%v?funcid=string&length=%v&ascii_range=%v&permutation_range=%v,%v", apiURL, length, fmt.Sprintf("%v,%v", asciiRange[0], asciiRange[1]), permutation, permutation+1))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if body == nil {
		return "", errors.New("genlib : invalid dimension paramter or permutation out of range")
	}

	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	return tempInterface.([]interface{})[0].(string), nil
}

// GenStringRange returns a string permutation that is equal to the length passed in
// if asciiRange is {0,0}, all upper and lowercase letters will be used
func GenStringRange(length int, asciiRange [2]int, permutation [2]int, apiURL string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%v?funcid=string&length=%v&ascii_range=%v&permutation_range=%v,%v", apiURL, length, fmt.Sprintf("%v,%v", asciiRange[0], asciiRange[1]), permutation[0], permutation[1]))
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

	var ret []string
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		ret = append(ret, permutationInterface.(string))
	}
	return ret, nil
}
