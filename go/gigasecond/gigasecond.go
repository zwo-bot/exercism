// Package gigasecond
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1_000_000_000)
	return t
}
