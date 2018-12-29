package main

import (
	"context"
	"os"
	"os/exec"
	"sync"
	"time"
)

func daemon() (taskQ chan []string, wg *sync.WaitGroup) {
	taskQ = make(chan []string)
	wg = new(sync.WaitGroup)

	var ctx context.Context
	var cancel context.CancelFunc
	cancel = func() {}

	go func() {
		for {
			cl := <-taskQ
			if len(cl) == 0 {
				wg.Done()
				continue
			}

			// kill stil running command
			cancel()
			ctx, cancel = context.WithCancel(context.Background())

			cmd := exec.CommandContext(ctx, cl[0], cl[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				wg.Done()
				continue
			}

			go func() {
				cmd.Wait()
				wg.Done()
			}()
		}
	}()

	return taskQ, wg
}

func main() {
	taskQ, wg := daemon()
	wg.Add(1)
	taskQ <- []string{"sh", "-c", `sleep 1; echo "hello world"`}
	wg.Wait()

	wg.Add(1)
	taskQ <- []string{"sh", "-c", `for x in $(seq 3); do echo "$x loops"; sleep 1; done`}
	wg.Wait()

	wg.Add(2)
	taskQ <- []string{"sh", "-c", `echo "PID=$$ sleeping..."; sleep 100; echo "unreachable"`}
	time.Sleep(time.Second)
	taskQ <- []string{"sh", "-c", `echo "killed process!"`}
	wg.Wait()
}
