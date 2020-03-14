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
		if max := math.Abs(val); max > res {
			res = max
		}
	}

	return
}

func Iterations(matrix *Matrix, col Coloumn, eps float64) (Coloumn, int, error) {
	if matrix.n != matrix.m {
		return nil, 0, IncorrectColoumn
	}

	var (
		beta = make(Coloumn, len(col))
		line = 0
	)

	copy(beta, col)

	matrix = matrix.Copy()

	for i := 0; i < matrix.n; i++ {
		div := matrix.data[line+i]

		for elem := 0; elem < matrix.m; elem++ {
			matrix.data[line+elem] = -matrix.data[line+elem] / div
		}

		matrix.data[line+i] = 0 // диагональ должна быть нулевой
		beta[i] /= div
		line += matrix.m
	}

	if matrix.norm() > 1 {
		return nil, 0, errors.New("матрица не сходится по методу итераций")
	}

	res := make(Coloumn, len(beta))
	copy(res, beta)

	var (
		prevNorm   = norm(res)
		currNorm   float64
		iterations int
	)

	for math.Abs(prevNorm-currNorm) > eps {
		iterations++

		curr := make(Coloumn, len(res))
		idx := 0

		for line := 0; line < len(matrix.data); line += matrix.m {
			for i := 0; i < matrix.m; i++ {
				curr[idx] += matrix.data[line+i] * res[i]
			}

			curr[idx] += beta[idx]
			idx++
		}

		res = curr
		prevNorm, currNorm = currNorm, norm(res)
	}

	return res, iterations, nil
}
