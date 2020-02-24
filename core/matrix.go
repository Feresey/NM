package core

import (
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

// GetSize : return sizes
func (matrix *Matrix) GetSize() (rows, coloumns int) {
	return matrix.n, matrix.m
}

func (matrix Matrix) String() string {
	b := strings.Builder{}

	for i := 0; i < matrix.n; i++ {
		for j := 0; j < matrix.m; j++ {
			b.WriteString(fmt.Sprintf("%7.2f", matrix.data[i*matrix.m+j]))
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
func (matrix *Matrix) ProdMatrix(right *Matrix) *Matrix {
	if matrix.m != right.n {
		return nil
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

	return res
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

func (matrix *Matrix) findNotZeroIndexInCol(idx int) int {
	line := idx + matrix.m*idx

	for i := idx; i < matrix.n; i++ {
		if matrix.data[line] != 0 {
			return i
		}
		line += matrix.m
	}

	return idx
}
