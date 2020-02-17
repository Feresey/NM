package matrix

import (
	"errors"
	"fmt"
	"strings"
)

func (matrix *Matrix) String() string {
	b := strings.Builder{}

	for i := 0; i < matrix.m; i++ {
		for j := 0; j < matrix.n; j++ {
			b.WriteString(fmt.Sprintf("%f ", matrix.data[i*matrix.m+j]))
		}
		b.WriteString("\n")
	}

	return b.String()
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
			rightColoumn := 0
			//matrix.m == right.n
			for k := 0; k < right.n; k++ {
				// todo calculate rows step by step
				res.data[index] += matrix.data[leftLine+k] * right.data[rightColoumn+j]
				rightColoumn += right.m
			}
		}
	}

	return res, nil
}

func (matrix *Matrix) findNotZeroIndexInCol(idx int) (int, error) {
	line := idx * matrix.m

	for i := 0; i < matrix.n; i++ {
		if matrix.data[line+i] != 0 {
			return i, nil
		}
	}

	return 0, errors.New("Вырожденная матрица")
}

func (matrix *Matrix) LUDecomposition() (*Matrix, *Matrix, error) {
	L := NewMatrix(matrix.n, matrix.m)
	U := NewMatrix(matrix.n, matrix.m)

	for col := 0; col < matrix.m; col++ {
		idx, err := matrix.findNotZeroIndexInCol(col)
		if err != nil {
			return nil, nil, err
		}

		for line := 0; line < matrix.n; line++ {

		}
	}

	return L, U, nil
}
