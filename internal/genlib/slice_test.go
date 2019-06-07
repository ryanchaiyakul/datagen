package genlibinternal

import (
	"testing"
)

func TestGenArray(t *testing.T) {
	type args struct {
		dimensions       []int
		validValues      []int
		permutationRange [2]int
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{"2x2", args{[]int{2, 2}, []int{0, 1}, [2]int{0, 2}}, 2, false},
		{"2x2 Fail", args{[]int{2, 2}, []int{0, 1}, [2]int{0, 100}}, 2, true},
		{"10x10", args{[]int{10, 10}, []int{0, 1}, [2]int{0, 100}}, 100, false},
		{"100x100", args{[]int{100, 100}, []int{0, 1, 2}, [2]int{100, 1000}}, 900, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenArray(tt.args.dimensions, tt.args.validValues, tt.args.permutationRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if len(got.([][]int)) != tt.want.(int) {
					t.Errorf("len(GenArray() = %v, want %v", len(got.([][]int)), tt.want.(int))
				}
			}
		})
	}
}
