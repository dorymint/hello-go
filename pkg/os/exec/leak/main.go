package main

import (
	"os"
	"os/exec"
	"time"
)

// leak example
// need kill PID after run
// recomend open another terminal
// run command `kill PID` or `top` or `htop`

// for sh -c
const cmdline = `
x=1
while true; do
  printf "%s\n" "$x loops PLEAS \"kill $$\" on another terminal"
  x=$((x + 1))
  sleep 1
  [ $x -lt 100 ] || break
done
printf "stopped"
`

func main() {
	cmd := exec.Command("sh", "-c", cmdline)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
}
