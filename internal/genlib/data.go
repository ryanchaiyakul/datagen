package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"

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
	var ret []map[string]interface{}
	results := make(chan map[string]interface{})

	if permutationRange[0] == 0 && permutationRange[1] == 0 {
		permutationRange[1] = getPermutationData(config)
	}

	var err error
	switch permutaitonCount := permutationRange[1] - permutationRange[0]; {
	case permutaitonCount < 100:
		err = genDataHelper(config, permutationRange, results, 10)
	case permutaitonCount < 100:
		err = genDataHelper(config, permutationRange, results, 50)
	default:
		err = genDataHelper(config, permutationRange, results, 100)
	}

	if err != nil {
		return nil, err
	}

	for permutation := range results {
		if err, ok := permutation["error"]; ok {
			return nil, err.(error)
		}
		ret = append(ret, permutation)
	}
	return ret, nil
}

func genDataHelper(config []*DataParams, permutationRange [2]int, results chan map[string]interface{}, routineCount int) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, routineCount)

	for i := permutationRange[0]; i < permutationRange[1]; i++ {
		tempConfig := configCopy(config)
		if err := setPermutation(tempConfig, i); err == nil {
			wg.Add(1)
			sem <- struct{}{}
			go genDataMain(tempConfig, &wg, sem, results)
		} else {
			return err
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return nil
}

func genDataMain(config []*DataParams, wg *sync.WaitGroup, sem chan struct{}, results chan map[string]interface{}) {
	defer func() {
		<-sem
		wg.Done()
	}()

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
			results <- map[string]interface{}{"error": fmt.Errorf("GenData : unknown type : %v requested", v.Type)}
			return
		}
		if err != nil {
			results <- map[string]interface{}{"error": err}
			return
		}
		ret[v.Name] = retVal
	}
	fmt.Println(ret)
	results <- ret
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
			case "string_slice":
				sliceConfig := v.GenConfig.(*StringParams)
				sliceConfig.Permutations = []int{0}
			case "string":
				stringConfig := v.GenConfig.(*StringParams)
				stringConfig.Permutations = []int{0}
			}
		} else {
			permutationCount := 0
			switch v.Type {
			case "int_slice":
				sliceConfig := v.GenConfig.(*IntSliceParams)
				permutationCount = helperlib.GetPermutationIntSlice(sliceConfig.ValidValues)
				sliceConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "int":
				intConfig := v.GenConfig.(*IntParams)
				permutationCount = len(intConfig.ValidValues)
				intConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "complex_slice":
				sliceConfig := v.GenConfig.(*ComplexSliceParams)
				permutationCount = helperlib.GetPermutationComplexSlice(sliceConfig.ValidValues)
				sliceConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "complex":
				complexConfig := v.GenConfig.(*ComplexParams)
				permutationCount = len(complexConfig.ValidValues)
				complexConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "string_slice":
				sliceConfig := v.GenConfig.(*StringSliceParams)
				permutationCount = helperlib.GetPermutationString(sliceConfig.StringValues)
				sliceConfig.Permutations = []int{int(math.Mod(float64(permutation), float64(permutationCount)))}
			case "string":
				stringConfig := v.GenConfig.(*StringParams)
				permutationCount = helperlib.GetPermutationString(stringConfig.StringValues)
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

func configCopy(config []*DataParams) []*DataParams {
	ret := []*DataParams{}
	for _, v := range config {
		tempParam := *v
		var tempConfig interface{}
		switch tempParam.Type {
		case "int_slice":
			tempVal := *v.GenConfig.(*IntSliceParams)
			tempConfig = &tempVal
		case "int":
			tempVal := *v.GenConfig.(*IntParams)
			tempConfig = &tempVal
		case "complex_slice":
			tempVal := *v.GenConfig.(*ComplexSliceParams)
			tempConfig = &tempVal
		case "complex":
			tempVal := *v.GenConfig.(*ComplexParams)
			tempConfig = &tempVal
		case "string_slice":
			tempVal := *v.GenConfig.(*StringSliceParams)
			tempConfig = &tempVal
		case "string":
			tempVal := *v.GenConfig.(*StringParams)
			tempConfig = &tempVal
		}
		tempParam.GenConfig = tempConfig
		ret = append(ret, &tempParam)
	}
	return ret
}

func getPermutationData(config []*DataParams) int {
	permutationCount := 1
	for _, v := range config {
		switch v.Type {
		case "int_slice":
			sliceConfig := v.GenConfig.(*IntSliceParams)
			permutationCount *= helperlib.GetPermutationIntSlice(sliceConfig.ValidValues)
		case "int":
			intConfig := v.GenConfig.(*IntParams)
			permutationCount *= len(intConfig.ValidValues)
		case "complex_slice":
			sliceConfig := v.GenConfig.(*ComplexSliceParams)
			permutationCount = helperlib.GetPermutationComplexSlice(sliceConfig.ValidValues)
		case "complex":
			complexConfig := v.GenConfig.(*ComplexParams)
			permutationCount = len(complexConfig.ValidValues)
		case "string_slice":
			sliceConfig := v.GenConfig.(*StringSliceParams)
			permutationCount = helperlib.GetPermutationString(sliceConfig.StringValues)
		case "string":
			stringConfig := v.GenConfig.(*StringParams)
			permutationCount = helperlib.GetPermutationString(stringConfig.StringValues)
		}
	}
	return permutationCount
}
