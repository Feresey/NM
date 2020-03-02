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
			wantErr: true,
		},
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
			wantErr: true,
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lup, err := LUDecomposition(&tt.Matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("error: %v, wantErr: %t", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			got := lup.SwapMatrix(lup.L.ProdMatrix(lup.U), false)

			if !matrixEqual(got, &tt.Matrix, EPS) {
				t.Errorf("Given:\n%s\nWant:\n%s", got, tt.Matrix)
				// t.Error(lup.P)
			}
		})
	}
}

func sumRow(matrix *Matrix, prod Coloumn, row int) float64 {
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
		b      Coloumn
	}
	tests := []struct {
		name               string
		args               args
		wantErr1, wantErr2 bool
		det                float64
	}{
		{
			name: "empty",
			args: args{
				matrix: NewMatrix(0, 0),
				b:      Coloumn{},
			},
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
				b: Coloumn{10, 0, -10},
			},
			wantErr1: true,
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
				b: Coloumn{1, 2, 3},
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
				b: Coloumn{3, 1},
			},
			det: 1,
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
				b: Coloumn{3, 1},
			},
			det: -1,
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
				b: Coloumn{11, 12},
			},
			det: -7,
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
				b: Coloumn{0, 3, 5},
			},
			det: 3,
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
				b: Coloumn{-60, -10, 65, 18},
			},
			det: 356,
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
				b: Coloumn{24, 41, 0, 20},
			},
			det: -8,
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
				b: Coloumn{-3, 30, -90, 12},
			},
			det: 2239,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lup, err := LUDecomposition(tt.args.matrix)
			if (err != nil) != tt.wantErr1 {
				t.Error("Unexpected error: ", err)
			}
			if tt.wantErr1 {
				return
			}

			// gat := lup.L.ProdMatrix(lup.U)
			// tmp := lup.SwapMatrix(&Matrix{data: Row(tt.args.b), n: tt.args.matrix.n, m: 1})
			// fmt.Println("A:", gat, "b:", tmp)
			// fmt.Println("A:", tt.args.matrix, "b:", tt.args.b)

			got, err := lup.SolveSLAU(tt.args.b)
			if (err != nil) != tt.wantErr2 {
				t.Errorf("SolveSLAU() wantErr %v", tt.wantErr2)
				return
			}
			if tt.wantErr2 {
				return
			}
			for i := 0; i < lup.n; i++ {
				if tmp := sumRow(tt.args.matrix, got, i); math.Abs(tmp-tt.args.b[i]) > EPS {
					t.Errorf("incorrect answer. got: %v, but sum of %d is %f, expected %f", got, i, tmp, -tt.args.b[i])

					// if !matrixEqual(gat, tt.args.matrix, EPS) {
					// 	t.Errorf("Given:\n%s\nWant:\n%s", gat, tt.args.matrix)
					// 	t.Error(lup.P)
					// }
					// return
				}
			}
			inverse := lup.Inverse()
			if !matrixEqual(tt.args.matrix.ProdMatrix(inverse), EMatrix(tt.args.matrix.n), EPS) {
				t.Error("Inverse matrix does not correct")
			}
			if tmp := lup.Determinant(); math.Abs(tt.det-tmp) > EPS {
				t.Errorf("Incorrect det. got: %f, expected: %f", tmp, tt.det)
			}
		})
	}

	_ = DisplaySLAU{Matrix: EMatrix(3), Coloumn: Coloumn{1, 1, 1}}.String()
}

func TestLUP_SwapMatrix(t *testing.T) {
	type fields struct {
		P []int
	}
	type args struct {
		m *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		{
			name: "no swap",
			args: args{
				m: EMatrix(2),
			},
			fields: fields{
				P: []int{0, 1},
			},
			want: EMatrix(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lup := &LUP{
				P: tt.fields.P,
			}
			got := lup.SwapMatrix(tt.args.m, false)
			if !matrixEqual(got, tt.args.m, EPS) {
				t.Errorf("Given:\n%s\nWant:\n%s", got, tt.want)
				t.Error(lup.P)
			}
		})
	}
}
