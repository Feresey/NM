package core

import (
	"testing"
)

func TestMatrix_RunThrough(t *testing.T) {
	type args struct {
		col Coloumn
	}
	tests := []struct {
		name string
		*Matrix
		args    args
		want    Coloumn
		wantErr bool
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
			got, err := RunThrough(tt.Matrix, tt.args.col)
			if (err != nil) != tt.wantErr {
				t.Error("Error: ", err)
			}
			if !floatEqual(got, tt.want, eps) {
				t.Errorf("Matrix.RunThrough() = %v, want %v", got, tt.want)
			}
		})
	}
}
