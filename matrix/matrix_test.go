package matrix

import "testing"

type testGa struct {
	m Matrix
	solution []float64
}

type testIden struct {
	i Matrix
	n int8
}

type testSol struct {
	a Matrix
	b []float64
	solution []float64
}

var testGauss = []testGa{
	{ Matrix{[]float64{1,5,7},[]float64{-2,-7,-5}}, []float64{-8,3} },
	{ Matrix{[]float64{0,2,1,-8},[]float64{1,-2,-3,0},[]float64{-1,1,2,3}}, []float64{-4,-5,2} },
	{ Matrix{[]float64{1,-2,-6,12},[]float64{2,4,12,-17},[]float64{1,-4,-12,22}}, nil },
	{ Matrix{[]float64{1,1,1,1},[]float64{2,2,2,2}}, nil},
	{ Matrix{[]float64{0,-0.56789,61.01},[]float64{54.02,12.69,1.25}}, []float64{25.260495115726457,-107.4327774745109}},
}

var testIdentity = []testIden{
	{ Matrix{[]float64{1,0},[]float64{0,1}}, 2 },
	{ Matrix{[]float64{1,0,0},[]float64{0,1,0},[]float64{0,0,1}}, 3 },
	{ Matrix{[]float64{1,0,0,0},[]float64{0,1,0,0},[]float64{0,0,1,0},[]float64{0,0,0,1}}, 4 },
	{ Matrix{[]float64{1,0,0,0,0},[]float64{0,1,0,0,0},[]float64{0,0,1,0,0},[]float64{0,0,0,1,0},[]float64{0,0,0,0,1}}, 5 },
}

var testSolution = []testSol{
	{ Matrix{[]float64{0,2},[]float64{1,2}}, []float64{1,1}, []float64{0,0.5} },
	{ Matrix{[]float64{1,2,3},[]float64{5,0,6},[]float64{8,9,0}}, []float64{4,7,10}, []float64{0.5254237288135591,0.6440677966101694,0.728813559322034} },
	{ Matrix{[]float64{0,0,1},[]float64{0,0.47,10},[]float64{0,1,50}}, []float64{50,10,1}, nil },
	{ Matrix{[]float64{0,10,10},[]float64{0,10,10},[]float64{20,1,2}}, []float64{10,5,5}, nil },
	{ Matrix{[]float64{1,2},[]float64{3,4},[]float64{5,6}},[]float64{1,2,3}, nil },
}

func TestGauss(t *testing.T) {
	for _, pair := range testGauss {
		v :=  pair.m.Copy()

		x,err := v.Gauss()
		if err != nil {
			if x == nil && pair.solution == nil {
				continue
			} else {t.Error(err, pair.m)}
		}
		for i,_ := range x {
			if x[i] != pair.solution[i] {
				t.Error(
					"| Gauss | expected", pair.solution,
					"got", x,
					"from matrix", pair.m,
					)
			}
		}
	}
}

func TestIdentity(t *testing.T){
	for _, pair := range testIdentity {
		i := Identity(pair.n)
		if !(&i).Equals(&pair.i) {
			t.Error(
				"| Identity | sxpected", pair.i,
				"got", i,
				)
		}
	}
}

func TestSolution(t *testing.T) {
	for _, pair := range testSolution {
		v :=  pair.a.Copy()

		x,err := v.Solution(pair.b)
		if err != nil {
			if x == nil && pair.solution == nil {
				continue
			} else {t.Error(err, pair.a)}
		}
		if !solutionApproxEquals(x, pair.solution) {
			t.Error(
				"| Solution | expected", pair.solution,
				"got", x,
				"from matrix", pair.a,
				)
		}
	}
}

func TestCreateMatrix(t *testing.T) {
	for i:=0 ; i < 10 ; i++ {
		for j:=0 ; j < 10 ; j++ {
			a := CreateMatrix(uint8(i),uint8(j))
			if i != 0 && j != 0 {
				if a.Columns() != j {
					t.Error(
						"| CreateMatrix | expected", i,
						"rows,", j,
						"columns. Got",a.Rows(),
						"rows and" , a.Columns(),
						"columns",
						)
				}
				if a.Rows() != i {
					t.Error(
						"| CreateMatrix | expected", i,
						"rows,", j,
						"columns. Got",a.Rows(),
						"rows and" , a.Columns(),
						"columns",
						)
				}
			} else if a != nil {
				t.Error(
					"| CreateMatrix | expected", nil,
					"got", a,
					)
			}
		}
	}
}
