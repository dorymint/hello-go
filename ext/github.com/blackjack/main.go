package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/blackjack/webcam"
)

func run() error {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		return err
	}
	defer cam.Close()

	err = cam.StartStreaming()
	if err != nil {
		return err
	}

	tempDir, err := ioutil.TempDir("", "stream")
	if err != nil {
		return err
	}

	var frameLimit = 30
	var timeout uint32 = 60
	for i := 0; i < frameLimit; i++ {
		err = cam.WaitForFrame(timeout)
		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			fmt.Fprintln(os.Stderr, err)
			continue
		default:
			return err
		}

		frame, err := cam.ReadFrame()
		if len(frame) != 0 {
			// for frame
			filename := filepath.Join(tempDir, strconv.Itoa(i))
			if err := ioutil.WriteFile(filename, frame, 0600); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	fmt.Println("writed in " + tempDir)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
