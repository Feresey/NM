package core

import (
	"fmt"
	"math"
	"strings"

	"gonum.org/v1/gonum/mat"
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

func StringMatrix(matrix mat.Matrix) string {
	b := strings.Builder{}

	n, m := matrix.Dims()

	for line := 0; line < n; line++ {
		for col := 0; col < m; col++ {
			b.WriteString(fmt.Sprintf("%10.5f", matrix.At(line, col)))
		}

		b.WriteString("\n")
	}

	return b.String()
}

func (matrix Matrix) String() string {
	b := strings.Builder{}

	for line := 0; line < len(matrix.data); line += matrix.m {
		for _, val := range matrix.data[line : line+matrix.m] {
			b.WriteString(fmt.Sprintf("%7.2f", val))
		}

		b.WriteString("\n")
	}

	return b.String()
}

// ProdMatrix : перемножает матрицы и возвращает результат перемножения
func (matrix *Matrix) ProdMatrix(right *Matrix) *Matrix {
	if matrix.m != right.n {
		return nil
	}

	res := NewMatrix(matrix.n, right.m)

	for i := 0; i < matrix.n; i++ {
		for j := 0; j < right.m; j++ {
			var (
				index        = i*res.m + j
				leftLine     = i * matrix.m
				rightColoumn = j
			)

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

func (matrix *Matrix) maxInCol(col, from int) int {
	var (
		res = -1
		max = eps
	)

	for i := from; i < matrix.n; i++ {
		if tmp := math.Abs(matrix.data[matrix.m*i+col]); tmp > max {
			max = tmp
			res = i
		}
	}

	return res
}
