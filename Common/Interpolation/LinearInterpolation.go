package Interpolation

import (
	"errors"
	"math"
)

type LinearInterp struct {
	abscissas []float64
	ordinates []float64
	extrapolation string
}

func LinearInterpolater(xs []float64, ys []float64, extrapScheme string) *LinearInterp {
	return &LinearInterp{abscissas:xs, ordinates:ys, extrapolation:"Flat"}
}

func (interp *LinearInterp) Eval(x float64) (float64, error) {
	// sanity check to ensure abscissas is in ascending order
	v := interp.abscissas[0]
	for _, val := range interp.abscissas {
		if val < v {
			return math.NaN(), errors.New("Abscissas values are not in ascending order")
		} else {
			v = val
		}
	}

	// only flat extrapolation currently
	if x < interp.abscissas[0] {

		return interp.ordinates[0], nil
		
	} else if  x > interp.abscissas[len(interp.abscissas) - 1] {

		return interp.ordinates[len(interp.ordinates) - 1], nil
		
	} else {
	
		var leftIndex, rightIndex int
		
		for idx, val := range interp.abscissas {
			if val > x {
				leftIndex = idx - 1
				rightIndex = idx
				break
			}
		}

		x0 := interp.abscissas[leftIndex]
		y0 := interp.ordinates[leftIndex]
		x1 := interp.abscissas[rightIndex]
		y1 := interp.ordinates[rightIndex]

		y := y0 + (y1 - y0) / (x1 - x0) * (x - x0)

		return y, nil
	}
}
