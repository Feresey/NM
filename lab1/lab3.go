package main

import (
	"fmt"
	"io"
	"log"

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
func lab3(r io.Reader) {
	var (
		eps float64
		n   int
	)
	fscan(r, &eps)
	fscan(r, &n)
	matrix, b := readSLAU(r, n)
	fmt.Printf("A:\n%s\n", core.DisplaySLAU{Matrix: matrix, Coloumn: b})

	x, total, err := core.Iterations(matrix, b, eps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total iterations: ", total)
	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
