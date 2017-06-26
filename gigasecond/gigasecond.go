package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4
const gigasecondSize = 1e9

func AddGigasecond(t time.Time) time.Time {
	var duration time.Duration = time.Second * gigasecondSize
	return t.Add(duration)
}
