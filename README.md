#golinear ![golinear](golinear2.png)
Golang implementation of linear algebra manipulations using ```float64``` representations of vectors and matricies. To look at usage for available functions, just search for the comments and arguments 

#Installation
####Step 1
In your Golang project's ```src``` directory, clone the repository.
```
cd /path/to/my/project/src
git clone https://github.com/cdipaolo96/golinear.git
```
###Step 2
In *your* Go source files, import the packages you want to use.

```golang
import (
	"golinear/matrix"
	"golinear/vector"
	)
```

#Basic Usage
The library's representation of matricies and vectors are just one or two dimensional arrays of ```float64```'s. After importing, you can initialize vectors or matricies using ```make(Vector,capacity)``` or ```make(Matrix,capacity)``` and, for matricies, individually allocating the rows. Or using the library's functions as follows:

```golang
// create a zero-filled 4x4 matrix
matrix := CreateMatrix(4,4)
// create a zero-filled 9 dimensional vector
vector := CreateVector(9)
```

#Available Functions
Note– these are not in any specific order.

Matrix  |  Vector
:--------:|:---------:
```CreateMatrix(m uint8, n uint8) Matrix``` | ```CreateVector(i uint8) Vector```
```Identity(n int8) Matrix``` | ```(a Vector) Scale(u float64)```
```(m *Matrix) Equals(a *Matrix) bool``` | ```(a Vector) Equals(b Vector) bool```
```(m *Matrix) ApproxEquals(a *Matrix, ε float64) bool``` | ```(a Vector) ApproxEquals(b Vector) bool```
```(m *Matrix) Copy() Matrix``` | ```(z Vector) Copy() Vector```
```(a Matrix) Rows() int``` | ```(a Vector) Dot(b Vector) (float64, error)```
```(a Matrix) Columns() int``` | ```(a Vector) Cross(b Vector) (float64, error)```
```(a Matrix) Scale(n float64)``` | ```(a Vector) Angle(b Vector) (float64, error)```
```(a Matrix) Gauss() ([]float64, error)``` | ```(a Vector) Plus(b Vector) error```
```(a Matrix) Solution(b []float64) (x []float64,err error)``` | ```(a Vector) Minus(b Vector) error```
 | ```(a Vector) Abs() float64```



#Liscense - MIT
See LISCENCE.