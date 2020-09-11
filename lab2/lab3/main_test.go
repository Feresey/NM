package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func makePoints(x []float64, f func(float64) float64) (res points) {
	res.x = make([]float64, len(x))
	res.y = make([]float64, len(x))
	for i, p := range x {
		res.x[i] = p
		res.y[i] = f(p)
	}
	return res
}

func TestLagrange(t *testing.T) {
	tests := []struct {
		name   string
		points points
		x      float64
		want   float64
		delta  float64
	}{
		{
			name:   "ln",
			points: makePoints([]float64{0.1, 0.5, 0.9, 1.3}, math.Log),
			x:      0.8,
			want:   math.Log(0.8),
			delta:  0.02278,
		},
		{
			name: "sin",
			points: makePoints([]float64{0, 1, 2, 3}, func(x float64) float64 {
				return math.Sin(math.Pi / 6 * x)
			}),
			x:     1.5,
			want:  math.Sin(math.Pi / 6 * 1.5),
			delta: 0.0165,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			f := Lagrange(tt.points)

			require.InDelta(t, tt.want, f(tt.x), tt.delta)
		})
	}
}
