package core

import (
	"math/cmplx"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/mat"
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
			vals, _, err := QRValues(tt.m, tt.eps)
			require.NoError(t, err)

			var eig mat.Eigen
			ok := eig.Factorize(tt.m, mat.EigenLeft)
			require.True(t, ok)

			eigvals := eig.Values(nil)

			require.Len(t, vals, len(eigvals))

			sort.Sort(complexSlice(vals))
			sort.Sort(complexSlice(eigvals))
			for idx := range eigvals {
				require.InDelta(t, cmplx.Abs(eigvals[idx]), cmplx.Abs(vals[idx]), 0.01,
					"orig:\t%v\nmy:\t%v", eigvals, vals)
			}
		})
	}
}

type complexSlice []complex128

func (p complexSlice) Len() int { return len(p) }
func (p complexSlice) Less(i, j int) bool {
	return cmplx.Abs(p[i]) < cmplx.Abs(p[j]) || cmplx.IsNaN(p[i]) && !cmplx.IsNaN(p[j])
}
func (p complexSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
