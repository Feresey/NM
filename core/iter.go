package core

import (
	"errors"
	"math"
)

func (matrix *Matrix) norm() (res float64) {
	for line := 0; line < len(matrix.data); line += matrix.m {
		var max float64

		for i := 0; i < matrix.m; i++ {
			max += math.Abs(matrix.data[line+i])
		}

		if max > res {
			res = max
		}
	}

	return
}

func norm(data []float64) (res float64) {
	for _, val := range data {
		res += val * val
		// if max := math.Abs(val); max > res {
		// 	res = max
		// }
	}

	return math.Sqrt(res)
}

func makeEquivalent(m *Matrix, col Coloumn) (*Matrix, Coloumn) {
	var (
		matrix = m.Copy()
		line   = 0
		beta   = col.Copy()
	)

	for i := 0; i < matrix.n; i++ {
		div := matrix.data[line+i]

		for elem := 0; elem < matrix.m; elem++ {
			matrix.data[line+elem] = -matrix.data[line+elem] / div
		}

		matrix.data[line+i] = 0 // диагональ должна быть нулевой
		beta[i] /= div
		line += matrix.m
	}

	return matrix, beta
}

func Iterations(matrix *Matrix, col Coloumn, eps float64) (Coloumn, int, error) {
	if matrix.n != matrix.m {
		return nil, 0, IncorrectColoumn
	}

	m, beta := makeEquivalent(matrix, col)

	if m.norm() > 1 {
		return nil, 0, errors.New("матрица не сходится по методу итераций")
	}

	res := beta.Copy()

	var (
		prevNorm   = norm(res)
		currNorm   float64
		iterations int
	)

	for math.Abs(prevNorm-currNorm) > eps {
		iterations++

		var (
			curr = make(Coloumn, len(res))
			idx  = 0
		)

		for line := 0; line < len(m.data); line += m.m {
			for i := 0; i < m.m; i++ {
				curr[idx] += m.data[line+i] * res[i]
			}

			curr[idx] += beta[idx]
			idx++
		}

		res = curr
		prevNorm, currNorm = currNorm, norm(res)
	}

	return res, iterations, nil
}

func Zeidel(matrix *Matrix, col Coloumn, eps float64) (Coloumn, int, error) {
	if matrix.n != matrix.m {
		return nil, 0, IncorrectColoumn
	}

	m, beta := makeEquivalent(matrix, col)

	if m.norm() > 1 {
		return nil, 0, errors.New("матрица не сходится по методу итераций")
	}

	res := beta.Copy()

	var (
		prevNorm   = norm(res)
		currNorm   float64
		iterations int
	)

	for ; math.Abs(prevNorm-currNorm) > eps; iterations++ {
		curr := beta.Copy()

		for i := 0; i < m.n; i++ {
			for j := 0; j < i; j++ {
				curr[i] += curr[j] * m.At(i, j)
			}
			for j := i; j < m.m; j++ {
				curr[i] += res[j] * m.At(i, j)
			}
		}

		res = curr
		prevNorm, currNorm = currNorm, norm(res)
	}

	return res, iterations, nil
}
