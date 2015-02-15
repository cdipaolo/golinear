package vector

import (
	"errors"
	"math"
)

type Vector []float64

// copies a Vector and returns the copy
func (z Vector) Copy() Vector {
	v := make(Vector, cap(z))
	for i, val := range z {
		v[i] = val
	}
	return v
}

// creates a zero-filled vector of capacity given in the arguments
func CreateVector(i int) Vector {
	return make(Vector, i)
}

// returns a bool as to whether Vector a is almost (very) equal to Vector b
func (a Vector) ApproxEquals(b Vector) bool {
	if cap(a) != cap(b) {return false}
	ε := 1e-8
	for i,_ := range a {
		diff := math.Abs(a[i] - b[i])
		if diff > ε {return false}
	}
	return true
}

// returns a bool as to whether a vector is strictly equal to another
func (a Vector) Equals(b Vector) bool {
	if cap(a) != cap(b) {return false}
	for i,_ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// returns the dot product of a vector and another vector, along with an error
func (a Vector) Dot(b Vector) (float64, error) {
	if cap(a) != cap(b) {return 0, errors.New("Vectors not of same order") }
	var c float64
	for i,_ := range a {
		c += a[i] * b[i]
	}
	return c, nil
}

// returns the magnitude of the cross product of two vectors, along with an error
/* Cross product is only valid in 0,1,2,3, and 7 dimensional space */
func (a Vector) Cross(b Vector) (float64, error) {
	if cap(a) != cap(b) {
		return 0, errors.New("Vectors not of same order")
	}
	theta, err := a.Angle(b)
	return a.Abs() * b.Abs() * math.Sin(theta), err
}

// returns the angle between two vectors calculated with the dot product
func (a Vector) Angle(b Vector) (float64, error) {
	newA := a.Copy()
	dot, err := newA.Dot(b)
	if err != nil || cap(a) == 0 {
		return 0, err
	}

	cosTheta := dot / (a.Abs() * b.Abs())
	theta := math.Acos(cosTheta)
	return theta, nil
}

// returns the addition between two vectors and error if there is one
func (a Vector) Plus(b Vector) error {
	if cap(a) != cap(b) {return errors.New("Vectors not of same order")}
	for i, _ := range a {
		a[i] += b[i]
	}
	return nil
}

// returns the subtraction of the argument vector from the called, and an error
func (a Vector) Minus(b Vector) error {
	c := b.Copy()
	c.Scale(-1)
	return a.Plus(c)
}

// returns the magnitude of a vector (or absulute value)
func (a Vector) Abs() float64 {
	var sum float64
	for _,val := range a {
		sum += val * val
	}
	return math.Sqrt(sum)
}

// scales a vector by a scalar
func (a Vector) Scale(u float64) {
	for i,_ := range a {
		a[i] *= u
	}
}