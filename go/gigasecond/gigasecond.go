package gigasecond

import "time"

// AddGigasecond : Returns the moment in time that would be after a gigasecond has passed
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
