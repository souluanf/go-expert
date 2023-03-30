package tax

func CalculateTax(amount float64) float64 {
	if amount < 1000 {
		return 0.05
	}
	return 0.1
}
