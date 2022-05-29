package sliceutils_test

import (
	"testing"

	"github.com/google/uuid"
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

func TestIntIndexOf(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not found", args{slice: []int{1, 2, 4, 5}, value: 6}, -1},
		{"found at 0", args{slice: []int{1, 2, 4, 5}, value: 1}, 0},
		{"found at 1", args{slice: []int{1, 2, 4, 5}, value: 2}, 1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.IntIndexOf(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("IntIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSliceContains(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not found", args{slice: []int{1, 2, 4, 5}, value: 6}, false},
		{"found", args{slice: []int{1, 2, 4, 5}, value: 1}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.IntSliceContains(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("IntSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUIDIndexOf(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []uuid.UUID
		value uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"not found", args{slice: []uuid.UUID{uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d8"), uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d9")}, value: uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7")}, -1},
		{"found", args{slice: []uuid.UUID{uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7"), uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d9")}, value: uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7")}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.UUIDIndexOf(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("UUIDIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUIDSliceContains(t *testing.T) {
	t.Parallel()
	type args struct {
		slice []uuid.UUID
		value uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not found", args{slice: []uuid.UUID{uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d8"), uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d9")}, value: uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7")}, false},
		{"found", args{slice: []uuid.UUID{uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7"), uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d9")}, value: uuid.MustParse("ac5c2795-1fbe-4eff-be32-eb97b10470d7")}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sliceutils.UUIDSliceContains(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("UUIDSliceContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
