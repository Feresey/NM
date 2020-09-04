package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQR(t *testing.T) {
	eps := 1e-9

	tests := []struct {
		m    *Matrix
		eps  float64
		name string
	}{
		{
			name: "met",
			eps:  0.1,
			m: &Matrix{
				data: []float64{
					1, 3, 1,
					1, 1, 4,
					4, 3, 1,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "my",
			eps:  1e-9,
			m: &Matrix{
				data: []float64{
					5, -1, -2,
					-4, 3, -3,
					-2, -1, 1,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "not my",
			eps:  1e-9,
			m: &Matrix{
				data: []float64{
					-7, 6, 0,
					0, 7, 3,
					1, 5, -4,
				},
				n: 3,
				m: 3,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			Q, R, err := QR(tt.m)
			require.NoError(t, err)

			prod := Q.ProdMatrix(R)

			tr := Transponse(Q)
			inv, err := Q.Inverse()
			require.NoError(t, err)

			sub := tr.Copy()
			sub.Sub(inv)
			require.InDeltaSlice(t, tr.data, inv.data, eps)

			require.InDeltaSlice(t, tt.m.data, prod.data, eps)
		})
		t.Run(tt.name+"vals", func(t *testing.T) {
			vals, iters, err := QRValues(tt.m, tt.eps)
			require.NoError(t, err)
			require.NotEmpty(t, vals)
			require.NotZero(t, iters)
		})
	}
}
