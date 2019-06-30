package genlib

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

// DataParams is for GenData
type DataParams struct {
	Name      string
	GenConfig DataGen
}

// DataGen is the standard interface for the data types
type DataGen interface {
	Gen() error
	Extract(int) (interface{}, error)
	PermutationCount() int
	SetPermutation([]int)
}

//GenData returns a list of permutations of the requested data types in map form
func GenData(config []*DataParams, permutationRange [2]int) ([]map[string]interface{}, error) {
	if len(config) == 0 {
		return nil, errors.New("GenData : missing config")
	}
	if permutationRange[0] < 0 {
		return nil, fmt.Errorf("GenData : permutationRange lower bound : %v out of range", permutationRange[0])
	}

	permutationTotal := getPermutationData(config)
	if permutationRange[1] > permutationTotal {
		return nil, fmt.Errorf("GenData : permutationRange higher bound : %v out of range", permutationRange[1])
	} else if permutationRange[0] == 0 && permutationRange[1] == 0 {
		permutationRange[1] = permutationTotal
	}
	permutationCount := permutationRange[1] - permutationRange[0]
	permutationMax, permutationMap, err := setPermutation(config, permutationRange)

	if err != nil {
		return nil, err
	}

	routineCount := 0
	bufferCount := 0
	switch {
	case permutationCount < 100:
		routineCount = 10
		bufferCount = 5
	case permutationCount < 100:
		routineCount = 50
		bufferCount = 10
	default:
		routineCount = 100
		bufferCount = 20
	}
	results := make(chan map[int]map[string]interface{}, bufferCount)

	go genDataHelper(config, permutationMap, permutationMax, permutationRange, routineCount, results)
	ret := make([]map[string]interface{}, permutationCount)
	for permutation := range results {
		for k, v := range permutation {
			if err, ok := permutation[-1]; ok {
				return nil, err["error"].(error)
			}
			ret[k-permutationRange[0]] = v
		}
	}
	return ret, nil
}

func setPermutation(config []*DataParams, permutationRange [2]int) (map[string]int, map[string]map[int]int, error) {
	permutationsMax := map[string]int{}
	permutationMap := map[string]map[int]int{}
	tempRange := permutationRange

	for _, v := range config {
		genConfig := v.GenConfig
		permutationCount := genConfig.PermutationCount()

		tempMap := map[int]int{}
		tempPermutations := []int{}
		if tempRange[0] == 0 && tempRange[1] == 0 {
			tempPermutations, tempMap = setPermutationHelper(tempPermutations, tempMap, [2]int{0, 1})
		} else {
			lowerbound := 0
			upperbound := permutationCount
			if tempRange[1]-tempRange[0] < permutationCount {
				lowerbound = int(math.Mod(float64(tempRange[0]), float64(permutationCount)))
				upperbound = int(math.Mod(float64(tempRange[1]), float64(permutationCount)))
				// corner case when the upperbound extends past permutationCount but does not fully circle
				if upperbound < lowerbound {
					tempPermutations, tempMap = setPermutationHelper(tempPermutations, tempMap, [2]int{0, upperbound})
					upperbound = permutationCount
				}
			}
			tempPermutations, tempMap = setPermutationHelper(tempPermutations, tempMap, [2]int{lowerbound, upperbound})
			tempRange[0], tempRange[1] = tempRange[0]/permutationCount, tempRange[1]/permutationCount
		}
		permutationsMax[v.Name] = permutationCount
		permutationMap[v.Name] = tempMap
		genConfig.SetPermutation(tempPermutations)
		if err := genConfig.Gen(); err != nil {
			return nil, nil, err
		}
	}
	return permutationsMax, permutationMap, nil
}

func setPermutationHelper(tempPermutations []int, permutationMap map[int]int, bounds [2]int) ([]int, map[int]int) {
	index := len(tempPermutations)
	for i := bounds[0]; i < bounds[1]; i++ {
		permutationMap[i] = index
		tempPermutations = append(tempPermutations, i)
		index++
	}
	return tempPermutations, permutationMap
}

func genDataHelper(config []*DataParams, permutationMap map[string]map[int]int, permutationMax map[string]int, permutationRange [2]int, routineCount int, results chan map[int]map[string]interface{}) {
	sem := make(chan struct{}, routineCount)
	var wg sync.WaitGroup
	wg.Add(permutationRange[1] - permutationRange[0])

	for i := permutationRange[0]; i < permutationRange[1]; i++ {
		sem <- struct{}{}
		go genDataMain(config, permutationMap, permutationMax, i, &wg, sem, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func genDataMain(config []*DataParams, permutationMap map[string]map[int]int, permutationMax map[string]int, permutation int, wg *sync.WaitGroup, sem chan struct{}, results chan map[int]map[string]interface{}) {
	defer func() {
		<-sem
		wg.Done()
	}()
	copyPermutation := permutation
	ret := map[string]interface{}{}
	for _, v := range config {
		perm := permutationMap[v.Name][permutation]
		if permutation > permutationMax[v.Name] {
			perm = permutation
		}
		tempInterface, err := v.GenConfig.Extract(int(math.Mod(float64(perm), float64(permutationMax[v.Name]))))
		if err != nil {
			results <- map[int]map[string]interface{}{-1: {"error": err}}
			return
		}
		ret[v.Name] = tempInterface
		permutation = permutation / permutationMax[v.Name]
	}
	if permutation != 0 {
		results <- map[int]map[string]interface{}{-1: {"error": fmt.Sprintf("genDataMain : permutation : %v out of range", copyPermutation)}}
		return
	}
	results <- map[int]map[string]interface{}{copyPermutation: ret}
}

func getPermutationData(config []*DataParams) int {
	permutationCount := 1
	for _, v := range config {
		permutationCount *= v.GenConfig.PermutationCount()
	}
	return permutationCount
}
