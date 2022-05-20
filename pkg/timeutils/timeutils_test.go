package timeutils_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/helmecke/taskgopher/pkg/timeutils"
)

func TestDiff(t *testing.T) {
	t.Parallel()
	dt := time.Date(2020, 10, 28, 0, 0, 0, 0, time.UTC)
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

func TestAbbr(t *testing.T) {
	t.Parallel()
	dt := time.Date(2020, 10, 28, 15, 41, 54, 0, time.UTC)
	type args struct {
		abbr        string
		pointInTime time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"now", args{abbr: "now", pointInTime: dt}, time.Date(2020, 10, 28, 15, 41, 54, 0, time.UTC)},
		{"today", args{abbr: "today", pointInTime: dt}, time.Date(2020, 10, 28, 0, 0, 0, 0, time.UTC)},
		{"sod", args{abbr: "sod", pointInTime: dt}, time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC)},
		{"eod", args{abbr: "eod", pointInTime: dt}, time.Date(2020, 10, 28, 23, 59, 59, 0, time.UTC)},
		{"yesterday", args{abbr: "yesterday", pointInTime: dt}, time.Date(2020, 10, 27, 0, 0, 0, 0, time.UTC)},
		{"tomorrow", args{abbr: "tomorrow", pointInTime: dt}, time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC)},
		{"monday", args{abbr: "monday", pointInTime: dt}, time.Date(2020, 11, 02, 0, 0, 0, 0, time.UTC)},
		{"tuesday", args{abbr: "tuesday", pointInTime: dt}, time.Date(2020, 11, 03, 0, 0, 0, 0, time.UTC)},
		{"wednesday", args{abbr: "wednesday", pointInTime: dt}, time.Date(2020, 11, 04, 0, 0, 0, 0, time.UTC)},
		{"thursday", args{abbr: "thursday", pointInTime: dt}, time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC)},
		{"friday", args{abbr: "friday", pointInTime: dt}, time.Date(2020, 10, 30, 0, 0, 0, 0, time.UTC)},
		{"saturday", args{abbr: "saturday", pointInTime: dt}, time.Date(2020, 10, 31, 0, 0, 0, 0, time.UTC)},
		{"sunday", args{abbr: "sunday", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"january", args{abbr: "january", pointInTime: dt}, time.Date(2021, 01, 01, 0, 0, 0, 0, time.UTC)},
		{"february", args{abbr: "february", pointInTime: dt}, time.Date(2021, 02, 01, 0, 0, 0, 0, time.UTC)},
		{"march", args{abbr: "march", pointInTime: dt}, time.Date(2021, 03, 01, 0, 0, 0, 0, time.UTC)},
		{"april", args{abbr: "april", pointInTime: dt}, time.Date(2021, 04, 01, 0, 0, 0, 0, time.UTC)},
		{"may", args{abbr: "may", pointInTime: dt}, time.Date(2021, 05, 01, 0, 0, 0, 0, time.UTC)},
		{"june", args{abbr: "june", pointInTime: dt}, time.Date(2021, 06, 01, 0, 0, 0, 0, time.UTC)},
		{"july", args{abbr: "july", pointInTime: dt}, time.Date(2021, 07, 01, 0, 0, 0, 0, time.UTC)},
		{"august", args{abbr: "august", pointInTime: dt}, time.Date(2021, 8, 01, 0, 0, 0, 0, time.UTC)},
		{"september", args{abbr: "september", pointInTime: dt}, time.Date(2021, 9, 01, 0, 0, 0, 0, time.UTC)},
		{"october", args{abbr: "october", pointInTime: dt}, time.Date(2021, 10, 01, 0, 0, 0, 0, time.UTC)},
		{"november", args{abbr: "november", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"december", args{abbr: "december", pointInTime: dt}, time.Date(2020, 12, 01, 0, 0, 0, 0, time.UTC)},
		{"someday", args{abbr: "someday", pointInTime: dt}, time.Date(2038, 01, 18, 0, 0, 0, 0, time.UTC)},
		{"later", args{abbr: "later", pointInTime: dt}, time.Date(2038, 01, 18, 0, 0, 0, 0, time.UTC)},
		{"soy", args{abbr: "soy", pointInTime: dt}, time.Date(2021, 01, 01, 0, 0, 0, 0, time.UTC)},
		{"eoy", args{abbr: "eoy", pointInTime: dt}, time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)},
		// {"soq", args{abbr: "soq", pointInTime: dt}, time.Date(2021, 01, 01, 0, 0, 0, 0, time.UTC)},
		// {"soq2", args{abbr: "soq", pointInTime: time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC)}, time.Date(2020, 04, 01, 0, 0, 0, 0, time.UTC)},
		// {"eoq", args{abbr: "eoq", pointInTime: dt}, time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC)},
		// soq	Local date for the start of the next quarter (January, April, July, October), 1st, with time 00:00:00.
		// eoq	Local date for the end of the current quarter (March, June, September, December), last day of the month, with time 23:59:59.
		{"som", args{abbr: "som", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"socm", args{abbr: "socm", pointInTime: dt}, time.Date(2020, 10, 01, 0, 0, 0, 0, time.UTC)},
		{"eom", args{abbr: "eom", pointInTime: dt}, time.Date(2020, 10, 31, 23, 59, 59, 0, time.UTC)},
		{"eocm", args{abbr: "eocm", pointInTime: dt}, time.Date(2020, 10, 31, 23, 59, 59, 0, time.UTC)},
		{"sow", args{abbr: "sow", pointInTime: dt}, time.Date(2020, 11, 02, 0, 0, 0, 0, time.UTC)},
		{"socw", args{abbr: "socw", pointInTime: dt}, time.Date(2020, 10, 26, 0, 0, 0, 0, time.UTC)},
		{"eow", args{abbr: "eow", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"eocw", args{abbr: "eocw", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"soww", args{abbr: "soww", pointInTime: dt}, time.Date(2020, 11, 02, 0, 0, 0, 0, time.UTC)},
		{"eoww", args{abbr: "eoww", pointInTime: dt}, time.Date(2020, 10, 30, 23, 59, 59, 0, time.UTC)},
		{"1st", args{abbr: "1st", pointInTime: dt}, time.Date(2020, 11, 01, 0, 0, 0, 0, time.UTC)},
		{"2nd", args{abbr: "2nd", pointInTime: dt}, time.Date(2020, 11, 02, 0, 0, 0, 0, time.UTC)},
		{"3rd", args{abbr: "3rd", pointInTime: dt}, time.Date(2020, 11, 03, 0, 0, 0, 0, time.UTC)},
		{"4th", args{abbr: "4th", pointInTime: dt}, time.Date(2020, 11, 04, 0, 0, 0, 0, time.UTC)},
		{"5th", args{abbr: "5th", pointInTime: dt}, time.Date(2020, 11, 05, 0, 0, 0, 0, time.UTC)},
		{"6th", args{abbr: "6th", pointInTime: dt}, time.Date(2020, 11, 06, 0, 0, 0, 0, time.UTC)},
		{"7th", args{abbr: "7th", pointInTime: dt}, time.Date(2020, 11, 07, 0, 0, 0, 0, time.UTC)},
		{"8th", args{abbr: "8th", pointInTime: dt}, time.Date(2020, 11, 8, 0, 0, 0, 0, time.UTC)},
		{"9th", args{abbr: "9th", pointInTime: dt}, time.Date(2020, 11, 9, 0, 0, 0, 0, time.UTC)},
		{"10th", args{abbr: "10th", pointInTime: dt}, time.Date(2020, 11, 10, 0, 0, 0, 0, time.UTC)},
		{"11th", args{abbr: "11th", pointInTime: dt}, time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC)},
		{"12th", args{abbr: "12th", pointInTime: dt}, time.Date(2020, 11, 12, 0, 0, 0, 0, time.UTC)},
		{"13th", args{abbr: "13th", pointInTime: dt}, time.Date(2020, 11, 13, 0, 0, 0, 0, time.UTC)},
		{"14th", args{abbr: "14th", pointInTime: dt}, time.Date(2020, 11, 14, 0, 0, 0, 0, time.UTC)},
		{"15th", args{abbr: "15th", pointInTime: dt}, time.Date(2020, 11, 15, 0, 0, 0, 0, time.UTC)},
		{"16th", args{abbr: "16th", pointInTime: dt}, time.Date(2020, 11, 16, 0, 0, 0, 0, time.UTC)},
		{"17th", args{abbr: "17th", pointInTime: dt}, time.Date(2020, 11, 17, 0, 0, 0, 0, time.UTC)},
		{"18th", args{abbr: "18th", pointInTime: dt}, time.Date(2020, 11, 18, 0, 0, 0, 0, time.UTC)},
		{"19th", args{abbr: "19th", pointInTime: dt}, time.Date(2020, 11, 19, 0, 0, 0, 0, time.UTC)},
		{"20th", args{abbr: "20th", pointInTime: dt}, time.Date(2020, 11, 20, 0, 0, 0, 0, time.UTC)},
		{"21th", args{abbr: "21th", pointInTime: dt}, time.Date(2020, 11, 21, 0, 0, 0, 0, time.UTC)},
		{"22th", args{abbr: "22th", pointInTime: dt}, time.Date(2020, 11, 22, 0, 0, 0, 0, time.UTC)},
		{"23th", args{abbr: "23th", pointInTime: dt}, time.Date(2020, 11, 23, 0, 0, 0, 0, time.UTC)},
		{"24th", args{abbr: "24th", pointInTime: dt}, time.Date(2020, 11, 24, 0, 0, 0, 0, time.UTC)},
		{"25th", args{abbr: "25th", pointInTime: dt}, time.Date(2020, 11, 25, 0, 0, 0, 0, time.UTC)},
		{"26th", args{abbr: "26th", pointInTime: dt}, time.Date(2020, 11, 26, 0, 0, 0, 0, time.UTC)},
		{"27th", args{abbr: "27th", pointInTime: dt}, time.Date(2020, 11, 27, 0, 0, 0, 0, time.UTC)},
		{"28th", args{abbr: "28th", pointInTime: dt}, time.Date(2020, 11, 28, 0, 0, 0, 0, time.UTC)},
		{"29th", args{abbr: "29th", pointInTime: dt}, time.Date(2020, 10, 29, 0, 0, 0, 0, time.UTC)},
		{"30th", args{abbr: "30th", pointInTime: dt}, time.Date(2020, 10, 30, 0, 0, 0, 0, time.UTC)},
		{"31th", args{abbr: "31th", pointInTime: dt}, time.Date(2020, 10, 31, 0, 0, 0, 0, time.UTC)},
		// goodfriday	Local date for the next Good Friday, with time 00:00:00.
		// easter	Local date for the next Easter Sunday, with time 00:00:00.
		// eastermonday	Local date for the next Easter Monday, with time 00:00:00.
		// ascension	Local date for the next Ascension (39 days after Easter Sunday), with time 00:00:00.
		// pentecost	Local date for the next Pentecost (40 days after Easter Sunday), with time 00:00:00.
		// midsommar	Local date for the Saturday after June 20th, with time 00:00:00. Swedish.
		// midsommarafton	Local date for the Friday after June 19th, with time 00:00:00. Swedish.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := timeutils.Abbr(tt.args.abbr, tt.args.pointInTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Synonym() = %v, want %v", got, tt.want)
			}
		})
	}
}
