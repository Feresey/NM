package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQR(t *testing.T) {
	eps := 1e-9

	tests := []struct {
		m    *Matrix
		name string
	}{
		{
			name: "met",
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
			Q, R, err := QR(tt.m, eps)
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
	}
}
