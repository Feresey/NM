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
type DisplaySLAU struct {
	*Matrix
	Coloumn
}

type LUP struct {
	L, U, P *Matrix

	n, m int
}
