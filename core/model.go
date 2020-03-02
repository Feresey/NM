package core

import "errors"

var (
	IncorrectColoumn = errors.New("Размерность столбца не совпадает")

	EPS = 1e-9
)

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
	L, U *Matrix
	// индексы перестановок
	P    []int
	n, m int
}
