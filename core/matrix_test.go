package core

import (
	"reflect"
	"testing"
)

func TestMatrix_ProdMatrix(t *testing.T) {
	type fields struct {
		data []float64
		n    int
		m    int
	}
	tests := []struct {
		name    string
		fields  fields
		arg     Matrix
		want    Matrix
		wantErr bool
	}{
		{
			name: "E",
			fields: fields{
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
			fields: fields{
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
			fields: fields{
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
			fields: fields{
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
			core := Matrix{
				data: tt.fields.data,
				n:    tt.fields.n,
				m:    tt.fields.m,
			}
			got, err := core.ProdMatrix(&tt.arg)
			if err != nil {
				if !tt.wantErr {
					t.Error(err)
				}
				return
			}
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("Given:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}

	if tests[0].arg.String() == "" {
		t.Error("oops")
	}
}

func TestMatrixSwapLines(t *testing.T) {
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
			if !reflect.DeepEqual(tt.Matrix, tt.wantMatrix) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.wantMatrix)
			}
		})
	}
}

func TestMatrix_LU_decomposition(t *testing.T) {
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
		{
			name: "not square",
			Matrix: Matrix{
				data: []float64{
					1, 2, 3,
					4, 5, 6,
				},
				n: 2,
				m: 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			L, U, P, err := tt.Matrix.LUDecomposition()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Matrix.LU_decomposition() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(L, tt.L) {
				t.Errorf("Given:\n%s\nWant:\n%s", L, tt.L)
			}
			if !reflect.DeepEqual(U, tt.U) {
				t.Errorf("Given:\n%s\nWant:\n%s", U, tt.U)
			}
			if !reflect.DeepEqual(P, tt.P) {
				t.Errorf("Given:\n%s\nWant:\n%s", P, tt.P)
			}
		})
	}
}
