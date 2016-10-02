package fizzbuzz

import (
	"fmt"
	"time"
)

var fizzBuzzError = string(":negativ number and 0 is invalid")


// FizzBuzz fizzbuzz! count from to n
func FizzBuzz(n int) error {
	if n <= 0 { return fmt.Errorf("FizzBuzz()%s:%d", fizzBuzzError, n) }
	for i := 1; i <= n; i++ {
		time.Sleep(time.Millisecond * 100)
		str, err := SfizzBuzz(i)
		if err != nil { return err }
		fmt.Println(str)
	}
	return nil
}

// SfizzBuzz return single string
func SfizzBuzz(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("SfizzBuzz()%s:%d", fizzBuzzError, n)
	}
	if n%15 == 0 { return fmt.Sprint("fizzbuzz"), nil }
	if n%3 == 0 { return fmt.Sprint("fizz"), nil }
	if n%5 == 0 { return fmt.Sprint("buzz"), nil }
	return fmt.Sprintf("%d", n), nil
}


// ShowFizzBuzz count from to n
func ShowFizzBuzz(n int) error {
	if n <= 0 { return fmt.Errorf("FizzBuzz()%s:%d",fizzBuzzError, n) }
	for i := 1; i <= n; i++ {
		if i % 15 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
	return nil
}
