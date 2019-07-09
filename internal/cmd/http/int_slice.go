package httplib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPIntSliceParams extends IntSliceParams for HTTP parameter insertion
type HTTPIntSliceParams struct {
	*genlib.IntSliceParams
	ValidValuesRaw   []int
	ValidValuesIndex []int
}

func init() {
	HTTPParams["int_slice"] = &HTTPIntSliceParams{&genlib.IntSliceParams{}, []int{}, []int{}}
}

// SetParams allows for setting of parameters in HTTPIntSliceParams
func (curParams *HTTPIntSliceParams) SetParams(k string, v string) error {
	switch k {
	case "dimensions":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Dimensions = append(curParams.Dimensions, intV)
			}
		}
	case "valid_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.ValidValuesRaw = append(curParams.ValidValuesRaw, intV)
			}
		}
	case "valid_values_index":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.ValidValuesIndex = append(curParams.ValidValuesIndex, intV)
			}
		}
	default:
		if k != "funcid" {
			return fmt.Errorf("getSliceParam : unknown parameter : %v", k)
		}
	}

	if len(curParams.ValidValuesRaw) != 0 && len(curParams.ValidValuesIndex) != 0 && len(curParams.ValidValues) == 0 {
		validValuesCount := 0
		for _, v := range curParams.ValidValuesIndex {
			validValuesCount += v
		}
		if validValuesCount != len(curParams.ValidValuesRaw) {
			return fmt.Errorf("getSliceParam : mismatched ValidValuesIndex : %v and ValidValuesRaw : %v", curParams.ValidValuesIndex, curParams.ValidValuesRaw)
		}
		tempVal := 0
		for i := 0; i < len(curParams.ValidValuesIndex); i++ {
			tempSlice := []int{}
			for j := 0; j < curParams.ValidValuesIndex[i]; j++ {
				tempSlice = append(tempSlice, curParams.ValidValuesRaw[tempVal])
				tempVal++
			}
			curParams.ValidValues = append(curParams.ValidValues, tempSlice)
		}
	}
	return nil
}

// New returns an empty object of the same config type
func (curParams *HTTPIntSliceParams) New() DataGenHTTP {
	return &HTTPIntSliceParams{&genlib.IntSliceParams{}, []int{}, []int{}}
}
