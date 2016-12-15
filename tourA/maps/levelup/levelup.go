package main

import (
	"fmt"
)

func map1() {
	fmt.Println("\nmap1")

	mNew := new(map[string]string)
	mMake1 := make(map[string]string)
	mMake2 := make(map[string]string)

	// !!
	*mNew = mMake1 // Throw the address (*mMake1) ?
	mMake1 = mMake2
	(*mNew)["hello"] = "world"

	fmt.Println("new", mNew) // mNew have a address to instance?
	fmt.Println("make1", mMake1)
	fmt.Println("make2", mMake2)
}
func map2() {
	fmt.Println("\nmap2")

	mNew := new(map[string]string)
	mMake1 := make(map[string]string)
	mMake2 := make(map[string]string)

	// !!
	mNew = &mMake1 // Throw address not instance?
	mMake1 = mMake2
	(*mNew)["hello"] = "world" // mMake1["hello"] = "world"

	fmt.Println("new", mNew)
	fmt.Println("make1", mMake1)
	fmt.Println("make2", mMake2)
}

func map3() {
	fmt.Println("\nmap3")
	mNew := new(map[string]string)
	mMake1 := make(map[string]string)
	mMake2 := make(map[string]string)
	*mNew = mMake1 // (*mMake1)を*mNewが受けてる

	(*mNew)["test"] = "test"
	fmt.Println("mMake1", mMake1)

	// この代入で(*mMake1)は(*mNew)だけが参照を残す
	mMake1 = mMake2 // mMake1は単に(*mMake2)をもらって代入してるだけっぽい
	(*mNew)["hello"] = "world"

	fmt.Println("new", mNew)
	fmt.Println("make1", mMake1)
	fmt.Println("make2", mMake2)
}
func map4() {
	fmt.Println("\nmap4")
	mNew := new(map[string]string)
	mMake1 := make(map[string]string)
	mMake2 := make(map[string]string)
	mNew = &mMake1 // mNewがmMake1へのアドレスを持ってる

	(*mNew)["test"] = "test" // mMake["test"] = "test"
	fmt.Println("mMake1", mMake1)

	// この代入で(*mMake1)は参照を失う
	mMake1 = mMake2 // mMake1は単に(*mMake2)をもらって代入してるだけ
	(*mNew)["hello"] = "world"

	fmt.Println("new", mNew)
	fmt.Println("make1", mMake1)
	fmt.Println("make2", mMake2)
}

func map5() {
	fmt.Println("\nmap5")

	mMake := make(map[string]string)
	mMake["hello"] = "world"
	//fmt.Printf("%p\n", *mMake) // invalid indirect
	fmt.Printf("origin\n")
	fmt.Printf("%p mMake\n", mMake)
	fmt.Printf("%v %p\n", mMake, &mMake)

	mNew1 := new(map[string]string)
	mNew1 = &mMake
	fmt.Printf("%v %p %p\n", *mNew1, mNew1, &mNew1)

	mNew2 := new(*map[string]string)
	mNew2 = &mNew1
	fmt.Printf("%v %p %p %p\n", **mNew2, *mNew2, mNew2, &mNew2)

	mNew3 := new(**map[string]string)
	mNew3 = &mNew2
	fmt.Printf("%v %p %p %p %p\n", ***mNew3, **mNew3, *mNew3, mNew3, &mNew3)

	mNew4 := new(***map[string]string)
	mNew4 = &mNew3
	fmt.Printf("%v %p %p %p %p %p\n", ****mNew4, ***mNew4, **mNew4, *mNew4, mNew4, &mNew4)
	fmt.Printf("%p  ****mNew4\n", ****mNew4)
}

func main() {
	map1()
	map2()
	map3()
	map4()
	map5()
}
