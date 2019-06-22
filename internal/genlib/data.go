package genlib

import (
	"errors"
	"fmt"
	"math"

	helperlib "github.com/ryanchaiyakul/datagen/internal/helper"
)

// DataParams is for GenStruct
type DataParams struct {
	Name      string
	Type      string
	GenConfig interface{}
}

// GenData generates a 1D map of the requested data types
func GenData(config []*DataParams, permutationRange [2]int) ([]map[string]interface{}, error) {
	ret := []map[string]interface{}{}
	if permutationRange[0] == 0 && permutationRange[1] == 0 {
		if len(config) == 0 {
			return nil, errors.New("GenData : missing config")
		}
		i := 0
		for {
			if err := setPermutation(config, i); err == nil {
				if tempMap, err := genDataHelper(config); err == nil {
					ret = append(ret, tempMap)
				} else {
					return nil, err
				}
			} else {
				break
			}
			i++
		}
	} else {
		for i := permutationRange[0]; i < permutationRange[1]; i++ {
			if err := setPermutation(config, i); err == nil {
				if tempMap, err := genDataHelper(config); err == nil {
					ret = append(ret, tempMap)
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
	}
	return ret, nil
}

func genDataHelper(config []*DataParams) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	for _, v := range config {
		var retVal interface{}
		var err error
		switch v.Type {
		case "int_slice":
			sliceConfig := *v.GenConfig.(*IntSliceParams)
			retList, newErr := GenIntSlice(sliceConfig)
			if newErr == nil {
				retMultiList, shapeErr := helperlib.ReshapeIntSlice(sliceConfig.Dimensions, retList[0])
				if shapeErr == nil {
					retVal = retMultiList
				}
				newErr = shapeErr
			}
			err = newErr
		case "int":
			retList, newErr := GenInt(*v.GenConfig.(*IntParams))
			if newErr == nil {
				retVal = retList[0]
			}
			err = newErr
		case "complex_slice":
			sliceConfig := *v.GenConfig.(*ComplexSliceParams)
			retList, newErr := GenComplexSlice(sliceConfig)
			if newErr == nil {
				retMultiList, shapeErr := helperlib.ReshapeComplexSlice(sliceConfig.Dimensions, retList[0])
				if shapeErr == nil {
					retVal = retMultiList
				}
				newErr = shapeErr
			}
			err = newErr
		case "complex":
			retList, newErr := GenComplex(*v.GenConfig.(*ComplexParams))
			if newErr == nil {
				retVal = retList[0]
			}
			err = newErr
		case "string_slice":
			sliceConfig := *v.GenConfig.(*StringSliceParams)
			retList, newErr := GenStringSlice(sliceConfig)
			if newErr == nil {
				retMultiList, shapeErr := helperlib.ReshapeStringSlice(sliceConfig.Dimensions, retList[0])
				if shapeErr == nil {
					retVal = retMultiList
				}
				newErr = shapeErr
			}
			err = newErr
		case "string":
			retList, newErr := GenString(*v.GenConfig.(*StringParams))
			if newErr == nil {
				retVal = retList[0]
			}
			err = newErr
		default:
			return nil, fmt.Errorf("GenData : unknown type : %v requested", v.Type)
		}
		if err != nil {
			return nil, err
		}
		ret[v.Name] = retVal
	}
	return ret, nil
}
func setPermutation(config []*DataParams, permutation int) error {
	copyPermutation := permutation
	if len(config) == 0 {
		return errors.New("setPermutation : missing config")
	}
	for _, v := range config {
		if permutation == 0 {
			switch v.Type {
			case "int_slice":
				sliceConfig := v.GenConfig.(*IntSliceParams)
				sliceConfig.Permutations = []int{0}
			case "int":
				intConfig := v.GenConfig.(*IntParams)
				intConfig.Permutations = []int{0}
			case "complex_slice":
				sliceConfig := v.GenConfig.(*ComplexSliceParams)
				sliceConfig.Permutations = []int{0}
			case "complex":
				complexConfig := v.GenConfig.(*ComplexParams)
				complexConfig.Permutations = []int{0}
			case "string":
				stringConfig := v.GenConfig.(*StringParams)
				stringConfig.Permutations = []int{0}
			}
		} else {
			permutationCount := 0
			switch v.Type {
			case "int_slice":
				sliceConfig := v.GenConfig.(*IntSliceParams)
				permutationCount = helperlib.GetPermutationIntSlice(sliceConfig.Dimensions, sliceConfig.ValidValues)
				sliceConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "int":
				intConfig := v.GenConfig.(*IntParams)
				permutationCount = len(intConfig.ValidValues)
				intConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "complex_slice":
				sliceConfig := v.GenConfig.(*ComplexSliceParams)
				permutationCount = helperlib.FlatIntSliceLength(sliceConfig.Dimensions)
				sliceConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "complex":
				complexConfig := v.GenConfig.(*ComplexParams)
				permutationCount = len(complexConfig.ValidValues)
				complexConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "string":
				stringConfig := v.GenConfig.(*StringParams)
				permutationCount = helperlib.GetPermutationString(stringConfig.Length, stringConfig.StringValues)
				stringConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			}
			permutation = permutation / permutationCount
		}
	}
	if permutation == 0 {
		return nil
	}
	return fmt.Errorf("setPermutation : permutation : %v out of range", copyPermutation)
}
