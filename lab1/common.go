package common

import (
	"fmt"
	"io"
	"log"

	"github.com/Feresey/NM/core"
)

func Fscan(r io.Reader, ref interface{}) {
	_, err := fmt.Fscan(r, ref)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadSLAU(r io.Reader, n int) (*core.Matrix, core.Coloumn) {
	matrix := core.NewMatrix(n, n)
	col := make(core.Coloumn, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var num float64
			_, err := fmt.Fscan(r, &num)
			if err != nil {
				log.Fatal("Error reading matrix: ", err)
			}
			matrix.Set(i, j, num)
		}
		_, err := fmt.Fscan(r, &col[i])
		if err != nil {
			log.Fatal("Error reading params: ", err)
		}
	}

	return matrix, col
}
