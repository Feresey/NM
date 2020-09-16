package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/Arafatk/glot"
)

type points struct {
	x, y []float64
}

func makePoints(x []float64, f func(float64) float64) (res points) {
	res.x = make([]float64, len(x))
	res.y = make([]float64, len(x))
	for i, p := range x {
		res.x[i] = p
		res.y[i] = f(p)
	}
	return res
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

		fmt.Printf("i\tx_i\t\tf_i\t\tw'(x_i)\t\tf_i/w'(x_i)\tX*-x_i\n")
		for i := range w {
			fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n", i, points.x[i], points.y[i], w[i], points.y[i]/w[i], x-points.x[i])
			ww := points.y[i] * W / (x - points.x[i]) / w[i]
			res += ww
		}
		fmt.Println()
		return res
	}
}

func formatF(i, j int) string {
	if j == i {
		return fmt.Sprintf("x%d", i)
	}

	var b strings.Builder

	for idx := i; idx <= j; idx++ {
		b.WriteString(fmt.Sprintf("x_%d", idx))
		if idx != j {
			b.WriteByte(',')
		}
	}

	return b.String()
}

func Newthon(points points) func(float64) float64 {
	ff := make([]float64, len(points.x))
	for i := 1; i < len(ff); i++ {
		ff[i] = points.y[i]
	}

	for j := 0; j < len(ff); j++ {
		for i := 0; i < len(ff)-j-1; i++ {
			fmt.Printf("f(%s)=%f\tf(%s)=%f\t",
				formatF(i, i+j), ff[i],
				formatF(i+1, i+j+1), ff[i+1],
			)
			fmt.Printf("x%d=%f\tx%d=%f\n", i, points.x[i], i+j+1, points.x[j+1])
			ff[i] = (ff[i] - ff[i+1]) / (points.x[i] - points.x[i+j+1])
			fmt.Printf("f(%s)=%f\n",
				formatF(i, i+j+1), ff[i],
			)
			fmt.Println()
		}
	}

	return func(x float64) float64 {
		res := ff[0]
		for i := range ff {
			res = res*(x-points.x[i]) + ff[i]
		}
		return res
	}
}

func f(x float64) float64 {
	return math.Sin(x) + x
}

func solve(points points, x float64) {
	fmt.Printf("\npoints: %v, x: %f\n\n", points, x)

	interpolateLagrange := Lagrange(points)
	interpolateNewthon := Newthon(points)

	fmt.Printf("newthon: %3.15f\n", interpolateNewthon(x))
	fmt.Printf("lagrange: %3.15f\n", interpolateLagrange(x))
	fmt.Printf("real: %3.15f\n", f(x))

	plot, err := glot.NewPlot(2, true, false)
	if err != nil {
		panic(err)
	}

	more := make([]float64, len(points.x)*100)

	var (
		from float64 = points.x[0]
		to   float64 = points.x[len(points.x)-1]
	)
	step := (to - from) / float64(len(more))

	for idx := range more {
		more[idx] = from
		from += step
	}

	_ = plot.SetXrange(0, 2)
	_ = plot.SetYrange(0, 3)

	err = plot.AddFunc2d("original", "lines", more, f)
	if err != nil {
		panic(err)
	}
	err = plot.AddFunc2d("newthon", "lines", more, interpolateNewthon)
	if err != nil {
		panic(err)
	}
	err = plot.AddFunc2d("lagrange", "lines", more, interpolateLagrange)
	if err != nil {
		panic(err)
	}
}

func main() {
	X := 1.0
	A := []float64{0, math.Pi / 6, 2 * math.Pi / 6, 3 * math.Pi / 6}
	B := []float64{0, math.Pi / 6, math.Pi / 4, math.Pi / 2}

	apoints := makePoints(A, f)
	bpoints := makePoints(B, f)

	solve(apoints, X)
	solve(bpoints, X)
}
