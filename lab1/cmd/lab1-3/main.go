package main

import (
	"fmt"
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
func main() {
	var (
		eps float64
		n   int
	)
	common.Scan(&eps)
	common.Scan(&n)
	matrix, b := common.ReadSLAU(n)

	x, total := matrix.Iterations(b, eps)
	if x == nil {
		log.Fatal("Матрица пустая")
	}
	fmt.Printf("A:\n%s\n", core.DisplaySLAU{Matrix: matrix, Coloumn: b})

	fmt.Println("Total iterations: ", total)
	for idx := range x {
		fmt.Printf("x%d = %f\n", idx+1, x[idx])
	}
}
