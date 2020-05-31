package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestMatrix_Iterations(t *testing.T) {
	type args struct {
		col Coloumn
		eps float64
	}

	tests := []struct {
		name string
		*Matrix
		args           args
		wantRes        Coloumn
		wantIterations int
		wantErr        bool
	}{
		{
			name: "simple",
			args: args{
				col: Coloumn{12, 13, 14},
				eps: 0.01,
			},
			Matrix: &Matrix{
				data: Row{
					10, 1, 1,
					2, 10, 1,
					2, 2, 10,
				},
				m: 3,
				n: 3,
			},
			wantIterations: 5,
			wantRes:        Coloumn{1, 1, 1},
		},
		{
			name: "big",
			Matrix: &Matrix{
				data: Row{
					14, -4, -2, 3,
					-3, 23, -6, -9,
					-7, -8, 21, -5,
					-2, -2, 8, 18,
				},
				m: 4,
				n: 4,
			},
			args: args{
				col: Coloumn{38, -195, -27, 142},
				eps: 1e-9,
			},
			wantRes:        Coloumn{-1, -6, -2, 8},
			wantIterations: 19,
		},
		{
			name: "big",
			Matrix: &Matrix{
				data: Row{
					24, 2, 4, -9,
					-6, -27, -8, -6,
					-4, 8, 19, 6,
					4, 5, -3, -13,
				},
				m: 4,
				n: 4,
			},
			args: args{
				col: Coloumn{-9, -76, -79, -70},
				eps: 1e-9,
			},
			wantRes:        Coloumn{4, 2, -7, 9},
			wantIterations: 30,
		},
		{
			name: "abort",
			args: args{
				col: Coloumn{1, 1, 1},
			},
			Matrix: &Matrix{
				data: Row{
					1, -2, 1,
					0, 1, 0,
					1, 0, 1,
				},
				m: 3,
				n: 3,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(ttt *testing.T) {
			t := assert.New(ttt)

			gotRes, gotIterations, err := Iterations(
				mat.NewDense(tt.n, tt.m, tt.Matrix.data),
				mat.NewVecDense(tt.n, tt.args.col),
				tt.args.eps,
			)
			if !t.Equal(tt.wantErr, err != nil, "Err: %v, test: %s", err, tt.name) {
				return
			}
			if err != nil {
				return
			}

			t.Equal(tt.wantIterations, gotIterations, gotRes)

			if !t.InEpsilonSlice(
				tt.wantRes,
				gotRes.RawVector().Data,
				tt.args.eps,
				gotRes,
			) {
				return
			}
		})
	}
}

// func TestMatrix_norm(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		Matrix
// 		wantRes float64
// 	}{
// 		{
// 			name: "simple",
// 			Matrix: Matrix{
// 				data: Row{
// 					1, 0, 0,
// 					0, 1, 0,
// 					0, 0, -1,
// 				},
// 				m: 3,
// 				n: 3,
// 			},
// 			wantRes: 1,
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt

// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotRes := tt.Matrix.norm(); gotRes != tt.wantRes {
// 				t.Errorf("Matrix.norm() = %v, want %v", gotRes, tt.wantRes)
// 			}
// 		})
// 	}
// }

// func Test_norm(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		data    []float64
// 		wantRes float64
// 	}{
// 		{
// 			name:    "simple",
// 			data:    []float64{1, 2, -1, 3, 4, 0},
// 			wantRes: 4,
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt

// 		t.Run(tt.name, func(t *testing.T) {
// 			if gotRes := norm(tt.data); gotRes != tt.wantRes {
// 				t.Errorf("norm() = %v, want %v", gotRes, tt.wantRes)
// 			}
// 		})
// 	}
// }
