// ref <http://qiita.com/qt-luigi/items/ec6cd349259fe6cc29eb>.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

var (
	N int = 20
	//xm	float64	=	0.0
	//ym	float64	=	0.6
	//h	float64	=	0.6
)

var (
	width  int = 500
	height int = 500
)

var (
	// 白
	bgcolor color.Color = color.RGBA{255, 255, 255, 255}
	// 緑
	linecolor color.Color = color.RGBA{0, 128, 0, 255}
)

func W1x(x, y float64) float64 {
	return 0.836*x + 0.044*y
}
func W1y(x, y float64) float64 {
	return 0.836*y - 0.044*x + 0.169
}

func W2x(x, y float64) float64 {
	return -0.141*x + 0.302*y
}
func W2y(x, y float64) float64 {
	return 0.141*y + 0.302*x + 0.127
}

func W3x(x, y float64) float64 {
	return 0.141*x - 0.302*y
}
func W3y(x, y float64) float64 {
	return 0.141*y + 0.302*x + 0.169
}

func W4x(x, y float64) float64 {
	return 0
}
func W4y(x, y float64) float64 {
	return 0.175337 * y
}

func f(m *image.RGBA, k int, x, y float64) {
	comperison := func() bool { return rand.Float64() < 0.3 }
	if 0 < k {
		f(m, k-1, W1x(x, y), W1y(x, y))
		if comperison() {
			f(m, k-1, W2x(x, y), W2y(x, y))
		}
		if comperison() {
			f(m, k-1, W2x(x, y), W2y(x, y))
		}
		if comperison() {
			f(m, k-1, W3x(x, y), W3y(x, y))
		}
		if comperison() {
			f(m, k-1, W4x(x, y), W4y(x, y))
		}
	} else {
		var s float64 = 490.0
		m.Set(int(x*s+float64(width)*0.5), int(float64(height)-y*s), linecolor)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	// RGBA キャンバス
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	fmt.Printf("%T\n", m)

	fmt.Printf("%T, %v\n", draw.Src, draw.Src)
	fmt.Printf("%T, %v\n", image.ZP, image.ZP)

	imuni := &image.Uniform{bgcolor}
	fmt.Printf("%T, %v\n", imuni, imuni)

	draw.Draw(m, m.Bounds(), &image.Uniform{bgcolor}, image.ZP, draw.Src)

	f(m, N, 0, 0)

	filename := "shida.png"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("writed to %s\n", filename)
}
