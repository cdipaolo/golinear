package matrix

import "errors"
//import "fmt"

type Matrix [][]float64

func abs(n float64) float64 {
	if n < 0 {return -n}
	return n
}

// copies a one dimensional array of float64's
func copy(a []float64) []float64 {
	b := make([]float64, cap(a))
	for i,val := range a {
		b[i] = val
	}
	return b
}


// returns a boolean to whether a matrix is of the same size and strictly equal to another matrix
func (m *Matrix) IsEqualTo(a *Matrix) bool {
	if cap(*m) != cap(*a) || cap((*m)[0]) != cap((*a)[0]) {
		// matrix sizes are not equal
		return false
	}

	// now we know the matrices have the same size
	for i,_ := range *m {
		for j,_ := range (*m)[0] {
			// if entries are different return false
			if (*m)[i][j] != (*a)[i][j] {return false}
		}
	}

	// if function calls to this point, we know they are equal...
	return true
}

// returns bool to whether a matrix is approximately equal to another matrix under a given threshold ε
func (m *Matrix) IsApproxEqualTo(a *Matrix, ε float64) bool {
	if cap(*m) != cap(*a) || cap((*m)[0]) != cap((*a)[0]) {
		// matrix sizes are not equal
		return false
	}

	// now we know the matrices have the same size
	for i,_ := range *m {
		for j,_ := range (*m)[0] {
			// if entries are different return false
			diff := abs((*m)[i][j] - (*a)[i][j])
			if diff > ε {return false}
		}
	}

	// if function calls to this point, we know they are equal...
	return true
}


func (m *Matrix) Copy() Matrix {
	// allocate the matrix
	newMatrix := make(Matrix, cap(*m))
	for i,_ := range *m {
		newMatrix[i] = make([]float64, cap((*m)[i]))
	}

	// assign values to the matrix
	for i,_ := range *m {
		for j,_ := range (*m)[0] {
			newMatrix[i][j] = (*m)[i][j]
		}
	}

	return newMatrix
}


// returns the solution vector of a matrix from gaussian elimination
// as an array of integers, along with an error if there was any
// modifies original matrix to be gauss-reduced to an upper triangle
/* requires that the matrix already be an augmented square */
func (a Matrix) Gauss() ([]float64, error) {
	if cap((a)[0]) != cap(a)+1 {
		return nil, errors.New("Matrix is not an augmented square")
	}
	//fmt.Println(a)
	n := cap(a)

	for i:=0 ; i < n-1 ; i++ {
		// find the pivot, and move around matrix such that leading 
		// value in pivot row is not 0
		if a[i][i] == 0 {
			//fmt.Println("Zero")
			broken := false
			for r:=i ; r < n-1 ; r++ {
				if !broken {
					if a[r][i] != 0 {
						a[i], a[r] = a[r], a[i]
						broken = true
						continue
					}
				}
			}
			if !broken {return nil, errors.New("Matrix is singular")}
		}

		for j:=i+1 ; j < n ; j++ {
			ratio := a[j][i] / a[i][i]
			for k:=i ; k < n+1 ; k++ {
				a[j][k] -= ratio * a[i][k]
			}
		}
	}
	//fmt.Println("Gaussian Row Reduced: ",a)

	// back substitute
	b := make([]float64, cap(a))
	for i:=n-1 ; i > -1 ; i-- {
		if (a[i][i] == 0 && a[i][n] != 0) || (a[i][i] != 0 && a[i][n] == 0) {
			return nil, errors.New("Matrix inconsistant")
		} else if a[i][i] == 0 && a[i][n] == 0 {
			return nil, errors.New("Matrix has no solutions")
		}


		b[i] = a[i][n] / a[i][i]
		//fmt.Println("Matrix: ", a)
		//fmt.Println("Solutions: ", b)
		for j:=0 ; j < i ; j++ {
			a[j][n] -= a[j][i] * b[i]
			a[j][i] = 0
		}
	}
	//fmt.Println("Solutions: ",b)
	return b,nil
}

