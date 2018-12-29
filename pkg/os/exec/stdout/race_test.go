// go test -v -race
package main

import (
	"os"
	"os/exec"
	"sync"
	"testing"
)

func TestNotRace(t *testing.T) {
	cl := []string{"sh", "-c", `for x in $(seq 3); do echo "$x loops"; sleep 1; done`}
	wg := new(sync.WaitGroup)
	run := func() {
		defer wg.Done()
		cmd := exec.Command(cl[0], cl[1:]...)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			t.Error(err)
		}
	}
	worker := 4
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go run()
	}
	wg.Wait()
}
