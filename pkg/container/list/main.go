// cantainer list.
// TODO: fix
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println("l:", l)

	// duplication
	var exl list.List
	exl = *l
	fmt.Println("exl:", exl)
	fmt.Printf("%p\n%p\n%p\n%+v\n%+v\n", &l, &exl, l, exl, *l)
	// &exl is point to entity

	// l
	e := l.PushFront("hi")
	fmt.Println(e.Value)
	e = l.PushFront("list!")
	fmt.Println(e.Value, e.Next().Value)
	fmt.Println("l.Len:", l.Len(), "exl.Len:", exl.Len())

	// exl
	exe := exl.PushBack("hello world")
	// maybe incremental exl.Len, don't incremental l.Len
	fmt.Println("exl.Len:", exl.Len(), exe.Value, exl.Front().Value)
	// exl.Front().Value == <nil>
	e = l.Front() // not exl
	fmt.Println(e.Value, e.Next().Value, e.Next().Next().Value)
	// I was bit surprised but understanding
	// e are simple node {*Next, *Prev, Valu}
	// l is having: some methods for control of simple node
	//            : point to Front and Back
	//            : value of length(virtual length?)
	fmt.Println("l.Len:", l.Len(), "exl.Len:", exl.Len())

	// remove
	// if exl.Remove(e) then runtime panic
	// borrowing the methods of exl
	fmt.Println(exl.Remove(e.Next())) // e is slave of l
	fmt.Println(exl.Remove(e.Next().Next()))
	fmt.Println("l.Len:", l.Len(), "exl.Len:", exl.Len())
	fmt.Println(exl, l)

	// query
	each := func(e *list.Element) {
		for ; e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
	fmt.Printf("%+v\n%+v\n", *l, exl)
	each(l.Front())
	each(exl.Front())
	each(e)
}
