package img

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"strconv"
	"time"
)

type Img struct {
	Pixels [][]Pixel
}

type Pixel struct {
	R  uint32
	G  uint32
	B  uint32
	A  uint32
	px int
	py int
}

func (c *Pixel) Equal(pixel Pixel) bool {
	if c.A == pixel.A && c.R == pixel.R && c.B == pixel.B && c.G == pixel.G {
		return true
	}
	return false
}

func (c *Pixel) ToString() string {
	a := strconv.Itoa(int(c.A))
	b := strconv.Itoa(int(c.B))
	g := strconv.Itoa(int(c.G))
	r := strconv.Itoa(int(c.R))
	return fmt.Sprintf("%s%s%s%s", r, g, b, a)
}

func ReadImg(filePath string) (img Img, samePixel map[string][]Pixel, err error) {
	samePixel = make(map[string][]Pixel)
	start := time.Now()
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	imgBytes := bytes.NewBuffer(file)
	theImg, err := png.Decode(imgBytes)
	bounds := theImg.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	var pixels [][]Pixel
	for i := 0; i < dx; i++ {
		var linePixels []Pixel
		for j := 0; j < dy; j++ {
			colorRgb := theImg.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			pixel := Pixel{
				r,
				g,
				b,
				a,
				i,
				j,
			}
			samePixel[pixel.ToString()] = append(samePixel[pixel.ToString()], pixel)
			linePixels = append(linePixels, pixel)
		}
		pixels = append(pixels, linePixels)
	}
	img.Pixels = pixels
	fmt.Println(time.Now().Sub(start).Seconds())
	return
}
