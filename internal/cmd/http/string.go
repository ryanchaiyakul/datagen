package httplib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPStringParams extends StringParams for HTTP parameter insertion
type HTTPStringParams struct {
	*genlib.StringParams
	StringValuesRaw   []string
	StringValuesIndex []int
}

func init() {
	HTTPParams["string"] = &HTTPStringParams{&genlib.StringParams{}, []string{}, []int{}}
}

// SetParams allows for setting of parameters in HTTPStringParams
func (curParams *HTTPStringParams) SetParams(k string, v string) error {
	switch k {
	case "length":
		if len(v) == 1 {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Length = intV
			}
		} else {
			return fmt.Errorf("getStringParams : length : %v has more than 1 integer", v)
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
func (curParams *HTTPStringParams) New() DataGenHTTP {
	return &HTTPStringParams{&genlib.StringParams{}, []string{}, []int{}}
}
