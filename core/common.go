package core

// NewMatrix : create a new matrix
func NewMatrix(N, M int) *Matrix {
	res := &Matrix{n: N, m: M}
	res.data = make(Row, N*M)

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

func Transponse(m *Matrix) *Matrix {
	res := NewMatrix(m.m, m.n)

	lineM := 0
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
