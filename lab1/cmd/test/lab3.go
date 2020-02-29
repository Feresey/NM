package main

import (
	"fmt"
	"io"
	"log"

	"github.com/Feresey/NM/core"
	common "github.com/Feresey/NM/lab1"
)

/*
	Формат:
	eps n
	a11 a12 ... b1
	a21 a22 ... b2
	... ... ... ...
	... ... ann bn
*/
func lab3(r io.Reader) {
	var (
		eps float64
		n   int
	)
	common.Fscan(r, &eps)
	common.Fscan(r, &n)
	matrix, b := common.ReadSLAU(r, n)
	fmt.Printf("A:\n%s\n", core.DisplaySLAU{Matrix: matrix, Coloumn: b})

	x, total := matrix.Iterations(b, eps)
	if x == nil {
		log.Fatal("Матрица пустая")
	}

	fmt.Println("Total iterations: ", total)
	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
