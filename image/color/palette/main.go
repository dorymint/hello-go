package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/png"
	"log"
	"os"
)

func main() {
	out := flag.String("out", "img.png", "specify output file path")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("invalid argument:", flag.Args())
	}
	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	length := len(palette.WebSafe)
	img := image.NewNRGBA(image.Rect(0, 0, length, length))
	for y := 0; y < length; y++ {
		c := palette.WebSafe[y]
		for x := 0; x < length; x++ {
			img.Set(x, y, c)
		}
	}
	var b []byte
	buf := bytes.NewBuffer(b)
	if err := png.Encode(buf, img); err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, buf)
}
