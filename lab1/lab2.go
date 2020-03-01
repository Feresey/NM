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
func lab2(r io.Reader) {
	n := 0
	Fscan(r, &n)
	matrix, b := ReadSLAU(r, n)

	x := matrix.RunThrough(b)
	if x == nil {
		log.Fatal("Матрица пустая")
	}
	fmt.Printf("A:\n%s\n", core.DisplaySLAU{Matrix: matrix, Coloumn: b})

	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
