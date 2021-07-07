package hello

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

/*

学习处理图片，拿到图片rgba值，并处理为灰色图像，再生成一个新图片；

*/

func Test() {

	file, err := os.Open("./test.png")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Println("错误", err)
		return
	}

	result := parsePixels(img)

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	newImg, err := os.Create("./new.png")
	defer newImg.Close()
	buff := bufio.NewWriter(newImg)
	rgb := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			rgb.SetRGBA(x, y, result[x][y])
		}
	}

	png.Encode(buff, rgb)
	buff.Flush()
}

func parsePixels(img image.Image) [][]color.RGBA {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	var pixel [][]color.RGBA

	for x := 0; x < w; x++ {
		var row []color.RGBA
		for y := 0; y < h; y++ {
			row = append(row, transform(img.At(x, y).RGBA()))
		}
		pixel = append(pixel, row)
	}

	return pixel
}

func transform(r uint32, g uint32, b uint32, a uint32) color.RGBA {
	R, G, B := uint8(r/257), uint8(g/257), uint8(b/257)

	avg := (R + G + B) / 3

	return color.RGBA{avg, avg, avg, uint8(a / 257)}
}
