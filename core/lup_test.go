package core

import (
	"math"
	"testing"
)

func TestMatrix_LUDecomposition(t *testing.T) {
	tests := []struct {
		name string
		Matrix
		wantErr bool
	}{
		{
			name: "hard",
			Matrix: Matrix{
				data: Row{
					3, 4, -9, 5,
					-15, -12, 50, -16,
					-27, -36, 73, 8,
					9, 12, -10, -16,
				},
				n: 4,
				m: 4,
			},
		},
		{
			name: "simple",
			Matrix: Matrix{
				data: Row{
					1, 3,
					2, 1,
				},
				n: 2,
				m: 2,
			},
		},
		{
			name: "simple solve",
			Matrix: Matrix{
				data: Row{
					1, 0, 1,
					0, 1, 2,
				},
				n: 2,
				m: 3,
			},
			wantErr: true,
		},
		{
			name: "zero coloumn",
			Matrix: Matrix{
				data: Row{
					1, 2, 3,
					1, 2, 4,
					1, 2, 5,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "fuck",
			Matrix: Matrix{
				data: Row{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "zero",
			Matrix: Matrix{
				data: Row{
					0, 0, 0,
					0, 0, 0,
					0, 0, 0,
				},
				n: 3,
				m: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lup := LUDecomposition(&tt.Matrix)
			if (lup == nil) != tt.wantErr {
				t.Error("Matrix is nil")
				return
			}
			if tt.wantErr {
				return
			}
			got := lup.L.ProdMatrix(lup.U).ProdMatrix(lup.P)

			if !matrixEqual(got, &tt.Matrix) {
				t.Errorf("L Given:\n%s\nWant:\n%s", got, tt.Matrix)
			}
		})
	}
}

func sumRow(matrix *Matrix, prod Row, row int) float64 {
	var (
		total float64
		line  = matrix.m * row
	)
	for i := 0; i < matrix.m; i++ {
		total += prod[i] * matrix.data[line+i]
	}

	return total
}

func TestSolveSLAU(t *testing.T) {
	type args struct {
		matrix *Matrix
		b      Row
	}
	tests := []struct {
		name               string
		args               args
		wantErr1, wantErr2 bool
	}{
		{
			name: "empty",
			args: args{
				matrix: NewMatrix(0, 0),
				b:      Row{},
			},
		},
		{
			name:     "nil",
			wantErr1: true,
		},
		{
			name: "zero",
			args: args{
				matrix: &Matrix{
					data: Row{
						0, 0, 0,
						0, 0, 0,
						0, 0, 0,
					},
					n: 3,
					m: 3,
				},
				b: Row{10, 0, -10},
			},
		},
		{
			name: "not square_1",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, 0,
					},
					n: 1,
					m: 2,
				},
			},
			wantErr1: true,
		},
		{
			name: "not square_2",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, 0,
						0, 1,
					},
					n: 2,
					m: 2,
				},
				b: Row{1, 2, 3},
			},
			wantErr2: true,
		},
		{
			name: "simple",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, 0,
						0, 1,
					},
					n: 2,
					m: 2,
				},
				b: Row{3, 1},
			},
		},
		{
			name: "simple rotated",
			args: args{
				matrix: &Matrix{
					data: Row{
						0, 1,
						1, 0,
					},
					n: 2,
					m: 2,
				},
				b: Row{3, 1},
			},
		},
		{
			name: "hard",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, 2,
						3, -1,
					},
					n: 2,
					m: 2,
				},
				b: Row{11, 12},
			},
		},
		{
			name: "harder",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, -2, 1,
						2, 2, -1,
						4, -1, 1,
					},
					n: 3,
					m: 3,
				},
				b: Row{0, 3, 5},
			},
		},
		{
			name: "big",
			args: args{
				matrix: &Matrix{
					data: Row{
						-1, -8, 0, 5,
						6, -6, 2, 4,
						9, -5, -6, 4,
						-5, 0, -9, 1,
					},
					n: 4,
					m: 4,
				},
				b: Row{-60, -10, 65, 18},
			},
		},
		{
			name: "also big",
			args: args{
				matrix: &Matrix{
					data: Row{
						1, 2, -2, 6,
						-3, -5, 14, 13,
						1, 2, -2, -2,
						-2, -4, 5, 10,
					},
					n: 4,
					m: 4,
				},
				b: Row{24, 41, 0, 20},
			},
		},
		{
			name: "fuck",
			args: args{
				matrix: &Matrix{
					data: Row{
						-1, -3, 0, -4,
						3, 7, -8, 3,
						1, -6, 2, 5,
						-8, -4, -1, -1,
					},
					n: 4,
					m: 4,
				},
				b: Row{-3, 30, -90, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lup := LUDecomposition(tt.args.matrix)
			if (lup == nil) != tt.wantErr1 {
				t.Errorf("LUDecomposition() wantErr %v", tt.wantErr1)
				return
			}
			if tt.wantErr1 {
				return
			}
			got := lup.SolveSLAU(tt.args.b)
			if (got == nil) != tt.wantErr2 {
				t.Errorf("SolveSLAU() wantErr %v", tt.wantErr2)
				return
			}
			if tt.wantErr2 {
				return
			}
			for i := 0; i < lup.n; i++ {
				if tmp := sumRow(tt.args.matrix, got, i); math.Abs(tmp-tt.args.b[i]) > EPS {
					t.Errorf("incorrect answer. got: %v, but sum of %d is %f, expected %f", got, i, tmp, -tt.args.b[i])
				}
			}
		})
	}
}
