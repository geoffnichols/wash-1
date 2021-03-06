// Package cmdutil provides utilities for formatting CLI output.
package cmdutil

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// ErrPrintf formats and prints the provided format string and args on stderr and
// colors the output red.
func ErrPrintf(msg string, a ...interface{}) {
	_, err := fmt.Fprintf(color.Error, color.RedString(msg), a...)
	if err != nil {
		panic(err)
	}
}

// FormatDuration formats a duration as `[[dd-]hh:]mm:ss` according to
// http://pubs.opengroup.org/onlinepubs/9699919799/utilities/ps.html.
func FormatDuration(dur time.Duration) string {
	const Decisecond = 100 * time.Millisecond
	const Day = 24 * time.Hour
	d := dur / Day
	dur = dur % Day
	h := dur / time.Hour
	dur = dur % time.Hour
	m := dur / time.Minute
	dur = dur % time.Minute
	s := dur / time.Second
	dur = dur % time.Second
	f := dur / Decisecond
	if d >= 1 {
		return fmt.Sprintf("%02d-%02d:%02d:%02d.%02d", d, h, m, s, f)
	} else if h >= 1 {
		return fmt.Sprintf("%02d:%02d:%02d.%02d", h, m, s, f)
	} else {
		return fmt.Sprintf("%02d:%02d.%02d", m, s, f)
	}
}
