package main

import (
	"flag"
	"fmt"
	"math"
)

func f1(x1, x2 float64) float64 {
	return x1 - math.Cos(x2) - 3
}
func df1X1(x1, x2 float64) float64 {
	return 1
}
func df1X2(x1, x2 float64) float64 {
	return math.Sin(x2)
}

func f2(x1, x2 float64) float64 {
	return x2 - math.Sin(x1) - 3
}
func df2X2(x1, x2 float64) float64 {
	return 1
}
func df2X1(x1, x2 float64) float64 {
	return -math.Cos(x1)
}

func detA1(x1, x2 float64) float64 {
	return f1(x1, x2)*df2X2(x1, x2) - f2(x1, x2)*df1X2(x1, x2)
}

func detA2(x1, x2 float64) float64 {
	return df1X1(x1, x2)*f2(x1, x2) - df2X1(x1, x2)*f1(x1, x2)
}

func detJ(x1, x2 float64) float64 {
	return df1X1(x1, x2)*df2X2(x1, x2) - df2X1(x1, x2)*df1X2(x1, x2)
}

func phi1(x1, x2 float64) float64 {
	return 3 + math.Cos(x2)
}

func phi2(x1, x2 float64) float64 {
	return 3 + math.Sin(x2)
}

func newthon(x1, x2 float64, eps float64) (res1 float64, res2 float64, iterations int) {
	var (
		prevX1 = x1
		prevX2 = x2
	)

	iterations++
	for ; ; iterations++ {
		j := detJ(prevX1, prevX2)
		x1 -= detA1(prevX1, prevX2) / j
		x2 -= detA2(prevX1, prevX2) / j

		diff1 := math.Abs(prevX1 - x1)
		diff2 := math.Abs(prevX2 - x2)

		fmt.Printf("iteration: %d\n", iterations)
		fmt.Printf("x1: %3.15f, diff: %3.15f\n", x1, diff1)
		fmt.Printf("x2: %3.15f, diff: %3.15f\n", x2, diff2)

		if diff1 < eps && diff2 < eps {
			return x1, x2, iterations
		}
		prevX1, prevX2 = x1, x2
	}
}

func iterations(x1, x2 float64, q float64, eps float64) (res1 float64, res2 float64, iterations int) {
	var (
		prevX1 = x1
		prevX2 = x2
	)

	iterations++
	for ; ; iterations++ {
		x1 = phi1(x1, x2)
		x2 = phi2(x2, x1)

		diff1 := math.Abs(prevX1 - x1)
		diff2 := math.Abs(prevX2 - x2)

		fmt.Printf("iteration: %d\n", iterations)
		fmt.Printf("x1: %3.15f, diff: %3.15f\n", x1, diff1)
		fmt.Printf("x2: %3.15f, diff: %3.15f\n", x2, diff2)

		max := diff1
		if diff2 > max {
			max = diff2
		}

		if math.Pow(q, float64(iterations))/(1-q)*max < eps {
			return x1, x2, iterations
		}

		prevX1, prevX2 = x1, x2
	}
}

func main() {
	x1 := flag.Float64("x1", 4, "first point")
	x2 := flag.Float64("x2", 4, "second point")
	eps := flag.Float64("eps", 1e-9, "calculation accuracy")
	q := flag.Float64("q", 0.4, "magick number, |phi'(x)| <= q < 1, Any x in (a,b)")
	flag.Parse()

	x1Iter, x2Iter, totalIter := iterations(*x1, *x2, *q, *eps)
	fmt.Printf("\nIterations\niterations: %d\nx1=%3.15f,x2=%3.15f\n", totalIter, x1Iter, x2Iter)

	x1Newthon, x2Newthon, totalNewthon := newthon(*x1, *x2, *eps)
	fmt.Printf("\nNewthon\niterations: %d\nx1=%3.15f,x2=%3.15f\n", totalNewthon, x1Newthon, x2Newthon)
}
