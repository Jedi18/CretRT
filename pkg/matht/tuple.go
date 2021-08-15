package matht

import "math"

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t *Tuple) IsPoint() bool {
	return t.W == 1
}

func (t *Tuple) IsVector() bool {
	return t.W == 0
}

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func Add(t1 Tuple, t2 Tuple) Tuple {
	return Tuple{t1.X + t2.X, t1.Y + t2.Y, t1.Z + t2.Z, t1.W + t2.W}
}

func (t *Tuple) Add(t1 Tuple) {
	t.X += t1.X
	t.Y += t1.Y
	t.Z += t1.Z
	t.W += t1.W
}

func Sub(t1 Tuple, t2 Tuple) Tuple {
	return Tuple{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z, t1.W - t2.W}
}

func (t *Tuple) Sub(t1 Tuple) {
	t.X -= t1.X
	t.Y -= t1.Y
	t.Z -= t1.Z
	t.W -= t1.W
}

func Equal(t1 Tuple, t2 Tuple) bool {
	checkX := math.Abs(t1.X - t2.X) < 0.00001
	checkY := math.Abs(t1.Y - t2.Y) < 0.00001
	checkZ := math.Abs(t1.Z - t2.Z) < 0.00001
	checkW := math.Abs(t1.W - t2.W) < 0.00001
	return checkX && checkY && checkZ && checkW
}

func Negate(t1 Tuple) Tuple {
	return Tuple{-t1.X, -t1.Y, -t1.Z, -t1.W}
}

func (t *Tuple) Negate() {
	t.X *= -1
	t.Y *= -1
	t.Z *= -1
	t.W *= -1
}

func (t *Tuple) Multiply(scalar float64) {
	t.X *= scalar
	t.Y *= scalar
	t.Z *= scalar
	t.W *= scalar
}

func Multiply(t Tuple, scalar float64) Tuple {
	t.Multiply(scalar)
	return t
}

func (t *Tuple) Divide(scalar float64) {
	t.X /= scalar
	t.Y /= scalar
	t.Z /= scalar
	t.W /= scalar
}

func Divide(t Tuple, scalar float64) Tuple {
	t.Divide(scalar)
	return t
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t *Tuple) Normalize() {
	mag := t.Magnitude()
	t.Divide(mag)
}

func Normalize(t Tuple) Tuple {
	t.Normalize()
	return t
}

func (t *Tuple) Dot(t1 Tuple) float64 {
	res := t.X * t1.X
	res += t.Y * t1.Y
	res += t.Z * t1.Z
	res += t.W * t1.W
	return res
}

func Dot(t1 Tuple, t2 Tuple) float64 {
	return t1.Dot(t2)
}

// Cross should be only called for vectors
func (t *Tuple) Cross(t1 Tuple) Tuple {
	return Vector(t.Y*t1.Z-t.Z*t1.Y, t.Z*t1.X-t.X*t1.Z, t.X*t1.Y-t.Y*t1.X)
}

func Cross(t1 Tuple, t2 Tuple) Tuple {
	return t1.Cross(t2)
}

func Color(r float64, g float64, b float64) Tuple {
	return Tuple{r, g, b, 0}
}

func HadamardProduct(t1 Tuple, t2 Tuple) Tuple {
	t1.X *= t2.X
	t1.Y *= t2.Y
	t1.Z *= t2.Z
	return t1
}