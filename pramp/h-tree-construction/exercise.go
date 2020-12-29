package exercise

import "math"

type drawLine func(x1, y1, x2, y2 float32)

func drawHTree(x, y, length float32, depth int, dl drawLine) {
	if depth == 0 {
		return
	}

	dl(x-length/2, y, x+length/2, y)
	dl(x-length/2, y-length/2, x-length/2, y+length/2)
	dl(x+length/2, y-length/2, x+length/2, y+length/2)

	drawHTree(x+length/2, y+length/2, length/math.Sqrt2, depth-1, dl)
	drawHTree(x+length/2, y-length/2, length/math.Sqrt2, depth-1, dl)
	drawHTree(x-length/2, y-length/2, length/math.Sqrt2, depth-1, dl)
	drawHTree(x-length/2, y+length/2, length/math.Sqrt2, depth-1, dl)
}
