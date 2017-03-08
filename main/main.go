package main

import "github.com/s-gheldd/mandel"
import "math/cmplx"
import "image/color"

const (
	iterations = 100
)

func main() {

	image := mandel.NewImage()

	for row := 0; row < mandel.Height(); row++ {
		for col := 0; col < mandel.Width(); col++ {

			norm := mandel.Norm(col, row)
			man := mandel.Mandel(norm, norm)
			it := 1
			for cmplx.Abs(man) < 2 && it < iterations {
				man = mandel.Mandel(man, norm)
				it++
			}

			if it < iterations {
				image.Set(col, row, mandel.ColorCode[it%mandel.GetNumberOfRegisteredColors()])
			} else {
				image.Set(col, row, color.Black)
			}
		}
	}
	mandel.SaveImage(image)
}
