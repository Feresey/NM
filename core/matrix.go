package core

import (
	"errors"
	"fmt"
	"strings"
)

// Get : return elem
func (matrix *Matrix) Get(i, j int) float64 {
	return matrix.data[i*matrix.m+j]
}

// Set : set elem
func (matrix *Matrix) Set(i, j int, value float64) {
	matrix.data[i*matrix.m+j] = value
}

// Copy :
func (matrix *Matrix) Copy() *Matrix {
	res := &Matrix{
		data: make([]float64, matrix.n*matrix.m),
		n:    matrix.n,
		m:    matrix.m,
	}
	copy(res.data, matrix.data)

	return res
}

func (matrix Matrix) String() string {
	b := strings.Builder{}

	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			b.WriteString(fmt.Sprintf("%f ", matrix.data[i*matrix.m+j]))
		}
		b.WriteString("\n")
	}

	return b.String()
}

// EMatrix : создает единичную матрицу размера nxn
func EMatrix(n int) *Matrix {
	var (
		E   = NewMatrix(n, n)
		row int
	)

	for i := 0; i < n; i++ {
		E.data[row+i] = 1
		row += n
	}

	return E
}

// ProdMatrix : перемножает матрицы
func (matrix *Matrix) ProdMatrix(right *Matrix) (*Matrix, error) {
	if matrix.m != right.n {
		return nil, errors.New("Не совпадают размерности матриц")
	}
	res := NewMatrix(matrix.n, right.m)

	for i := 0; i < matrix.n; i++ {
		for j := 0; j < right.m; j++ {
			index := i*res.m + j
			leftLine := i * matrix.m
			rightColoumn := j

			//matrix.m == right.n
			for k := 0; k < right.n; k++ {
				res.data[index] += matrix.data[leftLine+k] * right.data[rightColoumn]
				rightColoumn += right.m
			}
		}
	}

	return res, nil
}

// SwapLines : swap lines a and b
func (matrix *Matrix) SwapLines(a, b int) {
	if a == b {
		return
	}

	var (
		lineA = a * matrix.m
		lineB = b * matrix.m
	)

	for i := 0; i < matrix.m; i++ {
		tmpA, tmpB := lineA+i, lineB+i
		matrix.data[tmpA], matrix.data[tmpB] = matrix.data[tmpB], matrix.data[tmpA]
	}
}

func (matrix *Matrix) findNotZeroIndexInCol(idx int, P *Matrix) error {
	line := idx * matrix.m

	for i := idx; i < matrix.n; i++ {
		if matrix.data[line+i] != 0 {
			matrix.SwapLines(idx, i)
			P.SwapLines(idx, i)
			return nil
		}
	}

	return errors.New("Вырожденная матрица")
}

// LUDecomposition : Разделяет матрицу на три:
// L - нижнетреугольную с еденицами на главной диогонали
// U - верхнетреугольная
// P - матрица перестановок (опциональная)
func (matrix *Matrix) LUDecomposition() (*Matrix, *Matrix, *Matrix, error) {
	if matrix.n != matrix.m {
		return nil, nil, nil, errors.New("Матрица не квадратная")
	}

	var (
		U = matrix.Copy()
		L = EMatrix(U.n)
		P = EMatrix(U.n)
	)

	for col := 0; col < U.m; col++ {
		err := U.findNotZeroIndexInCol(col, P)
		if err != nil {
			return nil, nil, nil, err
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

			for i := 0; i < U.m; i++ {
				U.data[processLine+i] -= U.data[currLine+i] * coeff
			}
		}
	}

	return L, U, P, nil
}
