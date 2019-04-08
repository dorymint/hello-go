// queue.
package main

import "fmt"

type queue struct {
	vs []interface{}
}

func (q *queue) Push(v interface{}) {
	q.vs = append(q.vs, v)
}

func (q *queue) Pop() interface{} {
	// to do fix for out of bound.
	v := q.vs[0]
	q.vs = q.vs[1:]
	return v
}

func (q *queue) Len() int {
	return len(q.vs)
}

func main() {
	q := new(queue)
	q.Push("hello")
	q.Push("world")
	q.Push(0)
	q.Push(1)

	for q.Len() != 0 {
		v := q.Pop()
		fmt.Printf("%v\n", v)
	}
}
