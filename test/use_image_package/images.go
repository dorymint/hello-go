
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
	"image"
)

func main2() {

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	main2()
}
