package main

import "fmt"

type points struct {
	x, y []float64
}

func Lagrange(points points) func(float64) float64 {
	w := make([]float64, len(points.x))
	for idx := range points.x {
		var res float64 = 1
		for curr, point := range points.x {
			if curr != idx {
				res *= points.x[idx] - point
			}
		}
		w[idx] = res
	}

	return func(x float64) float64 {
		var res float64
		var W float64 = 1
		for i := range points.x {
			W *= x - points.x[i]
		}

		// fmt.Printf("i\tx_i\t\tf_i\t\tw'(x_i)\t\tf_i/w'(x_i)\tX*-x_i\n")
		for i := range w {
			// fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n", i, points.x[i], points.y[i], w[i], points.y[i]/w[i], x-points.x[i])
			ww := points.y[i] * W / (x - points.x[i]) / w[i]
			res += ww
		}
		return res
	}
}

// n
// x1, x2, ..., xn
// y1, y2, ..., yn
// x*
func main() {
	var (
		n      int
		points points
		x      float64
	)
	fmt.Scanln(&n)
	points.x = make([]float64, n)
	points.y = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points.x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&points.y[i])
	}
	fmt.Scanln(&x)

	f := Lagrange(points)

	fmt.Printf("res: %3.15f\n", f(x))
}
