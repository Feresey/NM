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
func lab4(r io.Reader) {
	var (
		eps float64
		n   int
	)
	fscan(r, &eps)
	fscan(r, &n)
	matrix := readMatrix(r, n)
	fmt.Printf("A:\n%s\n", matrix)

	sz, sv, total, err := core.Rotations(matrix, eps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Количество итераций :", total)
	fmt.Println("\nСобственные значения:")
	for idx := range sz {
		fmt.Printf("sz%d = %f\n", idx+1, sz[idx])
	}

	fmt.Println("\nСобственные векторы:")
	vectors := core.Transponse(sv)

	for i := 0; i < n; i++ {
		fmt.Printf("x%d = (", i)
		for j := 0; j < n; j++ {
			fmt.Printf("%7.2f", vectors.At(i, j))
		}
		fmt.Println(")")
	}
}
