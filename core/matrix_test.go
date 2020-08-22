package core

import (
	"testing"
)

func matrixEqual(a, b *Matrix, eps float64) bool {
	return floatEqual(a.data, b.data, eps)
}

func floatEqual(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if tmp := a[i] - b[i]; tmp > 0 && tmp > eps || tmp < 0 && tmp < -eps {
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
			got := tt.Matrix.ProdMatrix(&tt.arg)

			if !tt.wantErr && !matrixEqual(got, &tt.want, eps) {
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
			if !matrixEqual(&tt.Matrix, &tt.wantMatrix, eps) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.wantMatrix)
			}
		})
	}
}

func TestMatrix_findMaxInCol(t *testing.T) {
	tests := []struct {
		name string
		Matrix
		col  int
		from int
		want int
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
			col:  0,
			from: 0,
			want: 1,
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
			col:  1,
			from: 1,
			want: -1,
		},
		{
			name: "last",
			Matrix: Matrix{
				data: []float64{
					0, 1,
					1, 1,
				},
				n: 2,
				m: 2,
			},
			col:  1,
			from: 1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Matrix.maxInCol(tt.col, tt.from); got != tt.want {
				t.Errorf("Given: %d. Want: %d", got, tt.want)
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
			if got := tt.Matrix.At(tt.args.i, tt.args.j); got != tt.want {
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
			if !matrixEqual(&tt.Matrix, &tt.want, eps) {
				t.Errorf("Given:\n%s\nWant:\n%s", tt.Matrix, tt.want)
			}
		})
	}
}

func TestMatrix_String(t *testing.T) {
	tests := []struct {
		name string
		Matrix
		want string
	}{
		{
			name: "suqare",
			Matrix: Matrix{
				data: Row{
					1, 2,
					3, 4,
				},
				n: 2,
				m: 2,
			},
			want: "   1.00   2.00\n   3.00   4.00\n",
		},
		{
			name: "not suqare",
			Matrix: Matrix{
				data: Row{
					1, 2, 3,
					4, 5, 6,
				},
				n: 2,
				m: 3,
			},
			want: "   1.00   2.00   3.00\n   4.00   5.00   6.00\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Matrix.String(); got != tt.want {
				t.Errorf("Matrix.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransponse(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix
		want *Matrix
	}{
		{
			name: "true",
			m: &Matrix{
				data: Row{
					1, 2,
					3, 4,
				},
				n: 2,
				m: 2,
			},
			want: &Matrix{
				data: Row{
					1, 3,
					2, 4,
				},
				n: 2,
				m: 2,
			},
		},
		{
			name: "none",
			m: &Matrix{
				data: Row{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				},
				n: 3,
				m: 3,
			},
			want: &Matrix{
				data: Row{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				},
				n: 3,
				m: 3,
			},
		},
		{
			name: "not square",
			m: &Matrix{
				data: Row{
					1, 2, 3,
					4, 5, 6,
				},
				n: 2,
				m: 3,
			},
			want: &Matrix{
				data: Row{
					1, 4,
					2, 5,
					3, 6,
				},
				n: 3,
				m: 2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := Transponse(tt.m)

			if !matrixEqual(got, tt.want, eps) {
				t.Errorf("L Given:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}
}
