package timeutils

import (
	"regexp"
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

// Abbr return time of date abbreviation
// nolint:gocyclo
func Abbr(abbr string, pointInTime time.Time) time.Time {
	switch abbr {
	// now	Current local date and time.
	case "now":
		return pointInTime.Truncate(time.Second)
	// today	Current local date, with time 00:00:00.
	case "today":
		return pointInTime.Truncate(time.Hour * 24)
	// sod	Local date of the start of the next day, with time 00:00:00. Same as tomorrow.
	case "sod":
		return pointInTime.Truncate(time.Hour*24).AddDate(0, 0, 1)
	// eod	Current local date, with time 23:59:59.
	case "eod":
		return pointInTime.Truncate(time.Hour * 24).Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	// yesterday	Local date for yesterday, with time 00:00:00.
	case "yesterday":
		return pointInTime.Truncate(time.Hour*24).AddDate(0, 0, -1)
	// tomorrow	Local date for tomorrow, with time 00:00:00. Same as sod.
	case "tomorrow":
		return pointInTime.Truncate(time.Hour*24).AddDate(0, 0, 1)
	// monday, tuesday ...	Local date for the specified day, after today, with time 00:00:00.
	case "monday":
		return nextAfterWeekday(1, pointInTime.Truncate(time.Hour*24))
	case "tuesday":
		return nextAfterWeekday(2, pointInTime.Truncate(time.Hour*24))
	case "wednesday":
		return nextAfterWeekday(3, pointInTime.Truncate(time.Hour*24))
	case "thursday":
		return nextAfterWeekday(4, pointInTime.Truncate(time.Hour*24))
	case "friday":
		return nextAfterWeekday(5, pointInTime.Truncate(time.Hour*24))
	case "saturday":
		return nextAfterWeekday(6, pointInTime.Truncate(time.Hour*24))
	case "sunday":
		return nextAfterWeekday(7, pointInTime.Truncate(time.Hour*24))
	// january, february ...	Local date for the specified month, 1st day, with time 00:00:00.
	case "january":
		return nextAfterMonth(1, firstInMonth(pointInTime))
	case "february":
		return nextAfterMonth(2, firstInMonth(pointInTime))
	case "march":
		return nextAfterMonth(3, firstInMonth(pointInTime))
	case "april":
		return nextAfterMonth(4, firstInMonth(pointInTime))
	case "may":
		return nextAfterMonth(5, firstInMonth(pointInTime))
	case "june":
		return nextAfterMonth(6, firstInMonth(pointInTime))
	case "july":
		return nextAfterMonth(7, firstInMonth(pointInTime))
	case "august":
		return nextAfterMonth(8, firstInMonth(pointInTime))
	case "september":
		return nextAfterMonth(9, firstInMonth(pointInTime))
	case "october":
		return nextAfterMonth(10, firstInMonth(pointInTime))
	case "november":
		return nextAfterMonth(11, firstInMonth(pointInTime))
	case "december":
		return nextAfterMonth(12, firstInMonth(pointInTime))
	// later, someday	Local 2038-01-18, with time 00:00:00. A date far away, with semantically meaningful to GTD users.
	case "someday":
		return time.Date(2038, 01, 18, 0, 0, 0, 0, pointInTime.Location())
	case "later":
		return time.Date(2038, 01, 18, 0, 0, 0, 0, pointInTime.Location())
	// soy	Local date for the next year, January 1st, with time 00:00:00.
	case "soy":
		return time.Date(pointInTime.Year()+1, 01, 01, 0, 0, 0, 0, pointInTime.Location())
	// eoy	Local date for this year, December 31st, with time 00:00:00.
	case "eoy":
		return time.Date(pointInTime.Year(), 12, 31, 0, 0, 0, 0, pointInTime.Location())
	// soq	Local date for the start of the next quarter (January, April, July, October), 1st, with time 00:00:00.
	case "soq":
		return nextAfterQuarter(firstInMonth(pointInTime))
		// eoq	Local date for the end of the current quarter (March, June, September, December), last day of the month, with time 23:59:59.
	// som	Local date for the 1st day of the next month, with time 00:00:00.
	case "som":
		y, m, _ := pointInTime.Date()

		return time.Date(y, m+1, 01, 0, 0, 0, 0, pointInTime.Location())
	// socm	Local date for the 1st day of the current month, with time 00:00:00.
	case "socm":
		y, m, _ := pointInTime.Date()

		return time.Date(y, m, 01, 0, 0, 0, 0, pointInTime.Location())
	// eom, eocm	Local date for the last day of the current month, with time 23:59:59.
	case "eom":
		y, m, _ := pointInTime.Date()

		return time.Date(y, m+1, 01, 0, 0, 0, 0, pointInTime.Location()).Add(-time.Second)
	case "eocm":
		y, m, _ := pointInTime.Date()

		return time.Date(y, m+1, 01, 0, 0, 0, 0, pointInTime.Location()).Add(-time.Second)
	// sow	Local date for the next Monday, with time 00:00:00.
	case "sow":
		return nextAfterWeekday(1, pointInTime.Truncate(time.Hour*24))
	// socw	Local date for the last Monday, with time 00:00:00.
	case "socw":
		y, m, d := pointInTime.Date()

		return time.Date(y, m, d-int(pointInTime.Weekday()-1), 0, 0, 0, 0, pointInTime.Location())
	// eow, eocw	Local date for the end of the week, Saturday night, with time 00:00:00.
	case "eow":
		y, m, d := pointInTime.Date()

		return time.Date(y, m, d+int(pointInTime.Weekday()+1), 0, 0, 0, 0, pointInTime.Location())
	case "eocw":
		y, m, d := pointInTime.Date()

		return time.Date(y, m, d+int(pointInTime.Weekday()+1), 0, 0, 0, 0, pointInTime.Location())
	// soww	Local date for the start of the work week, next Monday, with time 00:00:00.
	case "soww":
		return nextAfterWeekday(1, pointInTime.Truncate(time.Hour*24))
	// eoww	Local date for the end of the work week, Friday night, with time 23:59:59.
	case "eoww":
		y, m, d := pointInTime.Date()

		return time.Date(y, m, d+int(pointInTime.Weekday()-1), 23, 59, 59, 0, pointInTime.Location())
		// goodfriday	Local date for the next Good Friday, with time 00:00:00.
		// easter	Local date for the next Easter Sunday, with time 00:00:00.
		// eastermonday	Local date for the next Easter Monday, with time 00:00:00.
		// ascension	Local date for the next Ascension (39 days after Easter Sunday), with time 00:00:00.
		// pentecost	Local date for the next Pentecost (40 days after Easter Sunday), with time 00:00:00.
		// midsommar	Local date for the Saturday after June 20th, with time 00:00:00. Swedish.
		// midsommarafton	Local date for the Friday after June 19th, with time 00:00:00. Swedish.
	}

	// 1st, 2nd, ...	Local date for the next Nth day, with time 00:00:00.
	re := regexp.MustCompile(`(\d+)\w+`)
	matches := re.FindStringSubmatch(abbr)
	if matches != nil {
		day, err := strconv.Atoi(matches[1])
		if err != nil {
			return time.Time{}
		}
		y, m, d := pointInTime.Date()
		if d >= day {
			m++
		}

		return time.Date(y, m, day, 0, 0, 0, 0, pointInTime.Location())
	}

	return time.Time{}
}

func nextAfterWeekday(w int, t time.Time) time.Time {
	diff := w - int(t.Weekday())
	if diff <= 0 {
		diff += 7
	}

	return t.AddDate(0, 0, diff)
}

func nextAfterMonth(m int, t time.Time) time.Time {
	diff := m - int(t.Month())
	if diff <= 0 {
		diff += 12
	}

	return t.AddDate(0, diff, 0)
}

func nextAfterQuarter(t time.Time) time.Time {
	q := [12]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	diff := q[int(t.Month())]
	if diff <= 0 {
		diff += 4
	}

	return t.AddDate(0, diff, 0)
}

func firstInMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	loc := t.Location()

	return time.Date(y, m, 1, 0, 0, 0, 0, loc)
}
