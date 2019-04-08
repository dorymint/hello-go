// walker.
//
//	go run ./main [Path]...
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type Walker struct {
	fileQueue chan string
	dirQueue  chan []string

	checked map[string]bool

	mu sync.Mutex
	wg sync.WaitGroup
}

func NewWalker() *Walker {
	return &Walker{checked: make(map[string]bool)}
}

func (w *Walker) SendPath(path ...string) error {
	var dirs []string
	for _, p := range path {
		abs, err := filepath.Abs(p)
		if err != nil {
			return err
		}
		fi, err := os.Stat(abs)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			dirs = append(dirs, abs)
		} else if fi.Mode().IsRegular() {
			w.wg.Add(1)
			w.fileQueue <- abs
		}
	}
	if len(dirs) != 0 {
		w.wg.Add(1)
		w.dirQueue <- dirs
	}
	return nil
}

func (w *Walker) errorHandle(errQueue <-chan error, handler func(error)) {
	for err := range errQueue {
		if err != nil {
			handler(err)
		}
	}
}

func (w *Walker) dirWalker(done <-chan struct{}, errQueue chan<- error) {
	var nextDirs []string
	for ; ; w.wg.Done() {
		select {
		case <-done:
			return
		case dirs := <-w.dirQueue:
			for _, dir := range dirs {
				w.mu.Lock()
				if w.checked[dir] {
					w.mu.Unlock()
					continue
				}
				w.checked[dir] = true
				w.mu.Unlock()

				fis, err := ioutil.ReadDir(dir)
				errQueue <- err
				for _, fi := range fis {
					if fi.IsDir() {
						nextDirs = append(nextDirs, filepath.Join(dir, fi.Name()))
					} else if fi.Mode().IsRegular() {
						w.wg.Add(1)
						w.fileQueue <- filepath.Join(dir, fi.Name())
					}
				}
			}
			if len(nextDirs) != 0 {
				w.wg.Add(1)
				w.dirQueue <- nextDirs
				nextDirs = nextDirs[:0]
			}
		}
	}
}

func (w *Walker) fileWalker(done <-chan struct{}, do func(string) error, errQueue chan<- error) {
	for ; ; w.wg.Done() {
		select {
		case <-done:
			return
		case file := <-w.fileQueue:
			w.mu.Lock()
			if w.checked[file] {
				w.mu.Unlock()
				continue
			}
			w.checked[file] = true
			w.mu.Unlock()

			errQueue <- do(file)
		}
	}
}

// do something for files.
// do func(file string) error
//
// error handler for dirWalker, fileWalker and Do.
// if unexpected error coming then to panic is better.
// errorHandler func(error)
func (w *Walker) Start(do func(string) error, errorHandler func(error)) (wait func()) {
	nworker := runtime.NumCPU() / 4
	if nworker < 2 {
		nworker = 2
	}
	var nfileQueue = 128

	errQueue := make(chan error, nfileQueue)
	go w.errorHandle(errQueue, errorHandler)

	done := make(chan struct{})
	w.dirQueue = make(chan []string, nworker)
	w.fileQueue = make(chan string, nfileQueue)
	for i := 0; i != nworker; i++ {
		go w.dirWalker(done, errQueue)
		go w.fileWalker(done, do, errQueue)
	}

	return func() {
		w.wg.Wait()
		close(done)
		close(errQueue)
	}
}

func main() {
	flag.Parse()
	w := NewWalker()

	var rwm sync.RWMutex
	do := func(file string) error {
		rwm.Lock()
		fmt.Println(file)
		rwm.Unlock()
		return nil
	}

	exitcode := 0
	var once sync.Once
	errorHandler := func(err error) {
		once.Do(func() { exitcode = 1 })
		if os.IsNotExist(err) || os.IsPermission(err) {
			rwm.Lock()
			fmt.Fprintln(os.Stderr, err)
			rwm.Unlock()
			return
		}
		// unexpected error.
		panic(err)
	}

	wait := w.Start(do, errorHandler)

	err := w.SendPath(flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	wait()
	os.Exit(exitcode)
}
