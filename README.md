#golinear
Golang implementation of linear algebra manipulations using float64 representations of vectors and matricies. To look at usage for available functions, just search for the comments and arguments 

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
After importing, you can initialize vectors or matricies using ```make(Vector,capacity)``` or ```make(Matrix,capacity)``` and, for matricies, individually allocating the rows. Or using the libraries functions as follows:

```golang
// create a zero-filled 4x4 matrix
matrix := CreateMatrix(4,4)
// create a zero-filled 9 dimensional vector
vector := CreateVector(9)
```

#Liscense - MIT
See LISCENCE.