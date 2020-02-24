package main

import (
	"NM/core"
	"fmt"
	"log"
)

func scan(ref interface{}) {
	_, err := fmt.Scan(ref)
	if err != nil {
		log.Fatal(err)
	}
}

func readSLAU(n int) (*core.Matrix, core.Row) {
	matrix := core.NewMatrix(n, n)
	col := make(core.Row, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var num float64
			scan(&num)
			matrix.Set(i, j, num)
		}
		scan(&col[i])
	}

	return matrix, col
}

/*
	Формат:
	<rows> <coloumns>

	a11 a12 ... b1
	a21 a22 ... b2
	... ... ... ...
	... ... ann bn
*/
func main() {
	fmt.Print("Введите количество строк: ")
	n := 0
	scan(&n)
	fmt.Println("Введите элементы матрицы:")
	matrix, b := readSLAU(n)

	lup, err := core.LUDecomposition(matrix)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\nDet(A) = %f\n", core.DisplaySLAU{Matrix: matrix, Row: b}, lup.Determinant())

	x, err := lup.SolveSLAU(b)
	if err != nil {
		log.Fatal(err)
	}

	for idx := range x {
		fmt.Printf("x[%d] = %f\n", idx, x[idx])
	}
}
