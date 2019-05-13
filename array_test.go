package datagen

import (
	"reflect"
	"testing"
)

func Test_incrementOneLeft(t *testing.T) {
	type args struct {
		intList     []int
		validValues []int
		startingval int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"5 long", args{[]int{0, 1, 0, 1, 0}, []int{0, 1}, 0}, []int{1, 1, 0, 1, 0}},
		{"1 long", args{[]int{0}, []int{0, 1}, 0}, []int{1}},
		{"1 long corner case", args{[]int{1}, []int{0, 1}, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementOneLeft(tt.args.intList, tt.args.validValues, tt.args.startingval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incrementOneLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceIndex(t *testing.T) {
	testList := []int{0, 1, 2, 3, 4, 5}
	type args struct {
		limit     int
		predicate func(i int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"5 long", args{5, func(i int) bool { return testList[i] == 2 }}, 2},
		{"5 long", args{5, func(i int) bool { return testList[i] == 0 }}, 0},
		{"5 long", args{6, func(i int) bool { return testList[i] == 5 }}, 5},
		{"5 long corner case", args{6, func(i int) bool { return testList[i] == 7 }}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceIndex(tt.args.limit, tt.args.predicate); got != tt.want {
				t.Errorf("sliceIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenOneDimensionalRet(t *testing.T) {
	type args struct {
		axisOneLength int
		validValues   []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"3 long 3 items", args{3, []int{0, 1, 2}}, [][]int{{0, 1, 1}, {0, 2, 1}, {1, 2, 1}, {0, 0, 2}, {1, 0, 2}, {2, 0, 2},
			{0, 1, 2}, {1, 1, 2}, {2, 1, 2}, {0, 0, 0}, {1, 0, 0}, {0, 2, 0}, {2, 1, 0}, {2, 2, 0}, {0, 0, 1}, {2, 0, 1}, {1, 0, 1},
			{2, 0, 0}, {0, 1, 0}, {1, 1, 0}, {1, 2, 0}, {1, 1, 1}, {2, 1, 1}, {2, 2, 1}, {2, 2, 2}, {0, 2, 2}, {1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenOneDimensionalRet(tt.args.axisOneLength, tt.args.validValues)
			for _, val := range got {
				equal, _ := checkOneDimensionalInttoList(val, tt.want)
				if !equal {
					t.Errorf("%d not found in %d", val, tt.want)
				}
			}
		})
	}
}

func TestGenOneDimensional(t *testing.T) {
	resultschan := make(chan []int)
	type args struct {
		axisOneLength int
		validValues   []int
		results       chan []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"3 long 3 items", args{3, []int{0, 1, 2}, resultschan}, [][]int{{0, 1, 1}, {0, 2, 1}, {1, 2, 1}, {0, 0, 2}, {1, 0, 2}, {2, 0, 2},
			{0, 1, 2}, {1, 1, 2}, {2, 1, 2}, {0, 0, 0}, {1, 0, 0}, {0, 2, 0}, {2, 1, 0}, {2, 2, 0}, {0, 0, 1}, {2, 0, 1}, {1, 0, 1},
			{2, 0, 0}, {0, 1, 0}, {1, 1, 0}, {1, 2, 0}, {1, 1, 1}, {2, 1, 1}, {2, 2, 1}, {2, 2, 2}, {0, 2, 2}, {1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go GenOneDimensional(tt.args.axisOneLength, tt.args.validValues, tt.args.results)
			CheckOneDimensionalIntChantoList(resultschan, tt.want, t)
		})
	}
}

func benchmarkGenOneDimensionalHelper(axisOneLength int, validValues []int, results chan []int) {
	go GenOneDimensional(axisOneLength, validValues, results)
	for i := range results {
		result = i
	}
}
func benchmarkGenOneDimensional(axisOneLength int, validValues []int, b *testing.B) {
	r := []int{}
	results := make(chan []int, 10)
	for i := 0; i < b.N; i++ {
		results = make(chan []int, 10)
		benchmarkGenOneDimensionalHelper(axisOneLength, validValues, results)
	}
	for i := range results {
		r = i
	}
	result = r
}

var result []int

func BenchmarkGenOneDimensional1(b *testing.B) {
	benchmarkGenOneDimensional(10, []int{0, 1, 2, 3}, b)
}
func BenchmarkGenOneDimensional2(b *testing.B) {
	benchmarkGenOneDimensional(10, []int{0, 1, 2, 3, 4, 5}, b)
}

func benchmarkGenOneDimensionalRet(axisOneLength int, validValues []int, b *testing.B) {
	r := []int{}
	for i := 0; i < b.N; i++ {
		r = GenOneDimensionalRet(axisOneLength, validValues)[0]
	}
	result = r
}

func BenchmarkGenOneDimensionalRet1(b *testing.B) { benchmarkGenOneDimensionalRet(3, []int{0, 1}, b) }
func BenchmarkGenOneDimensionalRet2(b *testing.B) {
	benchmarkGenOneDimensionalRet(10, []int{0, 1, 2, 3, 4, 5}, b)
}
