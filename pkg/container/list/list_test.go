package tlist

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.New()
	t.Log("l:", l)

	// duplication
	var exl list.List
	t.Run("basic", func(t *testing.T) {
		exl = *l
		t.Log("exl:", exl)
		t.Logf("%p\n%p\n%p\n%q\n%q\n", &l, &exl, l, exl, *l)
	})
	// &exl is point to entity

	// l
	e := l.PushFront("hi")
	t.Run("l", func(t *testing.T) {
		t.Log(e.Value)
		e = l.PushFront("list!")
		t.Log(e.Value, e.Next().Value)
		t.Log("l.Len:", l.Len(), "exl.Len:", exl.Len())
	})

	// exl
	exe := exl.PushBack("hello world")
	t.Run("exl", func(t *testing.T) {
		// maybe incremental exl.Len, don't incremental l.Len
		t.Log("exl.Len:", exl.Len(), exe.Value, exl.Front().Value)
		// exl.Front().Value == <nil>
		e = l.Front() // not exl
		t.Log(e.Value, e.Next().Value, e.Next().Next().Value)
		// I was bit surprised but understanding
		// e are simple node {*Next, *Prev, Valu}
		// l is having: some methods for control of simple node
		//            : point to Front and Back
		//            : value of length(virtual length?)
		t.Log("l.Len:", l.Len(), "exl.Len:", exl.Len())
	})

	// remove
	// if exl.Remove(e) then runtime panic
	// borrowing the methods of exl
	t.Run("remove", func(t *testing.T) {
		t.Log(exl.Remove(e.Next())) // e is slave of l
		t.Log(exl.Remove(e.Next().Next()))
		t.Log("l.Len:", l.Len(), "exl.Len:", exl.Len())
		t.Log(exl, l)
	})

	// query
	t.Run("query", func(t *testing.T) {
		each := func(e *list.Element) {
			for ; e != nil; e = e.Next() {
				t.Log(e.Value)
			}
		}
		t.Logf("%q\n%q\n", *l, exl)
		each(l.Front())
		each(exl.Front())
		each(e)
	})
}
