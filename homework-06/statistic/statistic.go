// The Statistic package - example package for demonstration how testing lib works
package statistic

// Average - calculates average value of args
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Summa - calculates sum value of args
func Summa(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
