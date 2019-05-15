// Package statistic - example package for demonstration how testing lib works
package statistic

import (
	"fmt"
	"math"
)

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

// SqrEquation - find roots of square equation.
// Using universal method with discriminar
func SqrEquation(abc [3]float64) (roots []float64, err error) {
	a := abc[0]
	b := abc[1]
	c := abc[2]
	if a == 0 && b == 0 && c != 0 { // 0 + 0 + Num = 0, where Num!=0 - nonsense
		return nil, fmt.Errorf("%.2f != 0, no roots", abc[2])
	}
	if a == 0 && b != 0 { // 0 + bx + c = 0 -> x=-c/b
		roots = append(roots, -1*c/b)
		return
	}
	// discriminar calculate
	d := b*b - 4*a*c
	// calculate roots
	if d < 0 {
		err = fmt.Errorf("for %v discriminar=%.2f<0, no roots", abc, d)
	} else if d == 0 {
		roots = append(roots, float64(0))
		roots = append(roots, (-1*b)/(2*a))
	} else {
		sqrtD := math.Sqrt(d)
		roots = append(roots, (-1*b+sqrtD)/(2*a))
		roots = append(roots, (-1*b-sqrtD)/(2*a))
	}
	return
}

// compareSliceFloat64 -
func compareSliceFloat64(a, b []float64) bool {
	eq := true
	if len(a) == 0 && len(b) == 0 {
		return eq
	}
	if len(a) == len(b) {
		for i, v := range a {
			if b[i] != v {
				eq = false
			}
		}
	} else {
		return false
	}
	return eq
}
