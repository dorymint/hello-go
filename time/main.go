package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println(n)
	fmt.Println("day:", n.Day(), "hour:", n.Hour(), "minute:", n.Minute(), "month:", n.Month(), "year:", n.Year(), "yearday:", n.YearDay())
	fmt.Println(n.Date())
}
