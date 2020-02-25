package core

import (
	"reflect"
	"testing"
)

func TestMatrix_RunThrough(t *testing.T) {
	type args struct {
		col Coloumn
	}
	tests := []struct {
		name string
		*Matrix
		args args
		want Coloumn
	}{
		{
			name: "simple",
			Matrix: &Matrix{
				data: Row{
					8, -2, 0, 0,
					-1, 6, -2, 0,
					0, 2, 10, -4,
					0, 0, -1, 6,
				},
				n: 4,
				m: 4,
			},
			args: args{
				col: Coloumn{6, 3, 8, 5},
			},
			want: Coloumn{1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.Matrix.RunThrough(tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.RunThrough() = %v, want %v", got, tt.want)
			}
		})
	}
}
