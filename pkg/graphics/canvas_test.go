package graphics

import (
	"CretRT/pkg/matht"
	"strings"
	"testing"
)

func TestCanvasCreate(t *testing.T) {
	canvas := CreateCanvas(10, 20)

	if canvas.Width != 10 || canvas.Height != 20 {
		t.Fatal("Rows and columns of canvas not set properly")
	}

	if canvas.Width != len(canvas.Arr) || canvas.Height != len(canvas.Arr[0]) {
		t.Fatal("Rows and columns of canvas not set properly")
	}

	for r := 0; r < canvas.Width; r++ {
		for c := 0; c < canvas.Height; c++ {
			if !matht.Equal(canvas.Arr[r][c], matht.Color(0, 0, 0)) {
				t.Fatal("All colors should be initially initialized to zero")
			}
		}
	}
}

func TestWriteAndReadPixel(t *testing.T) {
	canvas := CreateCanvas(10, 20)
	col := matht.Color(1, 0, 0)
	canvas.WritePixel(4, 15, col)

	if !matht.Equal(canvas.Arr[4][15], col) {
		t.Fatal("Write pixel did not write the pixel properly")
	}

	if !matht.Equal(canvas.PixelAt(4, 15), col) {
		t.Fatal("Pixel at did not read the correct pixel")
	}
}

func TestCanvasToPPM(t *testing.T) {
	c := CreateCanvas(5, 3)
	c.WritePixel(0, 0, matht.Color(1.5, 0, 0))
	c.WritePixel(2, 1, matht.Color(0, 0.5, 0))
	c.WritePixel(4, 2, matht.Color(-0.5, 0, 1))

	str := "P3\n5 3\n255\n255 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0\n0 0 0 0 127 0 0 0 0\n0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 255\n"
	if str != CanvasToPPM(&c) {
		t.Fatal("Incorrect Canvas to PPM conversion")
	}
}

func TestCanvasPPMLineLength(t *testing.T) {
	c := CreateCanvas(10, 2)
	c.Fill(matht.Color(1, 0.8, 0.6))

	ppmStr := CanvasToPPM(&c)

	for _, line := range strings.Split(ppmStr, "\n") {
		if len(line) > 70 {
			t.Fatal("Line length of PPM file above 70 characters")
		}
	}
}