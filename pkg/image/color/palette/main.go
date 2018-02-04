package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	write := func(f *os.File, cp []color.Color) error {
		length := len(cp)
		img := image.NewNRGBA(image.Rect(0, 0, length, length))
		for y := 0; y < length; y++ {
			c := cp[y]
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

	writeFile := func(path string, cp []color.Color) error {
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := write(f, cp); err != nil {
			return err
		}
		return nil
	}

	if err := writeFile("websafe.png", palette.WebSafe); err != nil {
		log.Fatal(err)
	}
	if err := writeFile("plan9.png", palette.Plan9); err != nil {
		log.Fatal(err)
	}

	var pl []color.Color
	//pl = palette.WebSafe
	pl = []color.Color{color.RGBA{0xff, 0xff, 0xff, 0xff}}
	if err := writeFile("img.png", pl); err != nil {
		log.Fatal(err)
	}

	if err := whileMod("while.png", palette.WebSafe); err != nil {
		log.Fatal(err)
	}
}

func whileMod(path string, cp []color.Color) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	length := len(cp)
	img := image.NewNRGBA(image.Rect(0, 0, length, length))
	for i := 1; i < 10; i++ {
		f.Truncate(0)
		for y := 0; y < length; y++ {
			c := cp[y]
			for x := 0; x < length; x++ {
				if i%2 == 0 {
					img.Set(x, y, c)
				} else {
					img.Set(y, x, c)
				}
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

		buf.Reset()
		if _, err := f.Seek(0, 0); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
