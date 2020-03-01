package main

import (
	"fmt"
	"io"
	"log"

	"github.com/Feresey/NM/core"
)

/*
	Формат:
	n
	a11 a12 ... b1
	a21 a22 ... b2
	... ... ... ...
	... ... ann bn
*/
func lab1(r io.Reader) {
	n := 0
	fscan(r, &n)
	matrix, b := readSLAU(r, n)

	lup := core.LUDecomposition(matrix)
	if lup == nil {
		log.Fatal("Матрица пустая")
	}

	fmt.Printf("A:\n%s\n", core.DisplaySLAU{Matrix: matrix, Coloumn: b})
	// fmt.Printf("L*U*P:\n%s\n", lup.L.ProdMatrix(lup.U).ProdMatrix(lup.P))
	// fmt.Printf("L:\n%s\nU:\n%s\nP:\n%s\n", lup.L, lup.U, lup.P)
	fmt.Printf("Det(A) = %f\n", lup.Determinant())
	inv := lup.Inverse()
	if lup == nil {
		log.Fatal("Матрица пустая")
	}

	fmt.Printf("Inverse(A):\n%s\n", inv)
	// fmt.Printf("A*Inverse(A):\n%s\n", matrix.ProdMatrix(inv))

	x := lup.SolveSLAU(b)
	if x == nil {
		log.Fatal("Матрица пустая")
	}

	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
