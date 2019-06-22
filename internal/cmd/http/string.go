package httpmainlib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

func getStringParam(k string, v []string, curParams *genlib.StringParams) error {
	switch k {
	case "length":
		if len(v) == 1 {
			if intV, err := strconv.Atoi(v[0]); err == nil {
				curParams.Length = intV
			}
		} else {
			return fmt.Errorf("getStringParams : length : %v has more than 1 integer", v)
		}
	case "string_values":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			curParams.StringValuesRaw = append(curParams.StringValuesRaw, v)
		}
	case "string_values_index":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.StringValuesIndex = append(curParams.StringValuesIndex, intV)
			}
		}
	default:
		if k != "funcid" {
			return fmt.Errorf("getStringParams : unknown parameter : %v", k)
		}
	}

	if len(curParams.StringValuesRaw) != 0 && len(curParams.StringValuesIndex) != 0 && len(curParams.StringValues) == 0 {
		tempVal := 0
		for i := 0; i < len(curParams.StringValuesIndex); i++ {
			tempSlice := []string{}
			for j := 0; j < curParams.StringValuesIndex[i]; j++ {
				tempSlice = append(tempSlice, curParams.StringValuesRaw[tempVal])
				tempVal++
			}
			curParams.StringValues = append(curParams.StringValues, tempSlice)
		}
	}
	return nil
}

func getStringSliceParam(k string, v []string, curParams *genlib.StringSliceParams) error {
	switch k {
	case "dimensions":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Dimensions = append(curParams.Dimensions, intV)
			}
		}
	case "lengths":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Lengths = append(curParams.Lengths, intV)
			}
		}
	case "string_values":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			curParams.StringValuesRaw = append(curParams.StringValuesRaw, v)
		}
	case "string_values_index":
		strList := strings.Split(v[0], ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.StringValuesIndex = append(curParams.StringValuesIndex, intV)
			}
		}
	default:
		if k != "funcid" {
			return fmt.Errorf("getStringParams : unknown parameter : %v", k)
		}
	}

	if len(curParams.StringValuesRaw) != 0 && len(curParams.StringValuesIndex) != 0 && len(curParams.StringValues) == 0 {
		tempVal := 0
		for i := 0; i < len(curParams.StringValuesIndex); i++ {
			tempSlice := []string{}
			for j := 0; j < curParams.StringValuesIndex[i]; j++ {
				tempSlice = append(tempSlice, curParams.StringValuesRaw[tempVal])
				tempVal++
			}
			curParams.StringValues = append(curParams.StringValues, tempSlice)
		}
	}
	return nil
}
