package core

import (
	"fmt"
	"strings"
)

// LUDecomposition : Разделяет матрицу на три:
// L - нижнетреугольную с еденицами на главной диогонали
// U - верхнетреугольная
// P - матрица перестановок (опциональная)
func LUDecomposition(matrix *Matrix) *LUP {
	if matrix == nil || matrix.n != matrix.m {
		return nil
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

			for i := 0; i < U.n; i++ {
				U.data[processLine+i] -= U.data[currLine+i] * coeff
			}
		}
	}

	for line := 0; line < len(L.data); line += L.m + 1 {
		L.data[line] = 1
	}

	return &LUP{L, U, P, matrix.n, matrix.m}
}

func (lup *LUP) SolveSLAU(b Coloumn) Coloumn {
	if len(b) != lup.n {
		return nil
	}
	tmp := lup.P.ProdMatrix(&Matrix{data: Row(b), n: lup.n, m: 1})
	b = Coloumn(tmp.data)
	var (
		x = make(Coloumn, lup.n)
		y = make(Coloumn, lup.n)
	)

	for i := 0; i < lup.n; i++ {
		var (
			num  float64
			line = i * lup.m
		)
		for j := 0; j < i; j++ {
			num += lup.L.data[line+j] * y[j]
		}
		y[i] = b[i] - num
	}

	for i := lup.n - 1; i >= 0; i-- {
		var (
			num  float64
			line = i * lup.m
		)
		for j := i + 1; j < lup.n; j++ {
			num += lup.U.data[line+j] * x[j]
		}
		x[i] = (y[i] - num) / lup.U.data[line+i]
	}

	return x
}

func (lup *LUP) Determinant() float64 {
	if len(lup.P.data) == 0 {
		return 0
	}
	var res float64 = 1
	for line := 0; line < len(lup.U.data); line += lup.m + 1 {
		res *= lup.U.data[line]
	}

	var (
		used = make([]bool, lup.n)
		prod = false
		line = 0
	)
	for i := 0; i < lup.m; i++ {
		idx := lup.P.findNotZeroIndexInCol(i)
		if idx != i && !used[i] {
			prod = !prod
		}
		used[i] = true
		used[idx] = true
		line += lup.m + 1
	}

	if prod {
		res *= -1
	}
	return res
}

func (lup *LUP) Inverse() *Matrix {
	res := NewMatrix(lup.n, lup.m)

	for i := 0; i < lup.m; i++ {
		col := make(Coloumn, lup.n)
		col[i] = 1
		resCol := lup.SolveSLAU(col)

		line := i
		for _, val := range resCol {
			res.data[line] = val
			line += lup.m
		}
	}

	return res
}

func (d DisplaySLAU) String() string {
	b := strings.Builder{}

	for i := 0; i < d.n; i++ {
		line := i * d.m
		b.WriteString(fmt.Sprintf("%6.2f*x%d", d.data[line], 1))

		for j := 1; j < d.m; j++ {
			b.WriteString(fmt.Sprintf(" +%6.2f*x%d", d.data[line+j], j+1))
		}
		b.WriteString(fmt.Sprintf(" = %6.2f\n", d.Coloumn[i]))
	}

	return b.String()
}
