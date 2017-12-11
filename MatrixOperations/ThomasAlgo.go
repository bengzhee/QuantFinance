package MatrixOperations

import (
	"errors"
)

type ThomasAlgorithm struct {
	A []float64
	B []float64
	C []float64
	R []float64
}


func (mx *ThomasAlgorithm) Calculate() ([]float64, error) {
	N := len(mx.A)
	v := make([]float64, N)

	if mx.areSameLength(){
		CPrime := make([]float64, N)
		RPrime := make([]float64, N)

		CPrime[0] = mx.C[0] / mx.B[0]
		RPrime[0] = mx.R[0] / mx.B[0]
		
		for i := 1; i < N; i++ {
			CPrime[i] = mx.C[i] / (mx.B[i] - mx.A[i] * CPrime[i-1])
			RPrime[i] = (mx.R[i] - mx.A[i] * RPrime[i-1]) / (mx.B[i] - mx.A[i] * CPrime[i-1])
		}
		
		v[N-1] = RPrime[N-1]
		for i := N - 1; i > 0; i-- {
			v[i-1] = RPrime[i-1] - CPrime[i-1] * v[i]
		}

	} else {
		return nil, errors.New("InputError: Input arrays are not of the same length!")
	}

	return v, nil
}

func (mx *ThomasAlgorithm) areSameLength() bool {
	if len(mx.A) == len(mx.B) && len(mx.A) == len(mx.C) && len(mx.A) == len(mx.R) {
		return true
	}

	return false
}
