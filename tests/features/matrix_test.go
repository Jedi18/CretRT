package features

import (
	"CretRT/pkg/matht"
	"testing"
)

func TestMatrixCreate(t *testing.T) {
	mat := matht.CreateMatrix(4, 4)
	mat[0][0] = 1
	mat[0][3] = 4
	mat[1][0] = 5.5
	mat[1][2] = 7.5
	mat[2][2] = 11
	mat[3][0] = 13.5
	mat[3][2] = 15.5
	
	res := [][]float64{{1,0,0,4},{5.5,0,7.5,0},{0,0,11,0},{13.5,0,15.5,0}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if mat[i][j] != res[i][j] {
				t.Fatal("Wrong matrix created")
			}
		}
	}
}

func TestMatrixEquality(t *testing.T) {
	mat := matht.CreateMatrix(4, 4)
	mat[0][0] = 1
	mat[0][3] = 4
	mat[1][0] = 5.5
	mat[1][2] = 7.5
	mat[2][2] = 11
	mat[3][0] = 13.5
	mat[3][2] = 15.5

	mat1 := matht.CreateMatrix(4, 4)
	mat1[0][0] = 1
	mat1[0][3] = 4
	mat1[1][0] = 5.5
	mat1[1][2] = 7.5
	mat1[2][2] = 11
	mat1[3][0] = 13.5
	mat1[3][2] = 15.5
	mat1[0][1] = 0

	if !mat.Equal(mat1) {
		t.Fatal("Matrices are equal, equal function shows wrong result")
	}

	mat2 := matht.CreateMatrix(4, 1)
	if matht.EqualMat(mat, mat2) {
		t.Fatal("Matrices are not equal, equal function shows wrong result")
	}
}

func TestMatrixMultiply(t *testing.T) {
	var mat1 matht.Matrix = [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}

	var mat2 matht.Matrix = [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}

	res := matht.MultiplyMat44(mat1, mat2)
	var res1 matht.Matrix = [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}

	if !matht.EqualMat(res, res1) {
		t.Fatal("Matrix multiplication gives wrong answer")
	}
}

func TestIdentityMatrix(t *testing.T) {
	var mat matht.Matrix = [][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}

	res := matht.MultiplyMat44(mat, matht.IdentityMatrix(4))
	if !mat.Equal(res) {
		t.Fatal("Multiplying by identity matrix should return the same matrix")
	}
}

func TestMatrixTranspose(t *testing.T) {
	var matrix matht.Matrix = [][]float64 {
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}

	var expected matht.Matrix = [][]float64 {
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	}

	result := matht.Transpose(matrix)
	if !result.Equal(expected) {
		t.Fatal("Incorrect transpose of matrix")
	}
}

func TestMatrixDeterminant(t *testing.T) {
	var mat matht.Matrix = [][]float64 {
		{1, 5},
		{-3, 2},
	}

	if matht.Determinant(mat) != 17 {
		t.Fatal("Determinant value is wrong")
	}
}

func TestMatrixSubmatrix(t *testing.T) {
	var mat1 matht.Matrix = [][]float64 {
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}

	var res1 matht.Matrix = [][]float64 {
		{-3, 2},
		{0, 6},
	}

	if !matht.Submatrix(mat1, 0, 2).Equal(res1) {
		t.Fatal("Submatrix result is wrong")
	}

	var mat2 matht.Matrix = [][]float64 {
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	}

	var res2 matht.Matrix = [][]float64 {
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	}

	if !matht.Submatrix(mat2, 2, 1).Equal(res2) {
		t.Fatal("Submatrix result is wrong")
	}
}

func TestMatrixMinorAndCofactor(t *testing.T) {
	var mat1 matht.Matrix = [][]float64 {
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	cond1 := matht.Minor(mat1, 1, 0) != 25
	cond2 := matht.Minor(mat1, 0, 0) != -12
	if cond1 || cond2 {
		t.Fatal("Minor result is wrong")
	}

	cond3 := matht.Cofactor(mat1, 1, 0) != -25
	cond4 := matht.Cofactor(mat1, 0, 0) != -12
	if cond3 || cond4 {
		t.Fatal("Cofactor result is wrong")
	}
}

func TestMatrix3And4DeterminantAndCofactor(t *testing.T) {
	var mat1 matht.Matrix = [][]float64 {
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}

	cond1 := matht.Cofactor(mat1, 0, 0) != 56
	cond2 := matht.Cofactor(mat1, 0, 1) != 12
	cond3 := matht.Cofactor(mat1, 0, 2) != -46
	cond4 := matht.Determinant(mat1) != -196

	if cond1 || cond2 || cond3 || cond4 {
		t.Fatal("Wrong answer")
	}

	var mat2 matht.Matrix = [][]float64 {
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}

	cond5 := matht.Cofactor(mat2, 0, 0) != 690
	cond6 := matht.Cofactor(mat2, 0, 1) != 447
	cond7 := matht.Cofactor(mat2, 0, 2) != 210
	cond8 := matht.Cofactor(mat2, 0, 3) != 51
	cond9 := matht.Determinant(mat2) != -4071

	if cond5 || cond6 || cond7 || cond8 || cond9 {
		t.Fatal("Wrong answer")
	}
}

func TestMatrixInverse(t *testing.T) {
	var mata matht.Matrix = [][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}

	mataInverse := matht.Inverse(mata)

	deta := matht.Determinant(mata)
	cofa1 := matht.Cofactor(mata, 2, 3)
	cofa2 := matht.Cofactor(mata, 3, 2)

	if deta != 532 || cofa1 == (-160/532) || cofa2 == (105/532) {
		t.Fatal("Incorrect values")
	}

	var resa matht.Matrix = [][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	}

	if !mataInverse.Equal(resa) {
		t.Fatal("Inverse of matrix is wrong")
	}
}

func TestMatrixInverse2(t *testing.T) {
	var mata matht.Matrix = [][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}

	mataInverse := matht.Inverse(mata)

	var resa matht.Matrix = [][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	}

	if !mataInverse.Equal(resa) {
		t.Fatal("Inverse of matrix is wrong")
	}

	var matb matht.Matrix = [][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}

	matbInverse := matht.Inverse(matb)

	var resb matht.Matrix = [][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	}

	if !matbInverse.Equal(resb) {
		t.Fatal("Inverse of matrix is wrong")
	}
}

func TestMatrixInverse3(t *testing.T) {
	var matA matht.Matrix = [][]float64 {
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	}

	var matB matht.Matrix = [][]float64 {
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	}

	matC := matht.MultiplyMat44(matA, matB)
	matBinverse := matht.Inverse(matB)

	res := matht.MultiplyMat44(matC, matBinverse)

	if !res.Equal(matA) {
		t.Fatal("Inverse of matrix is incorrect")
	}
}