package matht

import "math"

func Translation(x float64, y float64, z float64) Matrix {
	res := IdentityMatrix(4)
	res[0][3] = x
	res[1][3] = y
	res[2][3] = z
	return res
}

func Scaling(x float64, y float64, z float64) Matrix {
	res := IdentityMatrix(4)
	res[0][0] = x
	res[1][1] = y
	res[2][2] = z
	return res
}

func RotationX(radians float64) Matrix {
	res := IdentityMatrix(4)
	res[1][1] = math.Cos(radians)
	res[1][2] = -math.Sin(radians)
	res[2][1] = math.Sin(radians)
	res[2][2] = math.Cos(radians)
	return res
}

func RotationY(radians float64) Matrix {
	res := IdentityMatrix(4)
	res[0][0] = math.Cos(radians)
	res[0][2] = math.Sin(radians)
	res[2][0] = -math.Sin(radians)
	res[2][2] = math.Cos(radians)
	return res
}

func RotationZ(radians float64) Matrix {
	res := IdentityMatrix(4)
	res[0][0] = math.Cos(radians)
	res[0][1] = -math.Sin(radians)
	res[1][0] = math.Sin(radians)
	res[1][1] = math.Cos(radians)
	return res
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	res := IdentityMatrix(4)
	res[0][1] = xy
	res[0][2] = xz
	res[1][0] = yx
	res[1][2] = yz
	res[2][0] = zx
	res[2][1] = zy
	return res
}