package genlib

import (
	"testing"

	helperlib "github.com/ryanchaiyakul/datagen/internal/helper"
)

func TestGenString(t *testing.T) {
	type args struct {
		length           int
		asciiValues      []int
		permutationRange [2]int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"2 long", args{2, []int{}, [2]int{0, 3}}, []string{"AA", "BA", "CA"}, false},
		{"2 long crash", args{2, []int{}, [2]int{0, 10000}}, nil, true},
		{"10 long", args{10, []int{}, [2]int{0, 10}}, []string{"AAAAAAAAA", "BAAAAAAAA", "CAAAAAAAA", "DAAAAAAAA", "EAAAAAAAA", "FAAAAAAAA", "GAAAAAAAA", "HAAAAAAAA", "IAAAAAAAA", "JAAAAAAAA"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asciiValues := [][]int{}
			for i := 0; i < helperlib.GetPermutationSliceUnanimous([]int{tt.args.length}, tt.args.asciiValues); i++ {
				asciiValues = append(asciiValues, tt.args.asciiValues)
			}
			got, err := GenString(tt.args.length, asciiValues, tt.args.permutationRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if len(got) != len(tt.want) {
					t.Errorf("GenString() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
