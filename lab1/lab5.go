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
	a11 a12 ...
	a21 a22 ...
	... ... ...
	... ... ann
*/
func lab5(r io.Reader) {
	var (
		eps float64
		n   int
	)
	fscan(r, &eps)
	fscan(r, &n)
	matrix := readMatrix(r, n)
	fmt.Printf("A:\n%s\n", matrix)

	sz, total, err := core.QRValues(matrix, eps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Количество итераций :", total)
	fmt.Println("\nСобственные значения:")
	for idx := range sz {
		fmt.Printf("sz%d = %f\n", idx+1, sz[idx])
	}
}
