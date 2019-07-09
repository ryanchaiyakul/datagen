package httplib

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
	New() DataGenHTTP
	genlib.DataGen
}

// HTTPParams is a dictionary to get the config
var HTTPParams = map[string]DataGenHTTP{}

func getData(r *http.Request) interface{} {
	if dataList, ok := r.URL.Query()["key"]; ok {
		permutationRange := [2]int{}
		if val, ok := r.URL.Query()["permutation_range"]; ok {
			strList := strings.Split(val[0], ",")
			if len(strList) != 2 {
				return fmt.Errorf("getData : permutation_range : %v does not have 2 bounds", strList)
			}
			for k, v := range strList {
				if intV, err := strconv.Atoi(v); err == nil {
					permutationRange[k] = intV
				}
			}
		}
		if val, ok := r.URL.Query()["imaginary_denotation"]; ok {
			imaginaryDenotation = val[0]
		}
		configParam := []*genlib.DataParams{}
		for _, configKey := range strings.Split(dataList[0], ",") {
			tempParam := genlib.DataParams{Name: configKey}
			if configType, ok := r.URL.Query()[configKey+"_type"]; ok {
				var config DataGenHTTP
				if baseConfig, ok := HTTPParams[configType[0]]; ok {
					config = baseConfig.New()
				} else {
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
				for inK, inV := range permutation {
					permutation[inK] = stringComplex(inV)
				}
			}
			return map[string]interface{}{"permutations": ret}
		}
		return err
	}
	return errors.New("getData : 'key' paramter missing")
}
