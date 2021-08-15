package matht

import "math"

type Matrix [][]float64

func CreateMatrix(m int, n int) Matrix {
	res := make([][]float64, m)
	for i := 0; i < m; i++ {
		res[i] = make([]float64, n)
	}

	return res
}

func (mat Matrix) Equal(mat1 Matrix) bool {
	if len(mat) != len(mat1) {
		return false
	}

	m := len(mat)
	for i := 0; i < m; i++ {
		if len(mat[i]) != len(mat1[i]) {
			return false
		}
		n := len(mat[i])

		for j := 0; j < n; j++ {
			diff := math.Abs(mat[i][j] - mat1[i][j])
			if diff >= 0.00001 {
				return false
			}
		}
	}

	return true
}

func EqualMat(mat1 Matrix, mat2 Matrix) bool {
	return mat1.Equal(mat2)
}

func MultiplyMat44(mat1 Matrix, mat2 Matrix) Matrix {
	m := len(mat1)

	if m != 4 || m != len(mat2) {
		panic("Invalid matrix sizes for multiplication")
	}

	n := len(mat1[0])

	if n != 4 || n != len(mat2[0]) {
		panic("Invalid matrix sizes for multiplication")
	}

	result := CreateMatrix(4, 4)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 4; k++ {
				result[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}

	return result
}

func IdentityMatrix(n int) Matrix {
	identityMat := CreateMatrix(n, n)
	for i := 0; i < n; i++ {
		identityMat[i][i] = 1
	}

	return identityMat
}

func Transpose(matrix Matrix) Matrix {
	m := len(matrix)
	if m == 0 {
		return matrix
	}

	n := len(matrix[0])
	result := CreateMatrix(n, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}

func Determinant(matrix Matrix) float64 {
	m := len(matrix)

	if m == 0 {
		return 0
	}

	n := len(matrix[0])

	if m != n {
		panic("M != N for calculating determinant")
	}

	if m == 1 {
		return matrix[0][0]
	}

	if m == 2 {
		return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
	}

	var det float64 = 0
	for j := 0; j < n; j++ {
		det += matrix[0][j] * Cofactor(matrix, 0, j)
	}

	return det
}

func Submatrix(matrix Matrix, r int, c int) Matrix {
	n := len(matrix)
	m := len(matrix[0])

	result := CreateMatrix(n-1, m-1)

	for i := 0; i < n; i++ {
		var newi int

		if i == r {
			continue
		} else if i < r {
			newi = i
		} else {
			newi = i - 1
		}

		for j := 0; j < m; j++ {
			var newj int

			if j == c {
				continue
			} else if j < c {
				newj = j
			} else {
				newj = j - 1
			}

			result[newi][newj] = matrix[i][j]
		}
	}

	return result
}

func Minor(matrix Matrix, r int, c int) float64 {
	return Determinant(Submatrix(matrix, r, c))
}

func Cofactor(matrix Matrix, r int, c int) float64 {
	minorVal := Minor(matrix, r, c)
	if (r+c) % 2 == 0 {
		return minorVal
	}else{
		return -minorVal
	}
}

func Inverse(matrix Matrix) Matrix {
	det := Determinant(matrix)

	if det == 0 {
		panic("Matrix is not invertible")
	}

	m := len(matrix)
	n := len(matrix[0])
	if n != m {
		panic("n != m")
	}

	res := CreateMatrix(m, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			cof := Cofactor(matrix, r, c)

			res[c][r] = cof / det
		}
	}

	return res
}

func (matrix Matrix) MultiplyTuple(tuple Tuple) Tuple {
	res := Tuple{}
	m := len(matrix)

	if m == 0 {
		return res
	}

	n := len(matrix[0])

	if m != 4 || n != 4 {
		panic("Can only multiple tuple with 4x4 matrix")
	}

	res.X = tuple.X * matrix[0][0] + tuple.Y * matrix[0][1] + tuple.Z * matrix[0][2] + tuple.W * matrix[0][3]
	res.Y = tuple.X * matrix[1][0] + tuple.Y * matrix[1][1] + tuple.Z * matrix[1][2] + tuple.W * matrix[1][3]
	res.Z = tuple.X * matrix[2][0] + tuple.Y * matrix[2][1] + tuple.Z * matrix[2][2] + tuple.W * matrix[2][3]
	res.W = tuple.X * matrix[3][0] + tuple.Y * matrix[3][1] + tuple.Z * matrix[3][2] + tuple.W * matrix[3][3]

	return res
}