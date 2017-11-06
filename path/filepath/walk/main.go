package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	root  = flag.String("root", "", "")
	match = flag.String("match", "", "")
	dirs  = flag.Bool("dirs", false, "")
)

func testErr(root string) {
	fmt.Println("testErr: always return error")

	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("[%s]", path)
		fmt.Println()
		return errors.New("always return error")
	}

	if err := filepath.Walk(root, f); err != nil {
		fmt.Println("walk err:", err)
	}
}

func matchDirectory(root string, match string) {
	fmt.Println("matchDirectory:")

	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.Contains(path, match) {
			return nil
		}
		fmt.Printf("[%s]", path)
		fmt.Println()
		return nil
	}

	if err := filepath.Walk(root, f); err != nil {
		fmt.Println("walk err:", err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("invalid args:", flag.Args())
	}
	var err error
	*root, err = filepath.Abs(*root)
	if err != nil {
		log.Fatal(err)
	}

	printPath := func() func(string, os.FileInfo) {
		if *dirs {
			return func(path string, info os.FileInfo) {
				if info != nil && info.IsDir() {
					fmt.Printf("[%s]\n", path)
				}
			}
		}
		return func(path string, info os.FileInfo) {
			fmt.Printf("[%s]\n", path)
		}
	}()

	f := func() func(string, os.FileInfo, error) error {
		switch *match {
		case "":
			return func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("err:%v", err)
				}
				printPath(path, info)
				return nil
			}
		default:
			return func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("err:%v", err)
				}
				if !strings.Contains(path, *match) {
					return nil
				}
				printPath(path, info)
				return nil
			}
		}
	}()

	err = filepath.Walk(*root, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- testErr ---")
	testErr(*root)
	fmt.Println("--- matchDirectory ---")
	matchDirectory(*root, *match)
}
