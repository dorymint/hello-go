package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

// exec test

func main() {
	//execName()
	//execStart()
	execStream(os.Stdin, os.Stdout, os.Stderr)
}

func execName() {
	fmt.Println("test2")
	// exit absolute file path
	fmt.Println(exec.LookPath("ls"))
	exe := exec.Command("ls")

	exist, err := exe.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(exist))

	exe = exec.Command("ls", "../../")
	exist2, err := exe.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(exist2))
}

func execStart() {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish ...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v\n", err)
}

func execStream(src io.Reader, dst io.Writer, errDst io.Writer) error {
	cmd := exec.Command("tr", "a-z", "A-Z")
	stdin, err := cmd.StdinPipe()
	if err !=nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(3)
	// stdin
	go func() {
		stdin.Write([]byte("hello "))
//		_, err = io.Copy(stdin, src)
//		fmt.Println("copy to stdin")
//		if err != nil {
//			log.Fatal(err)
//		}
		stdin.Write([]byte("world\n"))
		stdin.Close()
		fmt.Println("close stdin")
		stdin.Write([]byte("hi!!!"))
		wg.Done()
	}()
	// stdout
	go func() {
		fmt.Println("temporary stoped for stdout")
		_, err = io.Copy(dst, stdout)
		fmt.Println("copied to os.stdout")
		if err != nil {
			log.Fatal(err)
		}
		stdout.Close()
		fmt.Println("close stdout")
		wg.Done()
	}()
	// stderr
	go func() {
		fmt.Println("temporary stoped for stderr")
		_, err = io.Copy(errDst, stderr)
		fmt.Println("copied to os.stderr")
		if err != nil {
			log.Fatal(err)
		}
		stderr.Close()
		fmt.Println("close stderr")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("wg done!")
	return cmd.Wait()
}
