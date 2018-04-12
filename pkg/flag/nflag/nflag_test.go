package nflag

import (
	"flag"
	"testing"
)

func TestNFlag(t *testing.T) {
	fl := flag.NewFlagSet(t.Name(), flag.ContinueOnError)
	_ = fl.Bool("version", false, "bool var")
	_ = fl.Bool("v", false, "bool var")
	t.Logf("nflag: %d", fl.NFlag())
	fl.Parse([]string{"-version"})
	t.Logf("nflag: %d", fl.NFlag())
}
