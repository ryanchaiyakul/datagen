package httplib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// StringConfigUnanimous contains parameters to input to GenString
type StringConfigUnanimous struct {
	Length           int
	ASCIIValues      []int
	PermutationRange [2]int
	APIURL           string
}

// StringConfigUnique contains parameters to input to GenString
type StringConfigUnique struct {
	Length           int
	ASCIIValues      []int
	ASCIIValuesIndex []int
	PermutationRange [2]int
	APIURL           string
}

// GenStringUnanimous returns a []string containing permutations that is equal to the length passed in
// if asciiValues is {}, all upper and lowercase letters will be used
func GenStringUnanimous(config StringConfigUnanimous) ([]string, error) {
	stringASCIIValues := ""
	if config.ASCIIValues[0] == -1 {
		config.ASCIIValues[0] = 65
		for i := 66; i < 91; i++ {
			config.ASCIIValues = append(config.ASCIIValues, i)
		}
		for i := 97; i < 123; i++ {
			config.ASCIIValues = append(config.ASCIIValues, i)
		}
	}
	stringASCIIValues = strconv.Itoa(config.ASCIIValues[0])
	for _, asciiValuesVal := range config.ASCIIValues[1:] {
		stringASCIIValues = fmt.Sprintf("%v,%v", stringASCIIValues, strconv.Itoa(asciiValuesVal))
	}

	resp, err := http.Get(fmt.Sprintf("%v?funcid=string&length=%v&ascii_values_unanimous=%v&permutation_range=%v,%v", config.APIURL, config.Length, stringASCIIValues, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("GenStringUnanimous : invalid length paramter or permutation out of range")
	}

	var ret []string
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		ret = append(ret, permutationInterface.(string))
	}
	return ret, nil
}

// GenStringUnique returns a []string containing permutations that is equal to the length passed in
// if asciiValues is {}, all upper and lowercase letters will be used
func GenStringUnique(config StringConfigUnique) ([]string, error) {
	index := -1
	for i := 0; i < len(config.ASCIIValuesIndex); i++ {
		index += config.ASCIIValuesIndex[i]
		if config.ASCIIValuesIndex[i] == 1 && config.ASCIIValues[index] == -1 {
			config.ASCIIValues[index] = 65
			for j := 122; j > 96; j-- {
				config.ASCIIValues = append(config.ASCIIValues, 0)
				copy(config.ASCIIValues[index+2:], config.ASCIIValues[index+1:])
				config.ASCIIValues[index+1] = j
			}
			for j := 90; j > 65; j-- {
				config.ASCIIValues = append(config.ASCIIValues, 0)
				copy(config.ASCIIValues[index+2:], config.ASCIIValues[index+1:])
				config.ASCIIValues[index+1] = j
			}
			index += 51
			config.ASCIIValuesIndex[i] += 51
		}
	}
	stringASCIIValues := strconv.Itoa(config.ASCIIValues[0])
	for _, asciiValueVal := range config.ASCIIValues[1:] {
		stringASCIIValues = fmt.Sprintf("%v,%v", stringASCIIValues, strconv.Itoa(asciiValueVal))
	}
	stringASCIIValuesIndex := strconv.Itoa(config.ASCIIValuesIndex[0])
	for _, asciiValueVal := range config.ASCIIValuesIndex[1:] {
		stringASCIIValuesIndex = fmt.Sprintf("%v,%v", stringASCIIValuesIndex, strconv.Itoa(asciiValueVal))
	}

	resp, err := http.Get(fmt.Sprintf("%v?funcid=string&length=%v&ascii_values=%v&ascii_values_index=%v&permutation_range=%v,%v", config.APIURL, config.Length, stringASCIIValues, stringASCIIValuesIndex, config.PermutationRange[0], config.PermutationRange[1]))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if body == nil || string(body) == "null" {
		return nil, errors.New("GenStringUnique : invalid length paramter or permutation out of range")
	}
	var ret []string
	var tempInterface interface{}
	json.Unmarshal(body, &tempInterface)
	for _, permutationInterface := range tempInterface.([]interface{}) {
		ret = append(ret, permutationInterface.(string))
	}
	return ret, nil
}
