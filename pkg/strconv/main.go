package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func split(str string) {
	fmt.Println("----------", str, "----------")
}

func hello() {
	str := "128"
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v: %d\n", reflect.TypeOf(i), i*2)
}

func overflow() {
	ul := []uint{50, 2, 1, 9, 1<<64-1}
	fmt.Println(ul)
	for _, u := range ul {
		s := strconv.FormatUint(uint64(u), 10)
		fmt.Printf("%s\t",s)
		fmt.Printf("%s\n",strconv.Itoa(int(u)))
	}
}

func main() {
	split("hello")
	hello()
	split("overflow")
	overflow()
}
