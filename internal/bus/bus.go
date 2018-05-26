package bus

import "time"

type Bus struct {
	linie     string
	uhrzeit   time.Time
	wochentag []bool
}
