package matht

import (
	"math"
	"testing"
)

func TestTranslation(t *testing.T) {
	transformMatrix := Translation(5, -3, 2)
	p := Point(-3, 4, 5)
	res := transformMatrix.MultiplyTuple(p)

	if !Equal(res, Point(2, 1, 7)) {
		t.Fatal("Wrong translation value")
	}

	tinv := Inverse(transformMatrix)
	res2 := tinv.MultiplyTuple(p)

	if !Equal(res2, Point(-8, 7, 3)) {
		t.Fatal("Wrong translation value")
	}

	v := Vector(-3, 4, 5)
	res3 := transformMatrix.MultiplyTuple(v)

	if !Equal(v, res3) {
		t.Fatal("Translation should not affect vectors")
	}
}

func TestScaling(t *testing.T) {
	scalingMatrix := Scaling(2, 3, 4)
	p1 := Point(-4, 6, 8)
	res := scalingMatrix.MultiplyTuple(p1)

	if !Equal(res, Point(-8, 18, 32)) {
		t.Fatal("Wrong scaling answer")
	}

	v1 := Vector(-4, 6, 8)
	res1 := scalingMatrix.MultiplyTuple(v1)

	if !Equal(res1, Vector(-8, 18, 32)) {
		t.Fatal("Wrong scaling answer")
	}

	invScalingMatrix := Inverse(scalingMatrix)
	res2 := invScalingMatrix.MultiplyTuple(v1)

	if !Equal(res2, Vector(-2, 2, 2)) {
		t.Fatal("Wrong scaling answer")
	}

	// reflect on x axis
	scalingMatrix2 := Scaling(-1, 1, 1)
	p2 := Point(2, 3, 4)
	res3 := scalingMatrix2.MultiplyTuple(p2)

	if !Equal(res3, Point(-2, 3, 4)) {
		t.Fatal("Wrong scaling answer")
	}
}

func TestRotationX(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationX(math.Pi / 4)
	fullQuarter := RotationX(math.Pi / 2)
	halfQuarterInv := Inverse(halfQuarter)

	res1 := halfQuarter.MultiplyTuple(p)
	res2 := fullQuarter.MultiplyTuple(p)
	res3 := halfQuarterInv.MultiplyTuple(p)

	var tests = []struct{
		result   Tuple
		expected Tuple
	} {
		{res1, Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)},
		{res2, Point(0, 0, 1)},
		{res3, Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)},
	}

	for _, test := range tests {
		if !Equal(test.result, test.expected) {
			t.Errorf("Wrong rotation x answer, result:%v expected:%v", test.result, test.expected)
		}
	}
}

func TestRotationY(t *testing.T) {
	p := Point(0, 0, 1)
	halfQuarter := RotationY(math.Pi / 4)
	fullQuarter := RotationY(math.Pi / 2)

	res1 := halfQuarter.MultiplyTuple(p)
	res2 := fullQuarter.MultiplyTuple(p)

	var tests = []struct{
		result   Tuple
		expected Tuple
	} {
		{res1, Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)},
		{res2, Point(1, 0, 0)},
	}

	for _, test := range tests {
		if !Equal(test.result, test.expected) {
			t.Errorf("Wrong rotation y answer, result:%v expected:%v", test.result, test.expected)
		}
	}
}

func TestRotationZ(t *testing.T) {
	p := Point(0, 1, 0)
	halfQuarter := RotationZ(math.Pi / 4)
	fullQuarter := RotationZ(math.Pi / 2)

	res1 := halfQuarter.MultiplyTuple(p)
	res2 := fullQuarter.MultiplyTuple(p)

	var tests = []struct{
		result   Tuple
		expected Tuple
	} {
		{res1, Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)},
		{res2, Point(-1, 0, 0)},
	}

	for _, test := range tests {
		if !Equal(test.result, test.expected) {
			t.Errorf("Wrong rotation z answer, result:%v expected:%v", test.result, test.expected)
		}
	}
}

func TestShearing(t *testing.T) {
	p := Point(2, 3, 4)

	var tests = []struct{
		result   Tuple
		expected Tuple
	} {
		{Shearing(1, 0, 0, 0, 0, 0).MultiplyTuple(p),
			Point(5, 3, 4)},
		{Shearing(0, 1, 0, 0, 0, 0).MultiplyTuple(p),
			Point(6, 3, 4)},
		{Shearing(0, 0, 1, 0, 0, 0).MultiplyTuple(p),
			Point(2, 5, 4)},
		{Shearing(0, 0, 0, 1, 0, 0).MultiplyTuple(p),
			Point(2, 7, 4)},
		{Shearing(0, 0, 0, 0, 1, 0).MultiplyTuple(p),
			Point(2, 3, 6)},
		{Shearing(0, 0, 0, 0, 0, 1).MultiplyTuple(p),
			Point(2, 3, 7)},
	}

	for _, test := range tests {
		if !Equal(test.result, test.expected) {
			t.Errorf("Wrong shearing answer, result:%v expected:%v", test.result, test.expected)
		}
	}
}