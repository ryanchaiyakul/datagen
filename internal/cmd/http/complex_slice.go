package httpmainlib

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPComplexSliceParams extends ComplexSliceParams for HTTP parameter insertion
type HTTPComplexSliceParams struct {
	*genlib.ComplexSliceParams
	RealValues       []int
	ImaginaryValues  []int
	ValidValuesIndex []int
}

// SetParams allows for setting of parameters in HTTPComplexSliceParams
func (curParams *HTTPComplexSliceParams) SetParams(k string, v string) error {
	switch k {
	case "dimensions":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.Dimensions = append(curParams.Dimensions, intV)
			}
		}
	case "real_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.RealValues = append(curParams.RealValues, intV)
			}
		}
	case "imaginary_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if intV, err := strconv.Atoi(v); err == nil {
				curParams.ImaginaryValues = append(curParams.ImaginaryValues, intV)
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
			return fmt.Errorf("getComplexParam : unknown parameter : %v", k)
		}
	}
	if len(curParams.RealValues) != 0 && len(curParams.ImaginaryValues) != 0 && len(curParams.ValidValuesIndex) != 0 && len(curParams.ValidValues) == 0 {
		indexLength := 1
		for _, v := range curParams.ValidValuesIndex {
			indexLength *= v
		}

		if len(curParams.RealValues) != len(curParams.ImaginaryValues) || len(curParams.RealValues) != indexLength {
			return fmt.Errorf("getComplexParam : mismatched real_values : %v or imaginary_values : %v or index : %v", curParams.RealValues, curParams.ImaginaryValues, curParams.ValidValuesIndex)
		}

		index := 0
		for _, v := range curParams.ValidValuesIndex {
			tempValues := []complex128{}
			for i := 0; i < v; i++ {
				tempValues = append(tempValues, complex(float64(curParams.RealValues[index]), float64(curParams.ImaginaryValues[index])))
				index++
			}
			curParams.ValidValues = append(curParams.ValidValues, tempValues)
		}
	}
	return nil
}
