package sliceutils_test

import (
	"testing"

	"github.com/helmecke/taskgopher/pkg/sliceutils"
)

func TestStrIndexOf(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []string
		value string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not found", args{slice: []string{"val1", "val2", "val3"}, value: "val4"}, -1},
		{"found at 0", args{slice: []string{"val1", "val2", "val3"}, value: "val1"}, 0},
		{"found at 1", args{slice: []string{"val1", "val2", "val3"}, value: "val2"}, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.StrIndexOf(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("StrIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrSliceContains(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not found", args{slice: []string{"val1", "val2", "val3"}, value: "val4"}, false},
		{"found", args{slice: []string{"val1", "val2", "val3"}, value: "val2"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.StrSliceContains(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("StrSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
