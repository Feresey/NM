package matrix

// Matrix : type for process matrix actions
type Matrix struct {
	data []float64
	// Number lines
	n int
	// Number rows
	m int
}

// NewMatrix : create a new matrix
func NewMatrix(N, M int) *Matrix {
	res := &Matrix{n: N, m: M}
	res.data = make([]float64, N*M)

	return res
}
