package graphics

import (
	"CretRT/pkg/matht"
	"fmt"
	"math"
	"strings"
)

type Canvas struct {
	Width int
	Height int
	Arr [][]matht.Tuple
}

func CreateCanvas (width int, height int) Canvas {
	canvas := Canvas{Width: width, Height: height}
	canvas.Arr = make([][]matht.Tuple, width)
	for i := range canvas.Arr {
		canvas.Arr[i] = make([]matht.Tuple, height)
	}

	return canvas
}

func (canvas *Canvas) WritePixel(x int, y int, color matht.Tuple) {
	canvas.Arr[x][y] = color
}

func (canvas *Canvas) PixelAt(x int, y int) matht.Tuple {
	return canvas.Arr[x][y]
}

func (canvas *Canvas) Fill(color matht.Tuple) {
	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			canvas.Arr[x][y] = color
		}
	}
}

func CanvasToPPM(canvas *Canvas) string {
	var sb strings.Builder

	sb.WriteString("P3\n")
	sb.WriteString(fmt.Sprintf("%d %d\n", canvas.Width, canvas.Height))
	sb.WriteString("255\n")

	var sb1 strings.Builder
	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			if sb1.Len() + 12 > 70 {
				sb1.WriteRune('\n')
				sb.WriteString(sb1.String())
				sb1.Reset()
			}else{
				if sb1.Len() != 0 {
					sb1.WriteRune(' ')
				}
			}

			newR := int(math.Max(0, math.Min(255, canvas.Arr[x][y].X * 255)))
			newG := int(math.Max(0, math.Min(255, canvas.Arr[x][y].Y * 255)))
			newB := int(math.Max(0, math.Min(255, canvas.Arr[x][y].Z * 255)))

			sb1.WriteString(fmt.Sprintf("%d %d %d", newR, newG, newB))
		}

		sb1.WriteRune('\n')
		sb.WriteString(sb1.String())
		sb1.Reset()
	}

	return sb.String()
}