package t

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	/// init

	/// run
	exitCode := m.Run()

	/// clean up
	os.Exit(exitCode)
}

func TestSimple(t *testing.T) {
	got := 1
	want := 2
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
