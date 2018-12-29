package fizzbuzz

import (
	"errors"
	"fmt"
	"time"
)

var ErrUnexpectedNumber = errors.New("unexpected number")

// FizzBuzz fizzbuzz! count from to n
func FizzBuzz(n int) error {
	if n <= 0 {
		return fmt.Errorf("FizzBuzz():%s:%d", ErrUnexpectedNumber, n)
	}
	for i := 1; i <= n; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(SfizzBuzz(i))
	}
	return nil
}

// SfizzBuzz return single string
func SfizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return fmt.Sprint("fizzbuzz")
	case n%3 == 0:
		return fmt.Sprint("fizz")
	case n%5 == 0:
		return fmt.Sprint("buzz")
	default:
		return fmt.Sprintf("%d", n)
	}
}

// ShowFizzBuzz count from to n
func ShowFizzBuzz(n int) error {
	if n <= 0 {
		return fmt.Errorf("FizzBuzz()%s:%d", ErrUnexpectedNumber, n)
	}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
	return nil
}
