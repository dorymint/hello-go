package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
	"time"
)

// exec test

func main() {
	//execName()
	//execStart()
	//execStream(os.Stdin, os.Stdout, os.Stderr)
	//execStdin()
	execSkeleton()
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
	if err != nil {
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
		if true {
			_, err = io.Copy(stdin, src)
			fmt.Println("copied to stdin")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			stdin.Write([]byte("hello "))
			stdin.Write([]byte("world\n"))
		}
		stdin.Close()
		fmt.Println("close stdin")
		stdin.Write([]byte("hi!!!"))
		wg.Done()
	}()
	// stdout
	go func() {
		if true {
			_, err = io.Copy(dst, stdout)
			fmt.Println("copied to os.stdout")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			buf := make([]byte, 100, 100)
			_, err := stdout.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("buf:", string(buf))
		}
		stdout.Close()
		fmt.Println("close stdout")
		wg.Done()
	}()
	// stderr
	go func() {
		if false {
			_, err = io.Copy(errDst, stderr)
			fmt.Println("copied to os.stderr")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			buferr := make([]byte, 10, 10)
			_, err := stderr.Read(buferr)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			fmt.Println("buferr:", string(buferr))
		}
		stderr.Close()
		fmt.Println("close stderr")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("wg done!")
	return cmd.Wait()
}

func execStdin() {
	cmd := exec.Command("wc")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(stdin, "hoge hoge")
	stdin.Close()
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func execSkeleton() {
	cmd := exec.Command("go", "run", "../skeleton.go")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for sc := bufio.NewScanner(stdout); sc.Scan(); {
			fmt.Println(sc.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	n, err := io.WriteString(stdin, "\nhappy\n")
	fmt.Println(n, err)
	io.WriteString(stdin, "next")
	io.WriteString(stdin, "\n")
	time.Sleep(time.Second * 4)

}
