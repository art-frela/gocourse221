package statistic

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

type testpair struct {
	values []float64
	result float64
}

var testsAVG = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var testsSUM = []testpair{
	{[]float64{-100, 200, 50, -150, 5, 6, 9}, 20},
	{[]float64{-100000, 50000, 25000, -12500, 45321, 3456, 3235, -12456, -342, -1014}, 700},
	{[]float64{0, 0, 0, -0, +1, 2, 5}, 8},
}

// TestAverageSet - test func for Average()
func TestAverageSet(t *testing.T) {
	for _, pair := range testsAVG {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

// TestSummaSet - test func for Summa()
func TestSummaSet(t *testing.T) {
	for _, pair := range testsSUM {
		v := Summa(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
