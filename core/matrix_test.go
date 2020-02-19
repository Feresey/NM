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
			if (err != nil) != tt.wantErr {
				t.Error(err)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(*got, tt.want) {
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
			if !reflect.DeepEqual(tt.Matrix, tt.wantMatrix) {
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
			if !reflect.DeepEqual(L, tt.L) {
				t.Errorf("L Given:\n%s\nWant:\n%s", L, tt.L)
			}
			if !reflect.DeepEqual(U, tt.U) {
				t.Errorf("U Given:\n%s\nWant:\n%s", U, tt.U)
			}
			if !reflect.DeepEqual(P, tt.P) {
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
			if !tt.wantErr && !reflect.DeepEqual(tt.Matrix, tt.want) {
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
			if !reflect.DeepEqual(tt.Matrix, tt.want) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolveSLAU(tt.args.matrix, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveSLAU() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveSLAU() = %v, want %v", got, tt.want)
			}
		})
	}
}
