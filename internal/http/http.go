package httpapilib

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	slicelib "github.com/ryanchaiyakul/datagen/internal/slice"
	stringlib "github.com/ryanchaiyakul/datagen/internal/string"
)

type arrayParams struct {
	dimensions       []int
	validValues      []int
	permutationRange [2]int
}

type stringParams struct {
	length           int
	asciiRange       [2]int
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
	switch r.URL.Query()["funcid"][0] {
	case "array":
		curParams := arrayParams{nil, nil, [2]int{0, 0}}
		for k, v := range r.URL.Query() {
			switch k {
			case "dimensions":
				strList := strings.Split(v[0], ",")
				intList := []int{}
				for _, v := range strList {
					if intV, err := strconv.Atoi(v); err == nil {
						intList = append(intList, intV)
					}
				}
				curParams.dimensions = intList
			case "valid_values":
				strList := strings.Split(v[0], ",")
				intList := []int{}
				for _, v := range strList {
					if intV, err := strconv.Atoi(v); err == nil {
						intList = append(intList, intV)
					}
				}
				curParams.validValues = intList
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
					return
				}
			}
		}

		if curParams.dimensions == nil {
			http.Error(w, "400 missing dimensions.", http.StatusBadRequest)
			return
		}
		slicelibRet, err := slicelib.GenArray(curParams.dimensions, curParams.validValues, curParams.permutationRange)
		if err == nil {
			ret = slicelibRet
		}

	case "string":
		curParams := stringParams{0, [2]int{0, 0}, [2]int{0, 0}}
		for k, v := range r.URL.Query() {
			switch k {
			case "length":
				intLength, err := strconv.Atoi(string(v[0][0]))
				if err != nil {
					http.Error(w, "400 can not convert length to a string.", http.StatusBadRequest)
					return
				}
				curParams.length = intLength
			case "ascii_range":
				strList := strings.Split(v[0], ",")
				intList := [2]int{}
				for k, v := range strList {
					if intV, err := strconv.Atoi(v); err == nil {
						intList[k] = intV
					}
				}
				curParams.asciiRange = intList
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
					return
				}
			}
		}
		if curParams.length == 0 {
			http.Error(w, "400 missing length.", http.StatusBadRequest)
			return
		}
		stringlibRet, err := stringlib.GenString(curParams.length, curParams.asciiRange, curParams.permutationRange)
		if err == nil {
			ret = stringlibRet
		}
	}

	if retEncoded, err := json.Marshal(ret); err == nil {
		w.Write(retEncoded)
	}
}
