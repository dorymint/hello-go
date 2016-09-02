
// image packageのインターフェース
// {{{
// pakcage image
// 
// type Image interface {
// 	ColorModel() color.Model
// 	Bunds() Rectangle
// 	At(x, y int) color.Color
// }
// }}}

package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (im Image) ColorModel() color.Model {
	model := color.RGBAModel
	return model
}
func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.w, im.h)
}
func (im Image) At(x, y int) color.Color {
	return color.RGBA{ uint8(x), uint8(y), uint8(255), uint8(255) }
}


func main() {
	var c color.Gray
	fmt.Printf("%T, %v\n\n", c, c)
	c.Y = 128
	fmt.Println(c)

	r, g, b, a := c.RGBA()
	fmt.Println(r, g, b, a, "\n")

	m := Image{255, 255}

	// TODO:call pic.ShowImage
	pic.ShowImage(m)

}

// NOTE:難しく考えない
// インターフェースに求められているレシーバ関数を定義して
// インターフェースを共有する関数に投げればいい

