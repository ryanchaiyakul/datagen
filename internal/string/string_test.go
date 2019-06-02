package stringlib

import (
	"testing"
)

func TestGenString(t *testing.T) {
	type args struct {
		length           int
		asciiRange       [2]int
		permutationRange [2]int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"2 long", args{2, [2]int{0, 0}, [2]int{0, 3}}, []string{"AA", "BA", "CA"}, false},
		{"2 long crash", args{2, [2]int{0, 0}, [2]int{0, 10000}}, nil, true},
		{"10 long", args{10, [2]int{0, 0}, [2]int{0, 10}}, []string{"AAAAAAAAA", "BAAAAAAAA", "CAAAAAAAA", "DAAAAAAAA", "EAAAAAAAA", "FAAAAAAAA", "GAAAAAAAA", "HAAAAAAAA", "IAAAAAAAA", "JAAAAAAAA"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenString(tt.args.length, tt.args.asciiRange, tt.args.permutationRange)
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
