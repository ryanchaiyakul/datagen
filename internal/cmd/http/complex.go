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
	RealValues      []int
	ImaginaryValues []int
}

// SetParams allows for setting of parameters in HTTPComplexParams
func (curParams *HTTPComplexParams) SetParams(k string, v string) error {
	switch k {
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

func stringComplex(v interface{}) interface{} {
	switch v.(type) {
	case complex128:
		v = strings.Trim(strings.Trim(fmt.Sprint(v), "("), ")")
	case []complex128:
		ret := []string{}
		for _, complexVal := range v.([]complex128) {
			ret = append(ret, fmt.Sprint(strings.Trim(strings.Trim(fmt.Sprint(complexVal), "("), ")")))
		}
		v = ret
	case [][]complex128:
		ret := [][]string{}
		for _, complexSlice := range v.([][]complex128) {
			tempSlice := []string{}
			for _, complexVal := range complexSlice {
				tempSlice = append(tempSlice, fmt.Sprint(strings.Trim(strings.Trim(fmt.Sprint(complexVal), "("), ")")))
			}
		}
		v = ret
	case [][][]complex128:
		ret := [][][]string{}
		for _, complex2DSlice := range v.([][][]complex128) {
			temp2DSlice := [][]string{}
			for _, complexSlice := range complex2DSlice {
				tempSlice := []string{}
				for _, complexVal := range complexSlice {
					tempSlice = append(tempSlice, fmt.Sprint(strings.Trim(strings.Trim(fmt.Sprint(complexVal), "("), ")")))
				}
				temp2DSlice = append(temp2DSlice, tempSlice)
			}
			ret = append(ret, temp2DSlice)
		}
		v = ret
	case [][][][]complex128:
		ret := [][][][]string{}
		for _, complex3DSlice := range v.([][][][]complex128) {
			temp3DSlice := [][][]string{}
			for _, complex2DSlice := range complex3DSlice {
				temp2DSlice := [][]string{}
				for _, complexSlice := range complex2DSlice {
					tempSlice := []string{}
					for _, complexVal := range complexSlice {
						tempSlice = append(tempSlice, fmt.Sprint(strings.Trim(strings.Trim(fmt.Sprint(complexVal), "("), ")")))
					}
					temp2DSlice = append(temp2DSlice, tempSlice)
				}
				temp3DSlice = append(temp3DSlice, temp2DSlice)
			}
			ret = append(ret, temp3DSlice)
		}
		v = ret
	}
	return v
}
