package httpmainlib

import (
	"fmt"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// HTTPComplexParams extends ComplexParams for HTTP parameter insertion
type HTTPComplexParams struct {
	*genlib.ComplexParams
	RealValues      []float64
	ImaginaryValues []float64
}

func init() {
	HTTPParams["complex"] = &HTTPComplexParams{&genlib.ComplexParams{}, []float64{}, []float64{}}
}

// SetParams allows for setting of parameters in HTTPComplexParams
func (curParams *HTTPComplexParams) SetParams(k string, v string) error {
	switch k {
	case "real_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if floatV, err := strconv.ParseFloat(v, 64); err == nil {
				curParams.RealValues = append(curParams.RealValues, floatV)
			}
		}
	case "imaginary_values":
		strList := strings.Split(v, ",")
		for _, v := range strList {
			if floatV, err := strconv.ParseFloat(v, 64); err == nil {
				curParams.ImaginaryValues = append(curParams.ImaginaryValues, floatV)
			}
		}
	default:
		if k != "funcid" {
			return fmt.Errorf("getComplexParam : unknown parameter : %v", k)
		}
	}
	if len(curParams.RealValues) != 0 && len(curParams.ImaginaryValues) != 0 {
		if len(curParams.RealValues) != len(curParams.ImaginaryValues) {
			return fmt.Errorf("getComplexParam : mismatched real_values : %v and imaginary_values : %v ", curParams.RealValues, curParams.ImaginaryValues)
		}
		for i := 0; i < len(curParams.RealValues); i++ {
			curParams.ValidValues = append(curParams.ValidValues, complex(float64(curParams.RealValues[i]), float64(curParams.ImaginaryValues[i])))
		}
	}
	return nil
}

// New returns an empty object of the same config type
func (curParams *HTTPComplexParams) New() DataGenHTTP {
	return &HTTPComplexParams{&genlib.ComplexParams{}, []float64{}, []float64{}}
}
