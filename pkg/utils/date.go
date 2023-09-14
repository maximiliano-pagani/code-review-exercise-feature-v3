package utils

func IsValidYearInterval(startYear, endYear int) bool {
	if IsValidYear(startYear) && IsValidYear(endYear) && startYear <= endYear {
		return true
	}
	return false
}

func IsValidYear(year int) bool {
	if 1900 <= year && year <= 2100 {
		return true
	}
	return false
}
