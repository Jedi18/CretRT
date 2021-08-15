package matht

import (
	"math"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	p1 := Tuple{X: 1, Y: 2, Z: 3, W: 1}

	if !p1.IsPoint() {
		t.Fatal("Tuple is a point but IsPoint returned false")
	}

	p2 := Tuple{X: 1, Z: 3}

	if p2.IsPoint() {
		t.Fatal("Tuple is not a point but IsPoint return true")
	}
}

func TestTupleIsVector(t *testing.T) {
	p1 := Tuple{X: 1, Y: 2, Z: 3, W: 1}

	if p1.IsVector() {
		t.Fatal("Tuple is not a vector but IsPoint return true")
	}

	p2 := Tuple{X: 1, Z: 3}

	if !p2.IsVector() {
		t.Fatal("Tuple is a vector but IsPoint returned false")
	}
}

func TestTupleValues(t *testing.T) {
	p1 := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1.0}

	if p1.X != 4.3 || p1.Y != -4.2 || p1.Z != 3.1 || p1.W != 1.0 {
		t.Fatal("Tuple values don't match")
	}
}

func TestPointAndVectorFunc(t *testing.T) {
	p1 := Point(4, -4, 3)

	if p1.X != 4 || p1.Y != -4 || p1.Z != 3 || p1.W != 1 {
		t.Fatal("Tuple created by Point is wrong")
	}

	p2 := Vector(4, -4, 3)

	if p2.X != 4 || p2.Y != -4 || p2.Z != 3 || p2.W != 0 {
		t.Fatal("Tuple created by Vector is wrong")
	}
}

func TestAddTuple(t *testing.T) {
	p1 := Tuple{X: 3, Y: -2, Z: 5, W: 1}
	p2 := Tuple{X: -2, Y: 3, Z: 1}

	p3 := Add(p1, p2)

	if !Equal(p3, Tuple{X: 1, Y: 1, Z: 6, W: 1}) {
		t.Fatal("Addition of tuples gives wrong answer")
	}

	p1.Add(p2)

	if !Equal(p1, Tuple{X: 1, Y: 1, Z: 6, W: 1}) {
		t.Fatal("Addition of tuples gives wrong answer")
	}
}

func TestSubTuple(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	p3 := Sub(p1, p2)

	if !Equal(p3, Vector(-2, -4, -6)) {
		t.Fatal("Subtraction of tuples gives wrong answer")
	}

	p1.Sub(p2)

	if !Equal(p1, Vector(-2, -4, -6)) {
		t.Fatal("Subtraction of tuples gives wrong answer")
	}
}

func TestNegateTuple(t *testing.T) {
	p1 := Tuple{X: 1, Y: -2, Z: 3}
	p2 := Negate(p1)
	p1.Negate()

	if !Equal(p1, Tuple{X: -1, Y: 2, Z: -3}) {
		t.Fatal("Negation of tuple gives wrong answer")
	}

	if !Equal(p2, Tuple{X: -1, Y: 2, Z: -3}) {
		t.Fatal("Negation of tuple gives wrong answer")
	}
}

func TestMultiplyTuple(t *testing.T) {
	p1 := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	p2 := Multiply(p1, 3.5)
	p1.Multiply(0.5)

	if !Equal(p2, Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14}) {
		t.Fatal("Multiplication of tuple gives wrong answer")
	}

	if !Equal(p1, Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}) {
		t.Fatal("Multiplication of tuple gives wrong answer")
	}
}

func TestDivideTuple(t *testing.T) {
	p1 := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	p2 := Divide(p1, 2)
	p1.Divide(2)

	if !Equal(p2, Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}) {
		t.Fatal("Division of tuple gives wrong answer")
	}

	if !Equal(p1, Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2}) {
		t.Fatal("Division of tuple gives wrong answer")
	}
}

func TestMagnitude(t *testing.T) {
	p1 := Tuple{Y: 1}
	p2 := Tuple{W: 1}
	p3 := Tuple{X: 1, Y: 2, Z: 3}
	p4 := Tuple{X: -1, Y: -2, Z: -3}

	if p1.Magnitude() != 1 {
		t.Fatal("Magnitude of tuple gives wrong answer")
	}

	if p2.Magnitude() != 1 {
		t.Fatal("Magnitude of tuple gives wrong answer")
	}

	if p3.Magnitude() != math.Sqrt(14) {
		t.Fatal("Magnitude of tuple gives wrong answer")
	}

	if p4.Magnitude() != math.Sqrt(14) {
		t.Fatal("Magnitude of tuple gives wrong answer")
	}
}

func TestNormalize(t *testing.T) {
	p1 := Vector(4, 0, 0)
	p1.Normalize()

	if p1.Magnitude() != 1 {
		t.Fatal("After normalizing, magnitude should be 1")
	}

	if !Equal(p1, Vector(1, 0, 0)) {
		t.Fatal("Normalization of tuple gives wrong tuple")
	}

	p2 := Vector(1, 2, 3)
	p3 := Normalize(p2)

	if p3.Magnitude() != 1 {
		t.Fatal("After normalizing, magnitude should be 1")
	}

	if !Equal(p3, Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))) {
		t.Fatal("Normalization of tuple gives wrong tuple")
	}
}

func TestDot(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	dotProduct := a.Dot(b)
	dotProductA := Dot(a, b)

	if dotProduct != 20 || dotProductA != 20 {
		t.Fatal("Dot product of the tuples gives wrong value")
	}
}

func TestCross(t *testing.T) {
	a := Vector(1, 2, 3)
	b := Vector(2, 3, 4)

	res1 := a.Cross(b)
	res2 := b.Cross(a)
	res3 := Cross(a, b)

	if !Equal(res1, Vector(-1, 2, -1)) {
		t.Fatal("Cross product of the tuples gives wrong answer")
	}

	if !Equal(res2, Vector(1, -2, 1)) {
		t.Fatal("Cross product of the tuples gives wrong answer")
	}

	if !Equal(res3, Vector(-1, 2, -1)) {
		t.Fatal("Cross product of the tuples gives wrong answer")
	}
}

func TestColorOperations(t *testing.T) {
	c1 := Color(0.9, 0.6, 0.75)
	c2 := Color(0.7, 0.1, 0.25)

	res := Add(c1, c2)

	if !Equal(res, Color(1.6, 0.7, 1.0)) {
		t.Fatal("Addition of colors gives wrong answer")
	}

	res1 := Sub(c1, c2)
	if !Equal(res1, Color(0.2, 0.5, 0.5)) {
		t.Fatal("Subtraction of colors gives wrong answer")
	}

	res2 := Multiply(Color(0.2, 0.3, 0.4), 2)

	if !Equal(res2, Color(0.4, 0.6, 0.8)) {
		t.Fatal("Multiplication of colors gives wrong answer")
	}
}

func TestHadamardProduct(t *testing.T) {
	c1 := Color(1, 0.2, 0.4)
	c2 := Color(0.9, 1, 0.1)
	res := HadamardProduct(c1, c2)

	if !Equal(res, Color(0.9, 0.2, 0.04)) {
		t.Fatal("Hadamard product of colors gives wrong answer")
	}
}