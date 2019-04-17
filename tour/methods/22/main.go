// Exercise: Readers.
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	if len(b) < 1 {
		return 0, nil
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
