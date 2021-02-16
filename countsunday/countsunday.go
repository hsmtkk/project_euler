package countsunday

import (
	"time"
)

func Answer() int {
	count := 0
	pos := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)
	for pos.Before(end) {
		if pos.Weekday() == time.Sunday {
			count++
		}
		pos = pos.AddDate(0, 1, 0)
	}
	return count
}
