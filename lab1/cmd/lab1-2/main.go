package main

import (
	"fmt"
	"log"

	common "github.com/Feresey/NM/lab1"
)

/*
	Формат:
	n
	a11 a12 ... b1
	a21 a22 ... b2
	... ... ... ...
	... ... ann bn
*/
func main() {
	n := 0
	common.Scan(&n)
	matrix, b := common.ReadSLAU(n)

	x := matrix.RunThrough(b)
	if x == nil {
		log.Fatal("Матрица пустая")
	}

	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
