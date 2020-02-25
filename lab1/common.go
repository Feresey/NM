package common

import (
	"fmt"
	"log"

	"github.com/Feresey/NM/core"
)

func Scan(ref interface{}) {
	_, err := fmt.Scan(ref)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadSLAU(n int) (*core.Matrix, core.Row) {
	matrix := core.NewMatrix(n, n)
	col := make(core.Row, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var num float64
			Scan(&num)
			matrix.Set(i, j, num)
		}
		Scan(&col[i])
	}

	return matrix, col
}
