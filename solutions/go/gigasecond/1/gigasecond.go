// For a given time, add one gigasecond.
package gigasecond

import "time"

// For t, add one gigasecond
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1000 * 1000 * 1000)
	return t.Add(time.Second * gigasecond)
}
