package main

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/eniehack/wiredize/pkg/wiredify"
)

func main() {
	h := &wiredify.Handler{
		In:  bufio.NewScanner(os.Stdin),
		Out: new(bytes.Buffer),
	}

	h.Wiredify()

	h.Out.WriteRune('\n')

	io.Copy(os.Stdout, h.Out)
}
