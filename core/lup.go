package core

import (
	"errors"
	"fmt"
	"strings"
)

// LUDecomposition : Разделяет матрицу на три:
// L - нижнетреугольную с еденицами на главной диогонали
// U - верхнетреугольная
// P - матрица перестановок (опциональная)
func LUDecomposition(matrix *Matrix) (*LUP, error) {
	if matrix == nil || matrix.n != matrix.m {
		return nil, errors.New("Матрица не квадратная")
	}

	var (
		L = NewMatrix(matrix.n, matrix.m)
		U = matrix.Copy()
		P = EMatrix(matrix.n)
	)

	for col := 0; col < U.m; col++ {
		idx := U.findNotZeroIndexInCol(col)
		L.SwapLines(col, idx)
		U.SwapLines(col, idx)
		P.SwapLines(col, idx)

		var (
			elem     = U.data[col*U.m+col]
			currLine = col * U.m
		)

		for line := col + 1; line < U.n; line++ {
			var (
				processLine        = line * U.m
				processLineWithCol = processLine + col
				coeff              = U.data[processLineWithCol] / elem
			)
			L.data[processLineWithCol] = coeff

			for i := 0; i < U.m; i++ {
				U.data[processLine+i] -= U.data[currLine+i] * coeff
			}
		}
	}

	for line := 0; line < len(L.data); line += L.m + 1 {
		L.data[line] = 1
	}

	return &LUP{l: L, u: U, p: P, n: matrix.n, m: matrix.m}, nil
}

func (lup *LUP) SolveSLAU(b Row) (Row, error) {
	if len(b) != lup.n {
		return nil, errors.New("Неверная размерность столбца")
	}
	tmp, _ := lup.p.ProdMatrix(&Matrix{data: b, n: lup.n, m: 1})
	b = tmp.data
	var (
		x = make(Row, lup.n)
		y = make(Row, lup.n)
	)

	for i := 0; i < lup.n; i++ {
		var (
			num  float64
			line = i * lup.m
		)
		for j := 0; j < i; j++ {
			num += lup.l.data[line+j] * y[j]
		}
		y[i] = b[i] - num
	}

	for i := lup.n - 1; i >= 0; i-- {
		var (
			num  float64
			line = i * lup.m
		)
		for j := i + 1; j < lup.n; j++ {
			num += lup.u.data[line+j] * x[j]
		}
		x[i] = (y[i] - num) / lup.u.Get(i, i)
	}

	return x, nil
}

func (lup *LUP) Determinant() float64 {
	lineIter := 0
	var res float64 = 1
	for lineIter < len(lup.u.data) {
		res *= lup.u.data[lineIter]
		lineIter += lup.m + 1
	}
	return res
}

func (d DisplaySLAU) String() string {
	b := strings.Builder{}

	for i := 0; i < d.n; i++ {
		line := i * d.m
		for j := 0; j < d.m; j++ {
			b.WriteString(fmt.Sprintf("%6.2f*x%d ", d.data[line+j], j+1))
		}
		b.WriteString(fmt.Sprintf(" = %f\n", d.Row[i]))
	}

	return b.String()
}
