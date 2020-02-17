package main

import (
	"NM/core"
	"fmt"
	"io"
	"log"
	"os"
)

// readMatrix : читает матрицу.
/*
 * Формат:
 * <rows> <coloumns>
 * a11 a12 ...
 * a21 a22 ...
 * ...     ann
 */
func readMatrix(r io.Reader) (*core.Matrix, error) {
	var (
		n int
		m int
	)

	_, err := fmt.Fscanf(r, "%d %d", &n, &m)
	if err != nil {
		return nil, err
	}

	res := core.NewMatrix(n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num float64
			_, err := fmt.Fscan(r, &num)
			if err != nil {
				return nil, err
			}
			res.Set(i, j, num)
		}
	}

	return res, nil
}

func main() {
	fmt.Printf("Введите матрицу:\n")

	m, err := readMatrix(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	L, U, P, err := m.LUDecomposition()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("L:\n%s\nU:\n%s\nP:\n%s\n", L, U, P)
}
