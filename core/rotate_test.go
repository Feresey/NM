package core

import (
	"testing"
)

func TestRotations(t *testing.T) {
	type args struct {
		matrix *Matrix
		eps    float64
	}
	tests := []struct {
		name           string
		args           args
		wantSz         Coloumn
		wantSv         *Matrix
		wantIterations int
	}{
		{
			name: "simple",
			args: args{
				matrix: &Matrix{
					data: Row{
						4, 2, 1,
						2, 5, 3,
						1, 3, 6,
					},
					m: 3,
					n: 3,
				},
				eps: 0.3,
			},
			wantSv: &Matrix{
				data: Row{
					0.78, -0.5064, 0.361,
					0.2209, 0.7625, 0.6,
					-0.58, -0.398, 0.7,
				},
				n: 3,
				m: 3,
			},
			wantSz: Coloumn{3.706, 1.929, 9.38},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSz, gotSv, gotIterations := Rotations(tt.args.matrix, tt.args.eps)
			if !floatEqual(gotSz, tt.wantSz, tt.args.eps) {
				t.Errorf("Rotations() gotSz = %v, want %v", gotSz, tt.wantSz)
			}
			if !matrixEqual(gotSv, tt.wantSv, tt.args.eps) {
				t.Errorf("Given:\n%s\nWant:\n%s", gotSv, tt.wantSv)
			}
			if gotIterations != tt.wantIterations {
				t.Errorf("Rotations() gotIterations = %v, want %v", gotIterations, tt.wantIterations)
			}
		})
	}
}

func Test_getMaxAbsElem(t *testing.T) {
	tests := []struct {
		name  string
		m     *Matrix
		wantI int
		wantJ int
	}{
		{
			name: "simple",
			m: &Matrix{
				data: Row{
					1, 0, 0,
					0, 1, 2,
					0, 0, 1,
				},
				m: 3,
				n: 3,
			},
			wantI: 1,
			wantJ: 2,
		},
		{
			name: "negate",
			m: &Matrix{
				data: Row{
					1, 0, 0,
					0, 1, -2,
					0, 0, 1,
				},
				m: 3,
				n: 3,
			},
			wantI: 1,
			wantJ: 2,
		},
		{
			name: "empty",
			m: &Matrix{
				data: Row{
					1, 0, 0,
					0, 1, 0,
					0, 0, 1,
				},
				m: 3,
				n: 3,
			},
			wantI: -1,
			wantJ: -1,
		},
		{
			name: "fool",
			m: &Matrix{
				data: Row{
					1, 0, 0,
					5, 1, 0,
					12, 123, 1,
				},
				m: 3,
				n: 3,
			},
			wantI: -1,
			wantJ: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if i, j := getMaxAbsElem(tt.m); i != tt.wantI || j != tt.wantJ {
				t.Errorf("got:\n(%d, %d)\nwant:\n(%d, %d)", i, j, tt.wantI, tt.wantJ)
			}
		})
	}
}
