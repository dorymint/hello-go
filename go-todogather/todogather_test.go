// TODO:
// TEST for TODOGather!!
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Create tempfile, Return filename.
func tempFile(content string) (string, error) {
	f, err := ioutil.TempFile("", "prefix")
	if err != nil {
		return "", err
	}
	defer func() {
		if errclose := f.Close(); errclose != nil {
			fmt.Fprintf(os.Stderr, "tempFile:%q\n", errclose)
		}
	}()
	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

// define from main
// str is from function infomation.
//func fatalIF(str string, err error) {
//	if err != nil {
//		log.Fatalf("%s:%q\n", str, err)
//	}
//}

// TODO:
func TestFileTODOGather(t *testing.T) {
	tmp, err := tempFile(` test TODO: hello`)
	fatalIF("tempFIle:", err)

	out, err := fileInGather(tmp)
	fatalIF("fileInGather:", err)

	expected := []string{`L1:TODO: hello`}
	if len(out) != len(expected) {
		t.Fatalf("expected %q but %q\n", expected, out)
	}
	for i := range out {
		if out[i] != expected[i] {
			t.Fatalf("expected %q but %q\n", expected, out)
		}
	}

	// executable test
	test, err := fileInGather("./testdir/todogather")
	t.Fatalf("%q, %q\n", test, err)
}
