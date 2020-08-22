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
	fscan(r, &n)
	matrix, b := readSLAU(r, n)

	x, err := core.RunThrough(matrix, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("A:\n%s\n", core.DisplaySLAU(matrix, b))

	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
