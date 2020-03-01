package core

import (
	"testing"
)

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
		t.Run(tt.name, func(t *testing.T) {
			got := Transponse(tt.m)

			if !matrixEqual(got, tt.want, EPS) {
				t.Errorf("L Given:\n%s\nWant:\n%s", got, tt.want)
			}
		})
	}
}
