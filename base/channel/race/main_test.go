package main

// go test -v -race

import (
	"fmt"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			ch <- fmt.Sprintf("go func [1]: i=%d", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			ch <- fmt.Sprintf("go func [2]: i=%d", i)
		}
	}()

	go func() {
		for {
			t.Logf("%s\n", <-ch)
			wg.Done()
		}
	}()

	wg.Wait()
}

func TestClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("go func [1]: i=%d", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("go func [2]: i=%d", i)
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := 0; true; i++ {
		s, ok := <-ch
		if !ok {
			t.Logf("total %d loops\n", i)
			break
		}
		t.Logf("%s\n", s)
	}
}

func TestRace(t *testing.T) {
	ch := make(chan string)
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		ch <- "race"
	}()

	go func() {
		<-ch
		wg.Done()
	}()

	wg.Wait()
}
