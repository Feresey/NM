package main

import (
	"NM/core"
	"fmt"
	"log"
)

// readMatrix : читает матрицу.
/*
 * Формат:
 * <rows> <coloumns>
 * a11 a12 ...
 * a21 a22 ...
 * ...     ann
 */
func readMatrix() *core.Matrix {
	var (
		n int
		m int
	)

	fmt.Print("Введите матрицу:\nN, M: ")

	_, err := fmt.Scanf("%d %d", &n, &m)
	if err != nil {
		log.Fatal(err)
	}

	res := core.NewMatrix(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num float64
			_, err := fmt.Scan(&num)
			if err != nil {
				log.Fatal(err)

			}
			res.Set(i, j, num)
		}
	}

	return res
}

func main() {
	m := readMatrix()

	L, U, P, err := m.LUDecomposition()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("L:\n%s\nU:\n%s\nP:\n%s\n", L, U, P)

	LU, err := L.ProdMatrix(U)
	if err != nil {
		log.Fatal(err)
	}
	LUP, err := LU.ProdMatrix(P)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("L*U*P:\n%s\n", LUP)
}
