package httpmainlib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPStringSliceParams extends StringSliceParams for HTTP parameter insertion
type HTTPStringSliceParams struct {
	*genlib.StringSliceParams
	StringValuesRaw   []string
	StringValuesIndex []int
}

func init() {
	HTTPParams["string_slice"] = &HTTPStringSliceParams{&genlib.StringSliceParams{}, []string{}, []int{}}
}

// SetParams allows for setting of parameters in HTTPStringSliceParams
func (curParams *HTTPStringSliceParams) SetParams(k string, v string) error {
	switch k {
	case "dimensions":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Dimensions = append(curParams.Dimensions, intV)
			}
		}
	case "lengths":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Lengths = append(curParams.Lengths, intV)
			}
		}
	case "string_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			curParams.StringValuesRaw = append(curParams.StringValuesRaw, v)
		}
	case "string_values_index":
		strList := strings.Split(v, ",")
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

// New returns an empty object of the same config type
func (curParams *HTTPStringSliceParams) New() DataGenHTTP {
	return &HTTPStringSliceParams{&genlib.StringSliceParams{}, []string{}, []int{}}
}
