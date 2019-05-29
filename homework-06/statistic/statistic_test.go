package statistic

import (
	"fmt"
	"testing"
)

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

type testpairSQR struct {
	values [3]float64
	result []float64
	err    error
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

var testSQREqu = []testpairSQR{
	{[3]float64{float64(-12), float64(5), float64(43)}, []float64{-1.6960658136710733, 2.11273248033774}, nil},
	{[3]float64{float64(-77), float64(-77), float64(-77)}, []float64{}, fmt.Errorf("no roots")},
	{[3]float64{float64(800), float64(-77000), float64(0)}, []float64{float64(96.25), float64(0)}, nil},
	{[3]float64{float64(0), float64(1), float64(99)}, []float64{float64(-99)}, nil},
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

// TestSqrEquation -
func TestSqrEquation(t *testing.T) {
	for i := 0; i < len(testSQREqu); i++ {
		roots, terr := SqrEquation(testSQREqu[i].values)
		if (terr == nil && testSQREqu[i].err != nil) || (terr != nil && testSQREqu[i].err == nil) || !compareSliceFloat64(roots, testSQREqu[i].result) {
			t.Errorf("For %v expected %v, but got %v", testSQREqu[i].values, testSQREqu[i].result, roots)
		}
	}
}
