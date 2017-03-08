package mandel

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Settings defines the static settings of an produced immage
type Settings struct {
	Width, Height, Zoom, XShift, YShift float64
}

var (
	// ColorCode is a map for the color scheme
	ColorCode map[int]color.Color
	colorNum  int

	//settings are the used settings
	settings = Settings{
		3000,
		2000,
		1,
		0,
		0,
	}

	minRe     = -2.0
	maxRe     = 1.0
	minIm     = -1.2
	maxIm     = minIm + (maxRe-minRe)*settings.Height/settings.Width
	re_factor = (maxRe - minRe) / (settings.Width - 1)
	im_factor = (maxIm - minIm) / (settings.Height - 1)
)

func init() {
	SetDefaultColors()
}

//Width returns the width of the used image
func Width() int {
	return int(settings.Width)
}

//Height returns the height of the used immage
func Height() int {
	return int(settings.Height)
}

// SetSettings changes the used Settings and updates all internal values
func SetSettings(set Settings) {
	settings = set
	maxIm = minIm + (maxRe-minRe)*settings.Height/settings.Width
	re_factor = (maxRe - minRe) / (settings.Width - 1)
	im_factor = (maxIm - minIm) / (settings.Height - 1)
}

// AddColor adds a new color to the ColorCode map in the next free Slot
func AddColor(r, g, b int) {
	ColorCode[colorNum] = makeColor(r, g, b)
	colorNum++
}

//GetNumberOfRegisteredColors returns the number of registered colors
func GetNumberOfRegisteredColors() int {
	return len(ColorCode)
}

// SetDefaultColors sets the default color scheme
func SetDefaultColors() {
	ColorCode = map[int]color.Color{
		0:  makeColor(66, 30, 15),
		1:  makeColor(25, 7, 26),
		2:  makeColor(9, 1, 47),
		3:  makeColor(4, 4, 73),
		4:  makeColor(0, 7, 100),
		5:  makeColor(12, 44, 138),
		6:  makeColor(24, 82, 177),
		7:  makeColor(57, 125, 209),
		8:  makeColor(134, 181, 229),
		9:  makeColor(211, 236, 248),
		10: makeColor(241, 233, 191),
		11: makeColor(248, 201, 95),
		12: makeColor(255, 170, 0),
		13: makeColor(204, 128, 0),
		14: makeColor(153, 87, 0),
		15: makeColor(106, 52, 3),
	}
}

var colorCode map[int]color.Color

// SaveImage saves a image to mandel.png
func SaveImage(im *image.RGBA) error {
	file, err := os.Create("mandel.png")
	if err != nil {
		return err
	}
	err = png.Encode(file, im)
	if err != nil {
		return err
	}
	return file.Close()
}

// NewImage creates a new immage from the settings
func NewImage() *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, int(settings.Width), int(settings.Height)))
}

// Mandel performs the Mandelbrot operation on two complex numbers z,c where c has to be constant between iterations
func Mandel(z, c complex128) complex128 {
	return z*z + c
}

// Norm return a new complex number from two integers that is scaled by to the image settings
func Norm(col, row int) complex128 {
	xi := minRe + (float64(col)+(settings.Zoom*settings.XShift))/settings.Zoom*re_factor
	yi := minIm + (float64(row)+(settings.Zoom*settings.YShift))/settings.Zoom*im_factor
	return complex(xi, yi)
}

func makeColor(r, g, b int) color.Color {
	return &color.RGBA{uint8(r), uint8(g), uint8(b), uint8(255)}
}
