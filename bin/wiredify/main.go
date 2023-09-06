package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := new(bytes.Buffer)
	for scanner.Scan() {
		s := scanner.Text()
		for _, c := range s {
			switch c {
			case 'バ':
				buf.WriteString("ヴァ")
			case 'ビ':
				buf.WriteString("ヴィ")
			case 'ブ':
				buf.WriteString("ヴ")
			case 'ベ':
				buf.WriteString("ヴェ")
			case 'ボ':
				buf.WriteString("ヴォ")
			default:
				buf.WriteRune(c)
			}
		}
	}
	buf.WriteRune('\n')

	io.Copy(os.Stdout, buf)
}
