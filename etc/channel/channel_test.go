package c

import (
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	wg := new(sync.WaitGroup)
	queue := make(chan rune, 128)
	worker := func(queue <-chan rune, resCh chan<- string) {
		t.Log("on worker")
		for {
			str := <-queue
			// do something
			resCh <- string(str) + " [appended]"
		}
	}
	general := func(queue chan rune, wg *sync.WaitGroup) {
		resch := make(chan string)
		for i := 0; i < 4; i++ {
			go worker(queue, resch)
		}
		t.Log("on general")
		for {
			str := <-resch
			// do something
			t.Log(str)
			wg.Done()
		}
	}

	tests := "hellow world"
	go general(queue, wg)
	for _, s := range tests {
		wg.Add(1)
		queue <- s
	}
	t.Log("wait")
	wg.Wait()
}
