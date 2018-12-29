package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// for sh -c
const cmdline = `x=0; while true; do echo "hello $x"; x=$(( x + 1 )); sleep 1; done`

func graceful() {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdline)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
	cancel()

	if err := cmd.Wait(); err != nil {
		fmt.Println("cmd.Wait:", err)
	}
}

func main() {
	graceful()
}
