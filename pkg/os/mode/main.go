package main

import (
	"fmt"
	"os"
)

func main() {
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
	fmt.Println(os.ModeDir|os.ModePerm, "ModeDir|ModePerm: ")
}
