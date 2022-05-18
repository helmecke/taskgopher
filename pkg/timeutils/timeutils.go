package timeutils

import (
	"strconv"
	"time"
)

// Diff returns duration between two dates as string.
func Diff(a, b time.Time, sign bool) string {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}

	after := ""

	if sign {
		after = "+"
	}

	if sign && a.After(b) {
		after = "-"
	}

	if a.After(b) {
		a, b = b, a
	}

	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := y2 - y1
	month := int(M2 - M1)
	day := d2 - d1
	hour := h2 - h1
	min := m2 - m1
	sec := s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}

	if min < 0 {
		min += 60
		hour--
	}

	if hour < 0 {
		hour += 24
		day--
	}

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}

	if month < 0 {
		month += 12
		year--
	}

	if year > 0 {
		return after + strconv.Itoa(year) + "y"
	}

	if month > 0 {
		return after + strconv.Itoa(month) + "M"
	}

	if day > 0 {
		if day >= 7 {
			return after + strconv.Itoa(day/7) + "w"
		}

		return after + strconv.Itoa(day) + "d"
	}

	if hour > 0 {
		return after + strconv.Itoa(hour) + "h"
	}

	if min > 0 {
		return after + strconv.Itoa(min) + "m"
	}

	if sec > 0 {
		return after + strconv.Itoa(sec) + "s"
	}

	return ""
}

// now	Current local date and time.
// today	Current local date, with time 00:00:00.
// sod	Local date of the start of the next day, with time 00:00:00. Same as tomorrow.
// eod	Current local date, with time 23:59:59.
// yesterday	Local date for yesterday, with time 00:00:00.
// tomorrow	Local date for tomorrow, with time 00:00:00. Same as sod.
// monday, tuesday ...	Local date for the specified day, after today, with time 00:00:00.
// january, february ...	Local date for the specified month, 1st day, with time 00:00:00.
// later, someday	Local 2038-01-18, with time 00:00:00. A date far away, with semantically meaningful to GTD users.
// soy	Local date for the next year, January 1st, with time 00:00:00.
// eoy	Local date for this year, December 31st, with time 00:00:00.
// soq	Local date for the start of the next quarter (January, April, July, October), 1st, with time 00:00:00.
// eoq	Local date for the end of the current quarter (March, June, September, December), last day of the month, with time 23:59:59.
// som	Local date for the 1st day of the next month, with time 00:00:00.
// socm	Local date for the 1st day of the current month, with time 00:00:00.
// eom, eocm	Local date for the last day of the current month, with time 23:59:59.
// sow	Local date for the next Sunday, with time 00:00:00.
// socw	Local date for the last Sunday, with time 00:00:00.
// eow, eocw	Local date for the end of the week, Saturday night, with time 00:00:00.
// soww	Local date for the start of the work week, next Monday, with time 00:00:00.
// eoww	Local date for the end of the work week, Friday night, with time 23:59:59.
// 1st, 2nd, ...	Local date for the next Nth day, with time 00:00:00.
// goodfriday	Local date for the next Good Friday, with time 00:00:00.
// easter	Local date for the next Easter Sunday, with time 00:00:00.
// eastermonday	Local date for the next Easter Monday, with time 00:00:00.
// ascension	Local date for the next Ascension (39 days after Easter Sunday), with time 00:00:00.
// pentecost	Local date for the next Pentecost (40 days after Easter Sunday), with time 00:00:00.
// midsommar	Local date for the Saturday after June 20th, with time 00:00:00. Swedish.
// midsommarafton	Local date for the Friday after June 19th, with time 00:00:00. Swedish.
