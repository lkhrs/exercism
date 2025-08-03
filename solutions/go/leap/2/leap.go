package leap

// IsLeapYear checks if the given year is a leap year
func IsLeapYear(year int) (leap bool) {
	if year%4 == 0 {
		leap = true
	}
	if year%100 == 0 {
		leap = false
		if year%400 == 0 {
			leap = true
		}
	}
	return
}
