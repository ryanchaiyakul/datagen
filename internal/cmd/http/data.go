package httpmainlib

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

// DataGenHTTP adds an additional function to the interface
type DataGenHTTP interface {
	SetParams(k string, v string) error
	genlib.DataGen
}

func getData(r *http.Request) interface{} {
	if dataList, ok := r.URL.Query()["key"]; ok {
		permutationRange := [2]int{}
		if val, ok := r.URL.Query()["permutation_range"]; ok {
			strList := strings.Split(val[0], ",")
			if len(strList) != 2 {
				return fmt.Errorf("getSliceParam : permutation_range : %v does not have 2 bounds", strList)
			}
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					permutationRange[k] = intV
				}
			}
		}

		configParam := []*genlib.DataParams{}
		for _, configKey := range strings.Split(dataList[0], ",") {
			tempParam := genlib.DataParams{Name: configKey}
			if configType, ok := r.URL.Query()[configKey+"_type"]; ok {
				var config DataGenHTTP
				switch configType[0] {
				case "int_slice":
					config = &HTTPIntSliceParams{&genlib.IntSliceParams{}, []int{}, []int{}}
				case "int":
					config = &HTTPIntParams{&genlib.IntParams{}}
				case "complex_slice":
					config = &HTTPComplexSliceParams{&genlib.ComplexSliceParams{}, []int{}, []int{}, []int{}}
				case "complex":
					config = &HTTPComplexParams{&genlib.ComplexParams{}, []int{}, []int{}}
				case "string_slice":
					config = &HTTPStringSliceParams{&genlib.StringSliceParams{}, []string{}, []int{}}
				case "string":
					config = &HTTPStringParams{&genlib.StringParams{}, []string{}, []int{}}
				default:
					return fmt.Errorf("getData : unknown type : %v", configType[0])
				}

				for k, v := range r.URL.Query() {
					if key := strings.TrimPrefix(k, tempParam.Name+"_"); key != k && key != "type" {
						err := config.SetParams(key, v[0])
						if err != nil {
							return err
						}
					}
				}
				tempParam.GenConfig = config
				configParam = append(configParam, &tempParam)
			} else {
				return fmt.Errorf("getData : type : %v does not exist", configKey+"_type")
			}
		}

		ret, err := genlib.GenData(configParam, permutationRange)
		if err == nil {
			for _, permutation := range ret {
				for k, v := range permutation {
					permutation[k] = stringComplex(v)
				}
			}
			return ret
		}
		return err
	}
	return errors.New("getData : 'key' paramter missing")
}
