package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	f1dot, err := os.Create("1dot.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f1dot.Close()

	f4dot, err := os.Create("4dot.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f4dot.Close()

	write := func(f *os.File, length int) error {
		img := image.NewNRGBA(image.Rect(0, 0, length, length))
		c := color.RGBA{0x00, 0xff, 0xff, 0xff} // cyan
		for y := 0; y < length; y++ {
			for x := 0; x < length; x++ {
				img.Set(x, y, c)
			}
		}
		var b []byte
		buf := bytes.NewBuffer(b)
		if err := png.Encode(buf, img); err != nil {
			return err
		}
		if _, err := fmt.Fprint(f, buf); err != nil {
			return err
		}
		return nil
	}

	if err := write(f1dot, 1); err != nil {
		log.Fatal(err)
	}
	if err := write(f4dot, 4); err != nil {
		log.Fatal(err)
	}
}
