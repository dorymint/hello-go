package main

import (
	"fmt"
)

func main() {
	const str = "hello world"
	sl := make([]byte, len(str))
	sl = []byte(str[:5])
	fmt.Printf("sl=%s len=%d\n\n", sl, len(sl))

	fmt.Printf("sl[:3]=%s\n", sl[:3])
	fmt.Printf("sl[3:]=%s\n\n", sl[3:])

	fmt.Printf("len(sl)/2=%d\n", len(sl)/2)
	fmt.Printf("sl[:len(sl)/2]=%s\n\n", sl[:len(sl)/2])

	fmt.Printf("%g\n", sl)
	fmt.Printf("%g\n", sl[0])
	fmt.Printf("%g\n\n", sl[:2])

	sl = append(sl, ':', 'p', 'p')
	sl = append(sl, "ended"...)
	fmt.Printf("sl=%s len=%d\n\n", sl, len(sl))
}
