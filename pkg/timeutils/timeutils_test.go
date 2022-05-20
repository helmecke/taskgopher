package timeutils_test

import (
	"testing"
	"time"

	"github.com/helmecke/taskgopher/pkg/timeutils"
)

func TestDiff(t *testing.T) {
	t.Parallel()
	dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)
	type args struct {
		a    time.Time
		b    time.Time
		sign bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"seconds", args{a: dt, b: dt.Add(time.Second), sign: false}, "1s"},
		{"seconds future", args{a: dt, b: dt.Add(time.Second), sign: true}, "+1s"},
		{"seconds past", args{a: dt.Add(time.Second), b: dt, sign: true}, "-1s"},
		{"minutes", args{a: dt, b: dt.Add(time.Minute), sign: false}, "1m"},
		{"minutes future", args{a: dt, b: dt.Add(time.Minute), sign: true}, "+1m"},
		{"minutes past", args{a: dt.Add(time.Minute), b: dt, sign: true}, "-1m"},
		{"hours", args{a: dt, b: dt.Add(time.Hour), sign: false}, "1h"},
		{"hours future", args{a: dt, b: dt.Add(time.Hour), sign: true}, "+1h"},
		{"hours past", args{a: dt.Add(time.Hour), b: dt, sign: true}, "-1h"},
		{"days", args{a: dt, b: dt.Add(time.Hour * 24), sign: false}, "1d"},
		{"days future", args{a: dt, b: dt.Add(time.Hour * 24), sign: true}, "+1d"},
		{"days past", args{a: dt.Add(time.Hour * 24), b: dt, sign: true}, "-1d"},
		{"months", args{a: dt, b: dt.Add(time.Hour * 24 * 7 * 5), sign: false}, "1M"},
		{"months future", args{a: dt, b: dt.Add(time.Hour * 24 * 7 * 5), sign: true}, "+1M"},
		{"months past", args{a: dt.Add(time.Hour * 24 * 7 * 5), b: dt, sign: true}, "-1M"},
		{"years", args{a: dt, b: dt.Add(time.Hour * 24 * 7 * 53), sign: false}, "1y"},
		{"years future", args{a: dt, b: dt.Add(time.Hour * 24 * 7 * 53), sign: true}, "+1y"},
		{"years past", args{a: dt.Add(time.Hour * 24 * 7 * 53), b: dt, sign: true}, "-1y"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := timeutils.Diff(tt.args.a, tt.args.b, tt.args.sign); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
