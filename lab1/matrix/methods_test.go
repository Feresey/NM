package matrix

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
	type args struct {
		right *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
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
			args: args{
				right: &Matrix{
					data: []float64{
						1, 2, 3,
						4, 5, 6,
						7, 8, 9,
					},
					n: 3,
					m: 3,
				},
			},
			want: &Matrix{
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
			args: args{
				right: &Matrix{
					data: []float64{
						1, 2,
						3, 4,
						5, 6,
					},
					n: 3,
					m: 2,
				},
			},
			want: &Matrix{
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
			args: args{
				right: &Matrix{
					data: []float64{
						8,
						7,
						2,
					},
					n: 3,
					m: 1,
				},
			},
			want: &Matrix{
				data: []float64{
					21,
					46,
					4,
				},
				n: 3,
				m: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			left := &Matrix{
				data: tt.fields.data,
				n:    tt.fields.n,
				m:    tt.fields.m,
			}
			if got := left.ProdMatrix(tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Given:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}
}
