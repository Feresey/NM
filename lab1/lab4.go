package main

import (
	"fmt"
	"io"

	"github.com/Feresey/NM/core"
)

/*
	Формат:
	eps n
	a11 a12 ... b1
	a21 a22 ... b2
	... ... ... ...
	... ... ann bn
*/
func lab4(r io.Reader) {
	var (
		eps float64
		n   int
	)
	fscan(r, &eps)
	fscan(r, &n)
	matrix := readMatrix(r, n)
	fmt.Printf("A:\n%s\n", matrix)

	sz, sv, total, err := core.Rotations(matrix, eps)
	_ = sz
	_ = sv
	_ = total
	_ = err
	// if x == nil {
	// 	log.Fatal("Матрица пустая")
	// }

	// fmt.Println("Total iterations: ", total)
	// for idx := range x {
	// 	fmt.Printf("x%d = %f\n", idx+1, x[idx])
	// }
}
