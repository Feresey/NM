package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQR(t *testing.T) {
	tests := []struct {
		m    *Matrix
		eps  float64
		name string
	}{
		{
			name: "met",
			eps:  1e-9,
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
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			Q, R, err := QR(tt.m, tt.eps)
			require.NoError(t, err)

			prod := Q.ProdMatrix(R)

			require.InEpsilonSlice(t, tt.m.data, prod.data, tt.eps)
		})
	}
}
