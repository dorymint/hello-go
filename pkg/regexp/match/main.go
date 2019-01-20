package main

import (
	"fmt"
	"regexp"
)

func main() {
	{
		reg := "^[a-z]+"
		match := regexp.MustCompile(reg)
		str := "hello world"
		loc := match.FindStringIndex(str)
		if loc != nil {
			fmt.Println("loc:", loc, str[loc[0]:loc[1]])
		}
	}

	{
		reg := "[0-9]+"
		match := regexp.MustCompile(reg)
		str := "test 123456 test" + "\n" + "test" + "\n" + "7890 test"
		loc := match.FindStringIndex(str)
		fmt.Println("loc:", loc)
		fmt.Println("str[[0]:loc[1]]:", str[loc[0]:loc[1]])
	}

	{
		reg := "^hello-(?:world|lily)$"
		match := regexp.MustCompile(reg)
		tests := []struct {
			str     string
			ismatch bool
		}{
			{str: "hello world", ismatch: false},
			{str: "hello-world", ismatch: true},
			{str: "hello-lily", ismatch: true},
			{str: "hello-", ismatch: false},
			{str: "hello-wo", ismatch: false},
			{str: "hello", ismatch: false},
			{str: "world", ismatch: false},
		}
		fmt.Println("reg:", reg)
		for i, test := range tests {
			if b := match.MatchString(test.str); b != test.ismatch {
				fmt.Printf("[Error %d]: unexpected result: reg:%v str:%v ismatch:%v bool:%v\n", i, reg, test.str, test.ismatch, b)
			} else if b {
				fmt.Printf("[Log %d]: matched:%v\n", i, test.str)
			}
		}
	}
}
