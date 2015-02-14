package matrix

import "testing"

type test struct {
	m Matrix
	solution []float64
}

var tests = []test{
	{ Matrix{[]float64{1,5,7},[]float64{-2,-7,-5}}, []float64{-8,3} },
	{ Matrix{[]float64{0,2,1,-8},[]float64{1,-2,-3,0},[]float64{-1,1,2,3}}, []float64{-4,-5,2} },
	{ Matrix{[]float64{1,-2,-6,12},[]float64{2,4,12,-17},[]float64{1,-4,-12,22}}, nil },
	{ Matrix{[]float64{1,1,1,1},[]float64{2,2,2,2}}, nil},
}

func TestGauss(t *testing.T) {
	for _, pair := range tests {
		v :=  pair.m.Copy()

		b,err := v.Gauss()
		if err != nil {
			if b == nil && pair.solution == nil {
				continue
			} else {t.Error(err, pair.m)}
		}
		for i,_ := range b {
			if b[i] != pair.solution[i] {
				t.Error(
					"Expected solution", pair.solution,
					"got", b,
					"from matrix", pair.m,
					)
			}
		}
	}
}