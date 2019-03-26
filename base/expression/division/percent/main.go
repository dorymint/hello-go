// percent.
package main

import "fmt"

func main() {
	is := []int{10, 1, 0, 133}
	for _, i := range is {
		fmt.Printf("i=%d %%=%d /=%d\n", i, i%10, i/10)
	}
}
