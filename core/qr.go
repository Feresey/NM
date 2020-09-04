package core

import (
	"math"
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

func QR(m *Matrix, eps float64) (*Matrix, *Matrix, error) {
	if m.n != m.m {
		return nil, nil, IncorrectColoumn
	}

	R := m.Copy()
	Q := EMatrix(m.n)
	for i := 0; i < m.n-1; i++ {
		vec := NewMatrix(m.n, 1)
		var sum float64
		for j := i; j < m.n; j++ {
			v := m.At(j, i)
			sum += v * v
		}
		diag := m.At(i, i)
		vec.Set(i, 0, diag+sign(diag)*math.Sqrt(sum))

		for j := i + 1; j < m.n; j++ {
			vec.Set(j, 0, m.At(j, i))
		}

		// fmt.Println("vec: ", mat.Formatted(vec))

		H := hausholder(vec)
		// fmt.Println("hausholder: ", mat.Formatted(H))
		R = H.ProdMatrix(R)
		Q = Q.ProdMatrix(H)
	}

	return Q, R, nil
}
