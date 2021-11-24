package thixo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHtmlDate(t *testing.T) {
	t.Skip()
	tpl := `{{ htmlDate 0}}`
	if err := runt(tpl, "1970-01-01"); err != nil {
		t.Error(err)
	}
}

func TestAgo(t *testing.T) {
	tpl := "{{ ago .Time }}"
	if err := runtv(tpl, "2m5s", map[string]interface{}{"Time": time.Now().Add(-125 * time.Second)}); err != nil {
		t.Error(err)
	}

	if err := runtv(tpl, "2h34m17s", map[string]interface{}{"Time": time.Now().Add(-(2*3600 + 34*60 + 17) * time.Second)}); err != nil {
		t.Error(err)
	}

	if err := runtv(tpl, "-5s", map[string]interface{}{"Time": time.Now().Add(5 * time.Second)}); err != nil {
		t.Error(err)
	}
}

func TestToDate(t *testing.T) {
	tpl := `{{toDate "2006-01-02" "2017-12-31" | date "02/01/2006"}}`
	if err := runt(tpl, "31/12/2017"); err != nil {
		t.Error(err)
	}
}

func TestToDateInLocation(t *testing.T) {
	tests := map[string]struct {
		location string
		expected string
	}{
		"CET": {
			location: "CET",
			expected: "2006-01-02 03:00:00 +0100 CET",
		},
		"unknown": {
			location: "unknown",
			expected: "2006-01-02 03:00:00 +0000 UTC",
		},
	}

	for title, tt := range tests {
		actual := toDateInLocation("2006-01-02 15:04", "2006-01-02 03:00", tt.location)

		require.Equal(t, tt.expected, actual.String(), title)
	}
}

func TestUnixEpoch(t *testing.T) {
	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "13 Jun 19 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}
	tpl := `{{unixEpoch .Time}}`

	if err = runtv(tpl, "1560458379", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}
}

func TestDateInZone(t *testing.T) {
	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "13 Jun 19 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}
	tpl := `{{ date_in_zone "02 Jan 06 15:04 -0700" .Time "UTC" }}`

	// Test time.Time input
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}

	// Test pointer to time.Time input
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": &tm}); err != nil {
		t.Error(err)
	}

	// Test no time input. This should be close enough to time.Now() we can test
	loc, _ := time.LoadLocation("UTC")
	if err = runtv(tpl, time.Now().In(loc).Format("02 Jan 06 15:04 -0700"), map[string]interface{}{"Time": ""}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int64
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int64(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int32
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int32(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test unix timestamp as int
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": int(1560458379)}); err != nil {
		t.Error(err)
	}

	// Test case of invalid timezone
	tpl = `{{ date_in_zone "02 Jan 06 15:04 -0700" .Time "foobar" }}`
	if err = runtv(tpl, "13 Jun 19 20:39 +0000", map[string]interface{}{"Time": tm}); err != nil {
		t.Error(err)
	}
}

func TestDuration(t *testing.T) {
	tpl := "{{ duration .Secs }}"
	tests := map[string]interface{}{
		"1m1s":     "61",
		"1m35s":    int64(95),
		"1m30s":    int(90),
		"2m30s":    int32(150),
		"40s":      uint32(40),
		"2m2s":     uint64(122),
		"1h0m0s":   "3600",
		"26h3m4s":  "93784",
		"0s":       []int{0, 1},
		"12.32s":   float64(12.32),
		"1m12.32s": float32(72.32),
		"1m19s":    uint(79),
	}

	for expected, in := range tests {
		if err := runtv(tpl, expected, map[string]interface{}{"Secs": in}); err != nil {
			t.Error(err)
		}
	}

}

func TestDurationRound(t *testing.T) {
	tpl := "{{ durationRound .Time }}"
	if err := runtv(tpl, "2h", map[string]interface{}{"Time": "2h5s"}); err != nil {
		t.Error(err)
	}
	if err := runtv(tpl, "1d", map[string]interface{}{"Time": "24h5s"}); err != nil {
		t.Error(err)
	}
	if err := runtv(tpl, "3mo", map[string]interface{}{"Time": "2400h5s"}); err != nil {
		t.Error(err)
	}
}

func TestIssue317(t *testing.T) {
	expectedCET := "2006-01-02 03:00:00 +0000 CET"
	expectedDZ := "2006-01-02 03:00 UTC"

	for _, name := range []string{"GMT", "CET", "America/New_York", "EET", "Africa/Bangui"} {
		var err error
		time.Local, err = time.LoadLocation(name)
		if err != nil {
			panic(err)
		}

		dateCET := toDate("2006-01-02 15:04 MST", "2006-01-02 03:00 CET")
		require.Equal(t, expectedCET, dateCET.String(), name)

		dz := dateInZone("2006-01-02 15:04 MST", dateCET, "UTC")
		require.Equal(t, expectedDZ, dz, name)
	}
}
