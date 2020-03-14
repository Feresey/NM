package core

import (
	"errors"
	"math"
)

// getMaxAbsElem : возвращает индексы максимального по абсолютному значению элемента верхнего треугольника матрицы
func getMaxAbsElem(m *Matrix) (int, int) {
	var (
		max float64
		i   = -1
		j   = -1
	)

	for line := 0; line < m.n; line++ {
		lineIter := line * m.m

		for col := line + 1; col < m.m; col++ {
			if tmp := math.Abs(m.data[lineIter+col]); tmp > max {
				max = tmp
				i = line
				j = col
			}
		}
	}

	return i, j
}

func getOrt(m *Matrix, i, j int) *Matrix {
	var (
		U     = EMatrix(m.n)
		lineI = i * U.m
		lineJ = j * U.m
		phi   = 0.5 * math.Atan(2*m.data[lineI+j]/(m.data[lineI+i]-m.data[lineJ+j]))
		sin   = math.Sin(phi)
		cos   = math.Cos(phi)
	)

	U.data[lineI+j] = -sin
	U.data[lineJ+i] = sin
	U.data[lineI+i] = cos
	U.data[lineJ+j] = cos

	return U
}

func getSquareSum(m *Matrix) (sum float64) {
	for line := 0; line < m.n; line++ {
		lineIter := line * m.m

		for col := line + 1; col < m.m; col++ {
			elem := m.data[lineIter+col]
			sum += elem * elem
		}
	}

	return math.Sqrt(sum)
}

func Rotations(matrix *Matrix, eps float64) (Coloumn, *Matrix, int, error) {
	if matrix.n != matrix.m {
		return nil, nil, 0, IncorrectColoumn
	}

	matrix = matrix.Copy()

	var (
		sz         = make(Coloumn, 0, matrix.n)
		sv         = EMatrix(matrix.n)
		iterations int
	)

	for {
		iterations++

		i, j := getMaxAbsElem(matrix)

		if i == -1 || j == -1 {
			return nil, nil, iterations, errors.New("вырожденная матрица")
		}

		U := getOrt(matrix, i, j)

		sv = sv.ProdMatrix(U)
		matrix = Transponse(U).ProdMatrix(matrix.ProdMatrix(U))

		if getSquareSum(matrix) < eps {
			break
		}
	}

	for line := 0; line < len(matrix.data); line += matrix.m + 1 {
		sz = append(sz, matrix.data[line])
	}

	return sz, sv, iterations, nil
}
