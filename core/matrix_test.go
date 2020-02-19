package core

import (
	"testing"
)

const EPS = 1e-9

func matrixEqual(a, b *Matrix) bool {
	return floatEqual(a.data, b.data)
}

func floatEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if tmp := a[i] - b[i]; tmp > 0 && tmp > EPS || tmp < 0 && tmp < -EPS {
			return false
		}
	}
	return true
}

func TestMatrix_ProdMatrix(t *testing.T) {
	tests := []struct {
		name string
		Matrix
		arg     Matrix
		want    Matrix
		wantErr bool
	}{
		{
			name: "E",
			Matrix: Matrix{
				data: []float64{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				},
				n: 3,
				m: 3,
			},
			arg: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
			want: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "2x3*3x2",
			Matrix: Matrix{
				data: []float64{
					1, 0, 1,
					0, 1, 0,
				},
				n: 2,
				m: 3,
			},
			arg: Matrix{
				data: []float64{
					1, 2,
					3, 4,
					5, 6,
				},
				n: 3,
				m: 2,
			},
			want: Matrix{
				data: []float64{
					6, 8,
					3, 4,
				},
				n: 2,
				m: 2,
			},
		},
		{
			name: "3x3*3x1",
			Matrix: Matrix{
				data: []float64{
					3, -1, 2,
					4, 2, 0,
					-5, 6, 1,
				},
				n: 3,
				m: 3,
			},
			arg: Matrix{
				data: []float64{
					8,
					7,
					2,
				},
				n: 3,
				m: 1,
			},
			want: Matrix{
				data: []float64{
					21,
					46,
					4,
				},
				n: 3,
				m: 1,
			},
		},
		{
			name: "error",
			Matrix: Matrix{
				data: []float64{
					3, -1, 2,
					4, 2, 0,
					-5, 6, 1,
				},
				n: 3,
				m: 3,
			},
			arg: Matrix{
				data: []float64{
					2,
				},
				n: 1,
				m: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.Matrix.ProdMatrix(&tt.arg)
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr && !matrixEqual(got, &tt.want) {
				t.Errorf("Given:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}

	if tests[0].arg.String() == "" {
		t.Error("oops")
	}
	if n, m := tests[0].arg.GetSize(); n == 0 || m == 0 {
		t.Error("oops")
	}
}

func TestMatrix_SwapLines(t *testing.T) {
	type arg struct {
		a int
		b int
	}
	tests := []struct {
		name string
		Matrix
		arg
		wantMatrix Matrix
	}{
		{
			name: "3",
			arg: arg{
				a: 0,
				b: 1,
			},
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
			wantMatrix: Matrix{
				data: []float64{
					4, 5, 6,
					1, 2, 3,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "2",
			arg: arg{
				a: 0,
				b: 2,
			},
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
			wantMatrix: Matrix{
				data: []float64{
					7, 8, 9,
					4, 5, 6,
					1, 2, 3,
				},
				n: 3,
				m: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Matrix.SwapLines(tt.arg.a, tt.arg.b)
			if !matrixEqual(&tt.Matrix, &tt.wantMatrix) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.wantMatrix)
			}
		})
	}
}

func TestMatrix_LUDecomposition(t *testing.T) {
	tests := []struct {
		name string
		Matrix
		L       *Matrix
		U       *Matrix
		P       *Matrix
		wantErr bool
	}{
		{
			name: "hard",
			Matrix: Matrix{
				data: []float64{
					3, 4, -9, 5,
					-15, -12, 50, -16,
					-27, -36, 73, 8,
					9, 12, -10, -16,
				},
				n: 4,
				m: 4,
			},
			L: &Matrix{
				data: []float64{
					1, 0, 0, 0,
					-5, 1, 0, 0,
					-9, 0, 1, 0,
					3, 0, -2.125, 1,
				},
				n: 4,
				m: 4,
			},
			U: &Matrix{
				data: []float64{
					3, 4, -9, 5,
					0, 8, 5, 9,
					0, 0, -8, 53,
					0, 0, 0, 81.625,
				},
				n: 4,
				m: 4,
			},
			P: EMatrix(4),
		},
		{
			name: "simple",
			Matrix: Matrix{
				data: []float64{
					1, 3,
					2, 1,
				},
				n: 2,
				m: 2,
			},
			L: &Matrix{
				data: []float64{
					1, 0,
					2, 1,
				},
				n: 2,
				m: 2,
			},
			U: &Matrix{
				data: []float64{
					1, 3,
					0, -5,
				},
				n: 2,
				m: 2,
			},
			P: EMatrix(2),
		},
		{
			name: "simple solve",
			Matrix: Matrix{
				data: []float64{
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
				data: []float64{
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
			name: "fuck",
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L, U, P, err := tt.Matrix.LUDecomposition()
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}
			if tt.wantErr {
				return
			}
			if !matrixEqual(L, tt.L) {
				t.Errorf("L Given:\n%s\nWant:\n%s", L, tt.L)
			}
			if !matrixEqual(U, tt.U) {
				t.Errorf("U Given:\n%s\nWant:\n%s", U, tt.U)
			}
			if !matrixEqual(P, tt.P) {
				t.Errorf("P Given:\n%s\nWant:\n%s", P, tt.P)
			}
		})
	}
}

func TestMatrix_findNotZeroIndexInCol(t *testing.T) {
	type args struct {
		idx int
		P   *Matrix
	}
	tests := []struct {
		name string
		Matrix
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "simple",
			Matrix: Matrix{
				data: []float64{
					0, 1,
					1, 0,
				},
				n: 2,
				m: 2,
			},
			args: args{
				idx: 0,
			},
			want: Matrix{
				data: []float64{
					1, 0,
					0, 1,
				},
				n: 2,
				m: 2,
			},
		},
		{
			name: "hard",
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					0, 0, 4,
					0, 0, 5,
				},
				n: 3,
				m: 3,
			},
			args: args{
				idx: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.Matrix.findNotZeroIndexInCol(tt.args.idx, tt.args.P); (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}
			if !tt.wantErr && !matrixEqual(&tt.Matrix, &tt.want) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.want)
			}
		})
	}
}

func TestMatrix_Get(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		Matrix
		args args
		want float64
	}{
		{
			name: "three",
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
					7, 8, 9,
				},
				n: 3,
				m: 3,
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: 5,
		},
		{
			name: "one",
			Matrix: Matrix{
				data: []float64{
					42,
				},
				n: 1,
				m: 1,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Matrix.Get(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Matrix.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type args struct {
		i     int
		j     int
		value float64
	}
	tests := []struct {
		name string
		Matrix
		want Matrix
		args args
	}{
		{
			name:   "one",
			Matrix: *NewMatrix(1, 1),
			args: args{
				i:     0,
				j:     0,
				value: 42,
			},
			want: Matrix{
				data: []float64{
					42,
				},
				n: 1,
				m: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Matrix.Set(tt.args.i, tt.args.j, tt.args.value)
			if !matrixEqual(&tt.Matrix, &tt.want) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.want)
			}
		})
	}
}

func TestSolveSLAU(t *testing.T) {
	type args struct {
		matrix *Matrix
		b      []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				matrix: NewMatrix(0, 0),
				b:      []float64{},
			},
			want: []float64{},
		},
		{
			name:    "nil",
			wantErr: true,
		},
		{
			name: "not square",
			args: args{
				matrix: &Matrix{
					data: []float64{
						1, 0,
					},
					n: 1,
					m: 2,
				},
			},
			wantErr: true,
		},
		{
			name: "simple",
			args: args{
				matrix: &Matrix{
					data: []float64{
						1, 0,
						0, 1,
					},
					n: 2,
					m: 2,
				},
				b: []float64{3, 1},
			},
			want: []float64{3, 1},
		},
		{
			name: "simple rotated",
			args: args{
				matrix: &Matrix{
					data: []float64{
						0, 1,
						1, 0,
					},
					n: 2,
					m: 2,
				},
				b: []float64{3, 1},
			},
			want: []float64{1, 3},
		},
		{
			name: "hard",
			args: args{
				matrix: &Matrix{
					data: []float64{
						1, 2,
						3, -1,
					},
					n: 2,
					m: 2,
				},
				b: []float64{11, 12},
			},
			want: []float64{5, 3},
		},
		{
			name: "harder",
			args: args{
				matrix: &Matrix{
					data: []float64{
						1, -2, 1,
						2, 2, -1,
						4, -1, 1,
					},
					n: 3,
					m: 3,
				},
				b: []float64{0, 3, 5},
			},
			want: []float64{1, 2, 3},
		},
		{
			name: "big",
			args: args{
				matrix: &Matrix{
					data: []float64{
						-1, -8, 0, 5,
						6, -6, 2, 4,
						9, -5, -6, 4,
						-5, 0, -9, 1,
					},
					n: 4,
					m: 4,
				},
				b: []float64{-60, -10, 65, 18},
			},
			want: []float64{7, 6, -6, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveSLAU(tt.args.matrix, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveSLAU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !floatEqual(got, tt.want) {
				t.Errorf("SolveSLAU() = %v, want %v", got, tt.want)
			}
		})
	}
}
