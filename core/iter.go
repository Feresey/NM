package core

import (
	"errors"
	"math"

	"gonum.org/v1/gonum/mat"
)

func makeEquivalent(matrix mat.Matrix, col mat.Vector) (res *mat.Dense, beta *mat.VecDense) {
	n, m := matrix.Dims()
	res = mat.NewDense(n, m, nil)
	beta = mat.NewVecDense(col.Len(), nil)

	res.Copy(matrix)
	beta.CopyVec(col)

	for line := 0; line < n; line++ {
		div := res.At(line, line)

		mLine := res.RowView(line).(*mat.VecDense)

		mLine.ScaleVec(-1, mLine)
		mLine.ScaleVec(1/div, mLine)

		res.Set(line, line, 0) // диагональ должна быть нулевой
		beta.SetVec(line, beta.AtVec(line)/div)
	}

	return res, beta
}

func Iterations(
	matrix mat.Matrix,
	col mat.Vector,
	eps float64,
) (res *mat.VecDense, iterations int, err error) {
	res = mat.NewVecDense(col.Len(), nil)
	m, n := matrix.Dims()

	if n != m || n != res.Len() {
		return nil, 0, IncorrectColoumn
	}

	eq, beta := makeEquivalent(matrix, col)

	alpha := mat.Norm(eq, 1)
	if alpha > 1 { // обычная линейная норма
		return nil, 0, errors.New("матрица не сходится по методу итераций")
	}

	res.CopyVec(beta)

	var (
		prevNorm = mat.Norm(res, 2) // среднеквадратичная норма
		currNorm float64
	)

	for math.Abs(prevNorm-currNorm) > eps {
		iterations++

		var (
			curr = mat.NewVecDense(res.Len(), nil)
			idx  int
		)

		for line := 0; line < n; line++ {
			for i := 0; i < m; i++ {
				curr.SetVec(idx, curr.AtVec(idx)+eq.At(line, i)*res.AtVec(i))
			}

			curr.SetVec(idx, curr.AtVec(idx)+beta.AtVec(idx))
			idx++
		}

		res = curr
		prevNorm, currNorm = currNorm, mat.Norm(res, 2)
	}

	return res, iterations, nil
}
