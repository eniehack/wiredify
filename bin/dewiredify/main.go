package main

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/eniehack/wiredize/pkg/queue"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := new(bytes.Buffer)
	for scanner.Scan() {
		s := scanner.Text()
		r := []rune(s)
		q := queue.New([]rune{r[0], r[1]})
		for i := 0; i < len(r); i++ {
			if q.Seek() == 'ヴ' {
				arr := []rune{q.Dequeue(), q.Dequeue()}
				switch string(arr) {
				case "ヴァ":
					buf.WriteRune('バ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴィ":
					buf.WriteRune('ビ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴェ":
					buf.WriteRune('ベ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴォ":
					buf.WriteRune('ボ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				default:
					buf.WriteRune('ブ')
					q.Enqueue(arr[1])
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
				}
			} else {
				buf.WriteRune(q.Dequeue())

				if i+2 < len(r) {
					q.Enqueue(r[i+2])
				}
			}
		}
	}
	buf.WriteRune('\n')

	io.Copy(os.Stdout, buf)
}
