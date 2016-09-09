package fizzbuzz

import (
	"log"
	"testing"
)

func ExampleFizzBuzz() {
	err := FizzBuzz(15)
	if err != nil {
		log.Fatalf("FizzBuzz(10):expected err==nil...out %q", err)
	}
	// Output:
	// 1
	// 2
	// fizz
	// 4
	// buzz
	// fizz
	// 7
	// 8
	// fizz
	// buzz
	// 11
	// fizz
	// 13
	// 14
	// fizzbuzz

	err = FizzBuzz(0)
	if err == nil {
		log.Fatalf("FizzBuzz(0):expected err!=nil...out %q", err)
	}
	err = FizzBuzz(-10)
	if err == nil {
		log.Fatalf("FizzBuzz(-10):expected err!=nil...out %q", err)
	}
}

func TestSfizzBuzz(t *testing.T) {
	_, err := SfizzBuzz(10)
	if err != nil {
		t.Fatalf("SfizzzBuzz(10):expected err==nil...out %q", err)
	}
	_, err = SfizzBuzz(0)
	if err == nil {
		t.Fatalf("SfizzzBuzz(0):expected err!=nil...out %q", err)
	}
	_, err = SfizzBuzz(-10)
	if err == nil {
		t.Fatalf("SfizzBuzz(-10):expected err!=nil...out %q", err)
	}
	t.Logf("SfizzBuzz:PATH error return check")

	testData := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1"},
		{input: 2, expected: "2"},
		{input: 3, expected: "fizz"},
		{input: 4, expected: "4"},
		{input: 5, expected: "buzz"},
		{input: 6, expected: "fizz"},
		{input: 7, expected: "7"},
		{input: 8, expected: "8"},
		{input: 9, expected: "fizz"},
		{input: 10, expected: "buzz"},
		{input: 11, expected: "11"},
		{input: 12, expected: "fizz"},
		{input: 13, expected: "13"},
		{input: 14, expected: "14"},
		{input: 15, expected: "fizzbuzz"},
	}
	for _, x := range testData {
		out, err := SfizzBuzz(x.input)
		if err != nil { t.Fatalf("Sfizzbuzz:%q", err) }
		if out != x.expected {
			t.Errorf("x.expected:%q ... out:%q", x.expected, out)
		}
	}
}

func ExampleShowFizzBuzz() {
	if err := ShowFizzBuzz(15); err != nil {
		log.Fatalf("ShowFizzBuzz:%q", err)
	}
	// Output:
	// 1
	// 2
	// fizz
	// 4
	// buzz
	// fizz
	// 7
	// 8
	// fizz
	// buzz
	// 11
	// fizz
	// 13
	// 14
	// fizzbuzz
}

func TestSimple(t *testing.T) {
	value := 1
	expected := 2
	if value != expected {
		//t.Fatalf("Expected %v, but %v:", expected, value)
	}
}
