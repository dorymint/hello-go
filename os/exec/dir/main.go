package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

type command struct {
	dir string
	wg  *sync.WaitGroup
}

func (c *command) run(w io.Writer) {
	defer c.wg.Done()
	// on linux
	cmd := exec.Command("sh", "-c", "pwd")
	//cmd := exec.Command("sh", "-c", "pwd; ls")
	cmd.Stdout = w
	cmd.Stderr = w
	cmd.Stdin = nil
	cmd.Dir = c.dir
	if err := cmd.Run(); err != nil {
		log.Println(c.dir, err)
	}
}

func main() {
	dirList := []string{
		"./",
		"none",
		"/home",
	}

	wg := new(sync.WaitGroup)
	for _, s := range dirList {
		c := &command{dir: s, wg: wg}
		wg.Add(1)
		c.run(os.Stdout)
	}
	wg.Wait()
}
