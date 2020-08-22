package core

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	IncorrectColoumn = errors.New("Размерность столбца не совпадает")

	eps = 1e-9
)

type (
	Row     []float64
	Coloumn []float64
)

func (r Row) Copy() Row {
	res := make(Row, len(r))
	copy(res, r)
	return res
}

func (c Coloumn) Copy() Coloumn {
	res := make(Coloumn, len(c))
	copy(res, c)
	return res
}

// Matrix : type for process core actions
type Matrix struct {
	data []float64
	// Number lines
	n int
	// Number rows
	m int
}

// Get : return elem
func (matrix *Matrix) At(i, j int) float64 {
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

// NewMatrix : create a new matrix
func NewMatrix(n, m int) *Matrix {
	res := &Matrix{n: n, m: m}
	res.data = make(Row, n*m)

	return res
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

// GetSize : return sizes
func (matrix *Matrix) GetSize() (rows, coloumns int) {
	return matrix.n, matrix.m
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

func Transponse(m *Matrix) *Matrix {
	var (
		res   = NewMatrix(m.m, m.n)
		lineM = 0
	)

	for i := 0; i < m.n; i++ {
		lineRes := i

		for j := 0; j < m.m; j++ {
			res.data[lineRes] = m.data[lineM+j]
			lineRes += m.n
		}

		lineM += m.m
	}

	return res
}

func DisplaySLAU(m *Matrix, col Coloumn) string {
	b := strings.Builder{}

	for i := 0; i < m.n; i++ {
		line := i * m.m
		b.WriteString(fmt.Sprintf("%6.2f*x%d", m.data[line], 1))

		for j := 1; j < m.m; j++ {
			b.WriteString(fmt.Sprintf(" +%6.2f*x%d", m.data[line+j], j+1))
		}

		b.WriteString(fmt.Sprintf(" = %6.2f\n", col[i]))
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
