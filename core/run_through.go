package core

func RunThrough(matrix *Matrix, col Coloumn) Coloumn {
	if matrix == nil || matrix.n != matrix.m {
		return nil
	}
	var (
		P   = make(Coloumn, matrix.n)
		Q   = make(Coloumn, matrix.n)
		res = make(Coloumn, matrix.n)
	)
	P[0] = -matrix.data[1] / matrix.data[0]
	Q[0] = col[0] / matrix.data[0]

	line := matrix.m
	for i := 1; i < matrix.n; i++ {
		if i == matrix.n-1 {
			P[i] = 0
		} else {
			P[i] = -matrix.data[line+2] / (matrix.data[line+1] + matrix.data[line]*P[i-1])
		}
		Q[i] = (col[i] - matrix.data[line]*Q[i-1]) / (matrix.data[line+1] + matrix.data[line]*P[i-1])
		line += matrix.m + 1
	}

	res[len(res)-1] = Q[len(res)-1]
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = P[i]*res[i+1] + Q[i]
	}

	return res
}
