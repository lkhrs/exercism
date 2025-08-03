package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pRate := float64(productionRate)
	return pRate * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	pRate := float64(productionRate/60)
	return int(pRate * (successRate/100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	batch := carsCount / 10
	batchPrice := 95000
	indPrice := 10000
	return uint(batch * batchPrice + (carsCount - batch * 10) * indPrice)
}
