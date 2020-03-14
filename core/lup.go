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
	if matrix.n != matrix.m {
		return nil, errors.New("матрица не квадратная")
	}

	var (
		L = NewMatrix(matrix.n, matrix.m)
		U = matrix.Copy()
		P = make([]int, matrix.n)
	)

	for i := 0; i < matrix.n; i++ {
		P[i] = i
	}

	for col := 0; col < U.m; col++ {
		idx := U.maxInCol(col, col)

		if idx == -1 {
			return nil, errors.New("матрица вырожденная")
		}

		if idx != col {
			P[col], P[idx] = P[idx], P[col]
			L.SwapLines(col, idx)
			U.SwapLines(col, idx)
		}

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

	return &LUP{L, U, P, matrix.n, matrix.m}, nil
}

func (lup *LUP) SolveSLAU(b Coloumn) (Coloumn, error) {
	if len(b) != lup.n {
		return nil, IncorrectColoumn
	}

	var (
		tmp = lup.SwapMatrix(&Matrix{data: Row(b), n: lup.n, m: 1}, true)
		x   = make(Coloumn, lup.n)
		y   = make(Coloumn, lup.n)
	)

	b = Coloumn(tmp.data)

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

	return x, nil
}

func (lup *LUP) Determinant() float64 {
	if lup.n == 0 {
		return 0
	}

	var (
		res  float64 = 1
		used         = make([]bool, lup.n)
		prod         = false
	)

	for line := 0; line < len(lup.U.data); line += lup.m + 1 {
		res *= lup.U.data[line]
	}

	for idx, val := range lup.P {
		if idx != val && !used[idx] {
			prod = !prod
		}

		used[idx] = true
		used[val] = true
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

		var (
			resCol, _ = lup.SolveSLAU(col)
			line      = i
		)

		for _, val := range resCol {
			res.data[line] = val
			line += lup.m
		}
	}

	return res
}

func (lup *LUP) SwapMatrix(m *Matrix, reverse bool) *Matrix {
	res := NewMatrix(m.n, m.m)

	for idx, val := range lup.P {
		var (
			from = m.m * idx
			to   = m.m * val
		)

		if reverse {
			to, from = from, to
		}

		copy(res.data[to:to+m.m], m.data[from:from+m.m])
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
