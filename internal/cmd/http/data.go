package httpmainlib

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	genlib "github.com/ryanchaiyakul/datagen/internal/genlib"
)

func getData(r *http.Request) interface{} {
	if dataList, ok := r.URL.Query()["key"]; ok {
		configParam := []*genlib.DataParams{}
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
		dataStrList := strings.Split(dataList[0], ",")
		for _, configKey := range dataStrList {
			tempParam := genlib.DataParams{Name: configKey}
			if configType, ok := r.URL.Query()[configKey+"_type"]; ok {
				tempParam.Type = configType[0]
			} else {
				return fmt.Errorf("getData : type : %v does not exist", configKey+"_type")
			}

			switch tempParam.Type {
			case "int_slice":
				tempParam.GenConfig = &genlib.IntSliceParams{}
			case "int":
				tempParam.GenConfig = &genlib.IntParams{}
			case "complex_slice":
				tempParam.GenConfig = &genlib.ComplexSliceParams{}
			case "complex":
				tempParam.GenConfig = &genlib.ComplexParams{}
			case "string_slice":
				tempParam.GenConfig = &genlib.StringSliceParams{}
			case "string":
				tempParam.GenConfig = &genlib.StringParams{}
			default:
				return fmt.Errorf("getData : unknown type : %v", tempParam.Type)
			}

			for k, v := range r.URL.Query() {
				if key := strings.TrimPrefix(k, tempParam.Name+"_"); key != k && key != "type" {
					if key == "permutation_range" {
						return fmt.Errorf("getData : cannot specifiy permutation_range for %v", tempParam.Name)
					}
					var err error
					switch tempParam.Type {
					case "int_slice":
						sliceConfig := tempParam.GenConfig.(*genlib.IntSliceParams)
						err = getIntSliceParam(key, v, sliceConfig)
					case "int":
						intConfig := tempParam.GenConfig.(*genlib.IntParams)
						err = getIntParam(key, v, intConfig)
					case "complex_slice":
						sliceConfig := tempParam.GenConfig.(*genlib.ComplexSliceParams)
						err = getComplexSliceParam(key, v, sliceConfig)
					case "complex":
						complexConfig := tempParam.GenConfig.(*genlib.ComplexParams)
						err = getComplexParam(key, v, complexConfig)
					case "string_slice":
						sliceConfig := tempParam.GenConfig.(*genlib.StringSliceParams)
						err = getStringSliceParam(key, v, sliceConfig)
					case "string":
						stringConfig := tempParam.GenConfig.(*genlib.StringParams)
						err = getStringParam(key, v, stringConfig)
					default:
						return fmt.Errorf("getData : unknown parameter : %v", v)
					}
					if err != nil {
						return err
					}
				}
			}
			configParam = append(configParam, &tempParam)
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
