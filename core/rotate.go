package core

import "math"

// getMaxAbsElem : возвращает индексы максимального по абсолютному значению элемента верхнего треугольника матрицы
func getMaxAbsElem(m *Matrix) (int, int) {
	var (
		max float64
		i   = -1
		j   = -1
	)
	for line := 1; line < m.n; line++ {
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

func getPhi(a_ii, a_ij, a_jj float64) float64 {
	if a_ii == a_jj {
		return math.Pi / 4
	}
	return 0.5 * math.Atan(2*a_ij/(a_ii-a_jj))
}

func getSquareSum(m *Matrix) (sum float64) {
	for line := 1; line < m.n; line++ {
		lineIter := line * m.m
		for col := line + 1; col < m.m; col++ {
			elem := m.data[lineIter+col]
			sum += elem * elem
		}
	}
	return math.Sqrt(sum)
}

func Rotations(matrix *Matrix, eps float64) (sz Coloumn, sv *Matrix, iterations int, err error) {
	if matrix.n != matrix.m {
		err = IncorrectColoumn
		return
	}
	sv = EMatrix(matrix.n)
	sz = make(Coloumn, 0, matrix.n)
	matrix = matrix.Copy()

	for {
		iterations++
		var (
			U    = EMatrix(matrix.n)
			i, j = getMaxAbsElem(matrix)
		)
		if i == -1 || j == -1 {
			return
		}
		var (
			lineI = i * U.m
			lineJ = j * U.m

			phi = getPhi(
				matrix.data[lineI+i],
				matrix.data[lineI+j],
				matrix.data[lineJ+j])
			sin = math.Sin(phi)
			cos = math.Cos(phi)
		)
		U.data[lineI+j] = -sin
		U.data[lineJ+i] = sin
		U.data[lineI+i] = cos
		U.data[lineJ+j] = cos

		U_T := Transponse(U)
		matrix = U_T.ProdMatrix(matrix.ProdMatrix(U))
		sv = sv.ProdMatrix(U)

		if getSquareSum(matrix) < eps {
			break
		}
	}

	for line := 0; line < len(matrix.data); line += matrix.m + 1 {
		sz = append(sz, matrix.data[line])
	}

	return
}
