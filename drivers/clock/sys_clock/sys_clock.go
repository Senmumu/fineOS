package sysclock

import "time"

func SysClock() time.Time {
	return time.Now()
}
