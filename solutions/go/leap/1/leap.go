package leap

import (
	"time"
)

// IsLeapYear uses package time to check if the given year is a leap year. It also completely ignores the instruction's suggested method.
func IsLeapYear(year int) bool {
	return time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay() == 366
}
