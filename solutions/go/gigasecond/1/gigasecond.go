
package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1000000000) * time.Second
	return t.Add(gigasecond)
}
