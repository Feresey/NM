package main

import (
	"flag"
	"fmt"
	"math"
)

// f(x) = 3^x-5*x^2+1
func f(x float64) float64 {
	return math.Pow(3, x) - 5*x*x + 1
}

// f'(x) = 3^x/10 * ln(3)/sqrt(3^x-1)
func f_derivatived(x float64) float64 {
	pow := math.Pow(3, x)
	return pow / 10 * math.Log(3) / math.Sqrt(pow-1)
}

func fIter(lambda func(float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return x - lambda(x)*f(x)
	}
}

func iterations(
	f func(float64) float64,
	from, to float64,
	q float64,
	eps float64,
) (res float64, iterations int) {
	fmt.Printf("Iterations\nfrom: %f\nto: %f\n", from, to)
	var prev, curr float64

	prev = (from + to) / 2
	fmt.Println(prev)

	for ; ; iterations++ {
		curr = f(prev)
		fmt.Printf("Iter: %d, value: %3.12f\n", iterations, curr)

		if q/(1-q)*math.Abs(prev-curr) < eps {
			return curr, iterations
		}
		prev = curr
	}
}

func main() {
	from := flag.Float64("from", 0.7, "left point")
	to := flag.Float64("to", 0.9, "right point")
	q := flag.Float64("q", 0.25, "magick number, |phi'(x)| <= q < 1, Any x in (a,b)")
	lambda := flag.Float64("lambda", -0.1, "second magick number")
	eps := flag.Float64("eps", 1e-9, "calculation precision")
	flag.Parse()

	res, iters := iterations(fIter(func(float64) float64 { return *lambda }), *from, *to, *q, *eps)
	fmt.Printf("Iterations method:\niter: %d\nresult: %3.12f\n", iters, res)
}
