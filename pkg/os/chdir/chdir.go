package main

import (
	"log"
	"os"
	"sync"
)

func chdir(i int, wg *sync.WaitGroup) {
	log.Println("call chfir", i)
	defer wg.Done()

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(i, err)
		return
	}
	log.Println(i, pwd)

	if err := os.Chdir("../"); err != nil {
		log.Println(i, err)
		return
	}

	pwd, err = os.Getwd()
	if err != nil {
		log.Println(i, err)
		return
	}
	log.Println(i, pwd)
}

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		log.Println(i, "go chdir")
		wg.Add(1)
		// not safe
		go chdir(i, wg)
	}
	wg.Wait()
}
