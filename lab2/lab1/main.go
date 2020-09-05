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

// f'(x) = 3^x * ln(3) - 10*x
func derived(x float64) float64 {
	return math.Pow(3, x)*math.Log(3) - 10*x
}

func fIter(lambda func(float64) float64) func(x float64) float64 {
	return func(x float64) float64 {
		return x - lambda(x)*f(x)
	}
}

func iterations(
	f func(float64) float64,
	finish func(iter int, prev, curr float64) bool,
	from, to float64,
) (res float64, iterations int) {
	var prev, curr float64

	iterations++
	prev = (from + to) / 2
	fmt.Printf("First x_0: %f\n", prev)

	for ; ; iterations++ {
		curr = f(prev)
		// if math.IsNaN(curr) {
		// 	panic(prev)
		// }

		fmt.Printf("Iter: %d, value: %3.15f\n", iterations, curr)
		if finish(iterations, prev, curr) {
			return curr, iterations
		}
		prev = curr
	}
}

func main() {
	from := flag.Float64("from", 0.7, "left point")
	to := flag.Float64("to", 0.9, "right point")
	q := flag.Float64("q", 0.4, "magick number, |phi'(x)| <= q < 1, Any x in (a,b)")
	lambda := flag.Float64("lambda", -0.1, "second magick number")
	eps := flag.Float64("eps", 1e-9, "calculation precision")
	flag.Parse()

	fmt.Printf("Iterations\nfrom: %f\nto: %f\n", *from, *to)
	iterRes, iters := iterations(
		fIter(func(float64) float64 { return *lambda }),
		func(iter int, prev, curr float64) bool {
			return math.Pow(*q, float64(iter))/(1-*q)*math.Abs(prev-curr) < *eps
		},
		*from, *to,
	)
	fmt.Printf("\nIterations method:\niterations: %d\nresult: %3.15f\n\n", iters, iterRes)

	fmt.Printf("Newthon\nfrom: %f\nto: %f\n", *from, *to)
	newthonRes, iters := iterations(
		fIter(func(x float64) float64 { return 1 / derived(x) }),
		func(iter int, prev, curr float64) bool {
			return math.Abs(prev-curr) < *eps
		},
		*from, *to,
	)
	fmt.Printf("\nNewthon method:\niterations: %d\nresult: %3.15f\n\n", iters, newthonRes)

	fmt.Printf("Iterations: %3.15f\n", f(iterRes))
	fmt.Printf("Newthon: %3.15f\n", f(newthonRes))
	fmt.Printf("diff: %3.15f\n", f(iterRes)-f(newthonRes))
}
