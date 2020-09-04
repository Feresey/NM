package core

import (
	"fmt"
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/mat"
)

func sign(num float64) float64 {
	if num > 0 {
		return 1
	} else {
		return -1
	}
}

func hausholder(v *Matrix) *Matrix {
	vT := Transponse(v)

	top := v.ProdMatrix(vT)
	bottom := vT.ProdMatrix(v)
	top.ProdNum(2 / bottom.At(0, 0))

	res := EMatrix(v.n)
	res.Sub(top)

	return res
}

func QR(m *Matrix) (*Matrix, *Matrix, error) {
	if m.n != m.m {
		return nil, nil, IncorrectColoumn
	}

	R := m.Copy()
	Q := EMatrix(m.n)
	for i := 0; i < m.n-1; i++ {
		vec := NewMatrix(m.n, 1)
		var sum float64
		for j := i; j < m.n; j++ {
			v := R.At(j, i)
			sum += v * v
		}
		diag := R.At(i, i)
		vec.Set(i, 0, diag+sign(diag)*math.Sqrt(sum))

		for j := i + 1; j < m.n; j++ {
			vec.Set(j, 0, R.At(j, i))
		}

		// fmt.Println("vec: \n", mat.Formatted(vec))
		H := hausholder(vec)
		// fmt.Println("hausholder: \n", mat.Formatted(H))
		R = H.ProdMatrix(R)
		// fmt.Printf("A%d:\n%v\n", i+1, mat.Formatted(R))
		Q = Q.ProdMatrix(H)
	}

	return Q, R, nil
}

func QRValues(A *Matrix, eps float64) ([]complex128, int, error) {
	res := make([]complex128, 0, A.n)
	A = A.Copy()
	var iterations int
	for i := 0; i < A.n; {
		values, err := getEigenvalue(A, eps, i, &iterations)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, values...)
		i += len(values)
	}

	return res, iterations, nil
}

func finishComplex(A *Matrix, eps float64, i int) (bool, error) {
	Q, R, err := QR(A)
	if err != nil {
		return false, err
	}
	Anext := R.ProdMatrix(Q)

	l11, l12 := GetRoots(A, i)
	l21, l22 := GetRoots(Anext, i)

	return cmplx.Abs(l11-l21) <= eps && cmplx.Abs(l12-l22) <= eps, nil
}

func getEigenvalue(A *Matrix, eps float64, i int, iter *int) ([]complex128, error) {
	for {
		Q, R, err := QR(A)
		if err != nil {
			return nil, err
		}
		A = R.ProdMatrix(Q)

		fmt.Printf("Q%d:\n%v\n", *iter, mat.Formatted(Q))
		fmt.Printf("R%d:\n%v\n", *iter, mat.Formatted(R))
		fmt.Printf("A%d:\n%v\n", *iter+1, mat.Formatted(A))
		fmt.Println()

		*iter++

		underlying := make([]float64, A.n-i-1)
		for j := i + 1; j < A.n; j++ {
			underlying[j-i-1] = A.At(j, i)
		}
		if norm(underlying) <= eps {
			return []complex128{complex(A.At(i, i), 1)}, nil
		}
		if norm(underlying[1:]) <= eps {
			finished, err := finishComplex(A, eps, i)
			if err != nil {
				return nil, err
			}

			if finished {
				r1, r2 := GetRoots(A, i)
				return []complex128{r1, r2}, nil
			}
		}
	}
}

func SolveQuadratic(a, b, c complex128) (xpos, xneg complex128) {
	negB := -b
	twoA := 2 * a
	bSquared := b * b
	fourAC := 4 * a * c
	discrim := bSquared - fourAC
	sq := cmplx.Sqrt(discrim)
	xpos = (negB + sq) / twoA
	xneg = (negB - sq) / twoA
	return
}

func GetRoots(A *Matrix, i int) (pos, neg complex128) {
	var a11, a12, a21, a22 float64

	lineMax := i + 1
	if lineMax >= A.n {
		lineMax = -1
	}
	colMax := i + 1
	if colMax >= A.m {
		colMax = -1
	}

	a11 = A.At(i, i)
	if lineMax != -1 {
		a12 = A.At(i, colMax)
	}
	if colMax != -1 {
		a21 = A.At(lineMax, i)
	}
	if lineMax != -1 && colMax != -1 {
		a22 = A.At(lineMax, colMax)
	}

	a := complex128(1)
	b := complex(-a11-a22, 0)
	c := complex(a11*a22-a12*a21, 0)

	return SolveQuadratic(a, b, c)
}
