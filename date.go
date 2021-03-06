package thixo

import (
	"strconv"
	"time"
)

// Given a format and a date, format the date string.
//
// Date can be a `time.Time` or an `int, int32, int64`.
// In the later case, it is treated as seconds since UNIX
// epoch.
func date(fmt string, date interface{}) string {
	return dateInZone(fmt, date, "Local")
}

func htmlDate(date interface{}) string {
	return dateInZone("2006-01-02", date, "Local")
}

func htmlDateInZone(date interface{}, zone string) string {
	return dateInZone("2006-01-02", date, zone)
}

func dateInZone(fmt string, date interface{}, zone string) string {
	var t time.Time
	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case *time.Time:
		t = *date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	case int32:
		t = time.Unix(int64(date), 0)
	}

	loc, err := time.LoadLocation(zone)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}

	return t.In(loc).Format(fmt)
}

func dateModify(fmt string, date time.Time) time.Time {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return date
	}
	return date.Add(d)
}

func mustDateModify(fmt string, date time.Time) (time.Time, error) {
	d, err := time.ParseDuration(fmt)
	if err != nil {
		return time.Time{}, err
	}
	return date.Add(d), nil
}

func dateAgo(date interface{}) string {
	var t time.Time

	switch date := date.(type) {
	default:
		t = time.Now()
	case time.Time:
		t = date
	case int64:
		t = time.Unix(date, 0)
	case int:
		t = time.Unix(int64(date), 0)
	}
	// Drop resolution to seconds
	duration := time.Since(t).Round(time.Second)
	return duration.String()
}

func duration(sec interface{}) string {
	return parseDuration(sec).String()
}

func durationRound(duration interface{}) string {
	d := parseDuration(duration)

	u := uint64(d)
	neg := d < 0
	if neg {
		u = -u
	}

	var (
		year   = uint64(time.Hour) * 24 * 365
		month  = uint64(time.Hour) * 24 * 30
		day    = uint64(time.Hour) * 24
		hour   = uint64(time.Hour)
		minute = uint64(time.Minute)
		second = uint64(time.Second)
	)
	switch {
	case u > year:
		return strconv.FormatUint(u/year, 10) + "y"
	case u > month:
		return strconv.FormatUint(u/month, 10) + "mo"
	case u > day:
		return strconv.FormatUint(u/day, 10) + "d"
	case u > hour:
		return strconv.FormatUint(u/hour, 10) + "h"
	case u > minute:
		return strconv.FormatUint(u/minute, 10) + "m"
	case u > second:
		return strconv.FormatUint(u/second, 10) + "s"
	}
	return "0s"
}

func toDate(fmt, str string) time.Time {
	t, _ := time.ParseInLocation(fmt, str, time.UTC)
	return t
}

func mustToDate(fmt, str string) (time.Time, error) {
	return time.ParseInLocation(fmt, str, time.Local)
}

func toDateInLocation(fmt, str, location string) time.Time {
	t, _ := mustToDateInLocation(fmt, str, location)

	return t
}

func mustToDateInLocation(fmt, str, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.UTC
	}

	return time.ParseInLocation(fmt, str, loc)
}

func unixEpoch(date time.Time) string {
	return strconv.FormatInt(date.Unix(), 10)
}

func parseDuration(i interface{}) time.Duration {
	var n int64
	switch value := i.(type) {
	default:
		n = 0
	case int64:
		n = value
	case int32:
		n = int64(value)
	case int:
		n = int64(value)
	case uint32:
		n = int64(value)
	case uint64:
		n = int64(value)
	case uint:
		n = int64(value)
	case float32:
		return time.Duration(value * float32(time.Second))
	case float64:
		return time.Duration(value * float64(time.Second))
	case string:
		d, err := time.ParseDuration(value)
		if err == nil {
			return d
		}
		n, _ = strconv.ParseInt(value, 10, 64)
	case time.Time:
		return time.Since(value)
	}
	return time.Duration(n) * time.Second
}
