package matrix

import "errors"
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

// evaluates if two arrays of float64's are very close to equal
func solutionApproxEquals(a []float64, b []float64) bool {
	if cap(a) != cap(b) {return false}
	ε := 1e-8
	for i,_ := range a {
		diff := abs(a[i] - b[i])
		if diff > ε {return false}
	}
	return true
}


// returns a boolean to whether a matrix is of the same size and strictly equal to another matrix
func (m *Matrix) Equals(a *Matrix) bool {
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
func (m *Matrix) ApproxEquals(a *Matrix, ε float64) bool {
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
	n := cap(a)

	for i:=0 ; i < n-1 ; i++ {
		// find the pivot, and move around matrix such that leading 
		// value in pivot row is not 0
		if a[i][i] == 0 {
			broken := false
			for r:=i+1 ; r < n ; r++ {
				if !broken {

					if a[r][i] != 0 {
						a[i], a[r] = a[r], a[i]
						broken = true
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
	// back substitute
	b := make([]float64, cap(a))
	for i:=n-1 ; i > -1 ; i-- {
		if a[i][i] == 0 && a[i][n] != 0 {
			return nil, errors.New("Matrix inconsistant")
		} else if a[i][i] == 0 && a[i][n] == 0 {
			return nil, errors.New("Matrix has infinitely many solutions")
		}

		if a[i][i] != 0 && a[i][n] == 0 {
			b[i] = 0
		} else {
			b[i] = a[i][n] / a[i][i]
		}

		for j:=0 ; j < i ; j++ {
			a[j][n] -= a[j][i] * b[i]
			a[j][i] = 0
		}
	}
	return b,nil
}


// called on a square matrix and takes the solution matrix 'b' as an argument
// returns an array of float64's representing the solution matrix [x0,x1,x2,...]
func (a Matrix) Solution(b []float64) (x []float64,err error) {
	// check if matrix is a square
	if cap(a) != cap(a[0]) {
		return nil, errors.New("Matrix is not a square")
	}

	n := cap(a)

	// create the augmented matrix
	m := make(Matrix, n)
	for i,_ := range m {
		m[i] = make([]float64, n+1)
	}
	for i:=0 ; i < n ; i++ {
		for j:=0 ; j <= n ; j++ {
			if j < n{
				m[i][j] = a[i][j]
			} else {
				m[i][j] = b[i]
			}
		}
	}
	x,err = m.Gauss()
	return
}

func (a Matrix) Rows() int {
	return cap(a)
}

func (a Matrix) Columns() int {
	return cap(a[0])
}

// scales a matrix by a scalar
func (a Matrix) Scale(n float64) {
	for i, _ := range a {
		for j, _ := range a[0] {
			a[i][j] *= n
		}
	}
}

// returns an identity matrix of size n x n
func Identity(n int8) Matrix {
	a := make(Matrix,n)
	for i,_ := range a {
		a[i] = make([]float64, n)
	}

	// assign the values
	for i:=0 ; i < int(n) ; i++ {
		a[i][i] = 1
	}

	return a
}

