// use {{ .Method }}.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

const tmplText = `Usage:
  {{ .Name }} [Options]

Description:
  Short description

Options:
{{ .SDefaults }}
Examples:
{{ .SExamples }}
`

type Mycmd struct {
	name string

	fs *flag.FlagSet

	w    io.Writer
	errw io.Writer
}

func NewMycmd() *Mycmd {
	p := &Mycmd{
		w:    os.Stdout,
		errw: os.Stderr,
		name: "cmdname",
	}
	p.fs = flag.NewFlagSet(p.name, flag.ExitOnError)
	p.fs.SetOutput(p.errw)
	return p
}

func (p *Mycmd) Name() string { return p.name }

func (p *Mycmd) SDefaults() string {
	tmpw := p.fs.Output()
	defer p.fs.SetOutput(tmpw)
	buf := bytes.NewBufferString("")
	p.fs.SetOutput(buf)
	p.fs.PrintDefaults()
	return buf.String()
}

func (p *Mycmd) SExamples() string {
	examples := []struct {
		e string
		c string
	}{
		{
			e: p.name + " -help",
			c: "Display help message",
		},
		{
			e: p.name + " -version",
			c: "Print version information",
		},
	}

	var maxwc int
	for _, example := range examples {
		if n := len(example.e); n > maxwc {
			maxwc = n
		}
	}

	var msg string
	for _, example := range examples {
		// two spaces
		msg += fmt.Sprintf("  $ %s %s# %s\n",
			example.e,
			strings.Repeat(" ", maxwc-len(example.e)),
			example.c)
	}
	return msg
}

func main() {
	cmd := NewMycmd()
	_ = cmd.fs.Bool("help", false, "help message")
	_ = cmd.fs.Bool("version", false, "version information")
	_ = cmd.fs.String("file", "", "path to file")
	cmd.fs.Parse([]string{})

	t := template.Must(template.New("usage").Parse(tmplText))
	if err := t.Execute(os.Stdout, cmd); err != nil {
		panic(err)
	}
}
