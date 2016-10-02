package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	// default filename
	readme = "./README.md"
	tree = "./tree.txt"
	bufsize = 16

	beginBlock = "```txt:./tree.txt"
	endBlock = "```"
)

type buffer struct {
	readme, tree []string
	block block
}
func (b *buffer) getBegin() int { return b.block.beginLine }
func (b *buffer) getEnd() int { return b.block.endLine }

type block struct {
	in, exit  bool
	beginLine, endLine int
}

// block開始位置の判定
func (b *block) setBegin(s string, line int) bool {
	if b.in { return false }
	if s == beginBlock {
		b.in = true
		b.beginLine = line
		return true
	}
	return false
}
// block終了位置の判定
func (b *block) setEnd(s string, line int) bool {
	if b.exit { return false }
	if s == endBlock && b.in {
		b.endLine = line
		b.exit = true
		return true
	}
	return false
}
//func (b *block) isIn() bool { return b.findFlag }





func fileRead() {

	var bl block

	fmt.Println("file read")

	// bufioを使ってファイルを掴む
	file, err := os.Open(readme)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error %q\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// bufio.Scanner
	fmt.Println("use bufio.Scanner")
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if errBufio := sc.Err(); errBufio != nil {
			fmt.Fprintf(os.Stderr, "sc.Scan error %q\n", errBufio)
			return
		}
		if bl.setBegin(sc.Text(), i) {
			fmt.Println("key true")
		}
		if bl.setEnd(sc.Text(), i) {
			fmt.Println("key end")
		}
		fmt.Printf("%4d行目: %s\n", i, sc.Text())
		fmt.Printf("%4d行目: %s\n", i, sc.Text())
	}
	fmt.Println("file scanner exit")



	// 同じファイルを同時に掴めるのか実験
	file2, err := os.Open(readme)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file2 open error %q\n", err)
		os.Exit(1)
	}
	defer file2.Close()

	// io.Reader
	fmt.Println("use io.Reader")
	buf := make([]byte, bufsize)
	i := 0
	for ; ;i++ {
		n, err := file2.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "file.Read error, %q\n", err)
			return
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
	}
	fmt.Println("buffer size", bufsize)
	fmt.Println("loop count", i)
	fmt.Println("io.Reader exit")


	fmt.Println("fileRead exit")
}

// DONE:
func addtree() {

	buf := new(buffer)

	// file reamde
	file1, err := os.Open(readme)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error %q\n", err)
		os.Exit(1)
	}
	defer file1.Close()

	// Scanner readme
	fmt.Println("Scanning README.md")
	for sc, i := bufio.NewScanner(file1), 0; sc.Scan(); i++ {
		if errsc := sc.Err(); errsc != nil {
			fmt.Fprintf(os.Stderr, "sc.Scan error %q\n", errsc)
			return
		}
		buf.block.setBegin(sc.Text(), i)
		buf.block.setEnd(sc.Text(), i)

		buf.readme = append(buf.readme, fmt.Sprintln(sc.Text()))
	}

	// file tree
	f2, err := os.Open(tree)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open error %q\n", err)
		return
	}
	defer f2.Close()

	// scanner tree
	fmt.Println("Scanning tree.txt")
	for sc := bufio.NewScanner(f2) ; sc.Scan(); {
		if errsc := sc.Err(); errsc != nil {
			fmt.Fprintf(os.Stderr, "sc.Scan() error %q\n", errsc)
			return
		}
		buf.tree = append(buf.tree, fmt.Sprintln(sc.Text()))
	}

	var str string
	for _, s := range buf.readme[:buf.getBegin()] { str += s }
	str += fmt.Sprintf("\n%s\n\n", beginBlock)
	for _, s := range buf.tree { str += s }
	str += fmt.Sprintf("\n%s\n\n", endBlock)
	for _, s := range buf.readme[buf.getEnd()+1:] { str += s }


	fmt.Println("readme overwrite")
	file3, err := os.Create(readme)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file3.Close()

	w := bufio.NewWriter(file3)
	n, err := w.WriteString(str)
	if err != nil {
		fmt.Fprintln(os.Stderr, "write bytes ", n, err)
		return
	}
	if err := w.Flush(); err != nil {
		fmt.Fprintln(os.Stderr, "Flush error ", err)
		return
	}
	fmt.Println("Write")


}

func main(){
	fileRead()

	fmt.Println()
	addtree()



	fmt.Println("main exit")
}
