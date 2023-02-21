package debounce

import (
	"time"
)

func debounce(delay time.Duration, f func()) func() {
	timer := time.NewTimer(delay)

	return func() {
		timer.Stop()
		timer = time.NewTimer(delay)
		<-timer.C
		timer.Reset(delay)
		f()
		timer.Stop()
	}
}
