package main

import (
	"fmt"
	"log"
	"os"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func q1() {
	const testFilePath = "./t/test.txt"
	f, err := os.Create(testFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		os.Remove(f.Name())
		fmt.Println("deleted:", f.Name())
		f.Close()
	}()
	fmt.Println("created:", f.Name())

	// helloworld
	b := make([]byte, len("helloworld")+1)
	reprint := func() {
		if _, err := f.ReadAt(b, 0); err != nil {
			log.Println(err) // maybe: io.EOF
		}
		fmt.Println(string(b), "[]byte:", b)
	}
	f.WriteString("hello")
	reprint()
	f.WriteString("world")
	reprint()
}

// for q2
func create(files []string) (fs []*os.File, fsClose func()) {
	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		fs = append(fs, f)
	}
	fsClose = func() {
		for _, f := range fs {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}
	}
	return
}

func q2() {
	var testFiles = []string{
		"./t/q2-1.txt",
		"./t/q2-2.txt",
	}
	// os.Create: override exists files
	fs, fsClose := create(testFiles)
	defer fsClose()
	for _, f := range fs {
		fmt.Println(f.Name())
	}

	if err := os.Mkdir("./t/q2", os.ModePerm); err != nil {
		fmt.Fprintln(os.Stderr, "err:", err)
	}
	for _, f := range append(testFiles, "./invalid", "./t/q2") {
		info, err := os.Stat(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "err:", err)
			continue
		}
		fmt.Println(info.Name(), info.IsDir(), info.Mode().IsRegular())
	}
}

func q3() {
	fmt.Println("list: os.FileMode")
	showMode := func(msg string, mode os.FileMode) {
		fmt.Println(mode, msg)
	}
	mods := map[string]os.FileMode{
		"ModeAppend":     os.ModeAppend,
		"ModeCharDevice": os.ModeCharDevice,
		"ModeDevice":     os.ModeDevice,
		"ModeDir":        os.ModeDir,
		"ModeExclusive":  os.ModeExclusive,
		"ModeNamedPipe":  os.ModeNamedPipe,
		"ModePerm":       os.ModePerm,
		"ModeSetgid":     os.ModeSetgid,
		"ModeSetuid":     os.ModeSetuid,
		"ModeSocket":     os.ModeSocket,
		"ModeSticky":     os.ModeSticky,
		"ModeSymlink":    os.ModeSymlink,
		"ModeTemporary":  os.ModeTemporary,
		"ModeType":       os.ModeType,
	}
	for key, m := range mods {
		showMode(key, m)
	}
	fmt.Println("ModeDir|ModePerm: ", os.ModeDir|os.ModePerm)
}

func main() {
	log.SetFlags(log.Lshortfile)
	split("q1 os.Create")
	q1()
	split("q2")
	q2()
	split("q3 os.FileMode")
	q3()
}
