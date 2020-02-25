package core

type (
	Row     []float64
	Coloumn []float64
)

// Matrix : type for process core actions
type Matrix struct {
	data Row
	// Number lines
	n int
	// Number rows
	m int
}

// NewMatrix : create a new matrix
func NewMatrix(N, M int) *Matrix {
	res := &Matrix{n: N, m: M}
	res.data = make(Row, N*M)

	return res
}

type DisplaySLAU struct {
	*Matrix
	Coloumn
}

type LUP struct {
	L, U, P *Matrix

	n, m int
}
