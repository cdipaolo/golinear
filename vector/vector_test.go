package vector

import (
	"testing"
	"errors"
	"math"
	)

type vecTest struct {
	v Vector
	approxEq Vector
	approxBool bool
	equal Vector
	equalBool bool
	dot float64
	dotError error
	cross float64
	crossError error
	multTestVec Vector
	angle float64
	angleError error
	plusOne Vector
	minusOne Vector
	abs float64
	scaleBy10 Vector
}

var vectors = []vecTest{
	{ 
		Vector{0,1,2}, 
		Vector{1e-9,1+1e-9,2+1e-9}, 
		true,
		Vector{0,1,2+1e-9},
		false,
		14.0,
		nil,
		7.348469228349534,
		nil,
		Vector{3,4,5},
		0.483361282211952,
		nil,
		Vector{1,2,3},
		Vector{-1,0,1},
		2.23606797749979,
		Vector{0,10,20}},
	{ 
		Vector{10,50}, 
		Vector{20,500}, 
		false,
		Vector{10,50},
		true,
		0.0,
		errors.New("Vectors not of same order"),
		0.0,
		errors.New("Vectors not of same order"),
		Vector{10,50,50},
		0.0,
		errors.New("Vectors not of same order"),
		Vector{11,51},
		Vector{9,49},
		50.99019513592785,
		Vector{100,500}},
	{ // test empty vector
		Vector{}, 
		Vector{10,1e-7}, 
		false,
		Vector{},
		true,
		0,
		nil,
		0,
		nil,
		Vector{},
		0.0,
		nil,
		Vector{},
		Vector{},
		0.0,
		Vector{}},
	{ 
		Vector{1,2,3,4,5}, 
		Vector{1,2,3-1e-9,4,5}, 
		true,
		Vector{1,2,0,4,5},
		false,
		115,
		nil,
		28.28427124746,
		nil,
		Vector{5,6,7,8,9},
		0.241163455304096,
		nil,
		Vector{2,3,4,5,6},
		Vector{0,1,2,3,4},
		7.416198487095663,
		Vector{10,20,30,40,50}},
}

func TestEqual(t *testing.T) {
	for _,pair := range vectors {
		if pair.v.ApproxEquals(pair.approxEq) != pair.approxBool {
			t.Error(
				"| Approx Equals | expected", pair.approxBool,
				"got", pair.v.ApproxEquals(pair.approxEq),
				"from vector", pair.v,
				)
		}
		if pair.v.Equals(pair.equal) != pair.equalBool {
			t.Error(
				"| Equals | expected", pair.equalBool,
				"got", pair.v.Equals(pair.equal),
				"from vector", pair.v,
				)
		}
	}
}

func TestDot(t *testing.T) {
	for _,pair := range vectors {
		dot, err := pair.v.Dot(pair.multTestVec)
		if err != nil && pair.crossError == nil {
			t.Error("| Dot |", err)
		}
		if math.Abs(dot - pair.dot) > 1e-8 {
			t.Error(
				"| Dot | expected", pair.dot,
				"got", dot,
				"from vector", pair.v,
				)
		}
	}
}

func TestCross(t *testing.T) {
	for _,pair := range vectors {
		cross, err := pair.v.Cross(pair.multTestVec)
		if err != nil && pair.crossError == nil {
			t.Error("| Cross |", err)
		}
		if math.Abs(cross - pair.cross) > 1e-8 {
			t.Error(
				"| Cross | expected", pair.cross,
				"got", cross,
				"from vector", pair.v,
				)
		}
	}
}

func TestAngle(t *testing.T) {
	for _,pair := range vectors {
		angle, err := pair.v.Angle(pair.multTestVec)
		if err != nil && pair.angleError == nil {
			t.Error("| Angle |", err)
		}
		if math.Abs(angle - pair.angle) > 1e-8 {
			t.Error(
				"| Angle | expected", pair.angle,
				"got", angle,
				"from vector", pair.v,
				)
		}
	}
}

func TestPlusMinus(t *testing.T) {
	for _,pair := range vectors {
		one := make(Vector, cap(pair.v))
		for i,_ := range one {
			one[i] = 1
		}

		v := pair.v.Copy()
		w := v.Copy()
		plusError := v.Plus(one)
		minusError := w.Minus(one)
		if plusError != nil {
			t.Error(
				"| Plus |", plusError,
				)
		}
		if minusError != nil {
			t.Error(
				"| Minus |", minusError,
				)
		}
		if !v.Equals(pair.plusOne) {
			t.Error(
				"| Plus | expected", pair.plusOne,
				"got", v,
				"from vector", pair.v,
				)
		}
		if !w.Equals(pair.minusOne) {
			t.Error(
				"| Minus | expected", pair.minusOne,
				"got", w,
				"from vector", pair.v,
				)
		}
	}
}

func TestAbs(t *testing.T) {
	for _,pair := range vectors {
		if math.Abs(pair.v.Abs() - pair.abs) > 1e-8 {
			t.Error(
				"| Abs | expected", pair.abs,
				"got", pair.v.Abs(),
				"from vector", pair.v,
				)
		}
	}
}

func TestScale(t *testing.T) {
	for _,pair := range vectors {
		v := pair.v.Copy()
		v.Scale(10)

		if !v.Equals(pair.scaleBy10) {
			t.Error(
				"| Scale | expected", pair.scaleBy10,
				"got", v,
				"from vector", pair.v,
				)
		}
	}
}

func TestCreateVector(t *testing.T) {
	for i:=0 ; i < 10 ; i++ {
		v := CreateVector(i)
		if cap(v) != i {
			t.Error(
				"| CreateVector | expected", i,
				"vector capacity got",cap(v),
				)
		}
	}
}
