package wiredify

import (
	"bufio"
	"bytes"

	"github.com/eniehack/wiredize/pkg/queue"
)

type Handler struct {
	In  *bufio.Scanner
	Out *bytes.Buffer
}

func (h *Handler) Wiredify() {
	for h.In.Scan() {
		s := h.In.Text()
		for _, c := range s {
			switch c {
			case 'バ':
				h.Out.WriteString("ヴァ")
			case 'ビ':
				h.Out.WriteString("ヴィ")
			case 'ブ':
				h.Out.WriteString("ヴ")
			case 'ベ':
				h.Out.WriteString("ヴェ")
			case 'ボ':
				h.Out.WriteString("ヴォ")
			default:
				h.Out.WriteRune(c)
			}
		}
	}
}

func (h *Handler) Dewiredify() {
	for h.In.Scan() {
		s := h.In.Text()
		r := []rune(s)
		q := queue.New([]rune{r[0], r[1]})
		for i := 0; i < len(r); i++ {
			if q.Seek() == 'ヴ' {
				arr := []rune{q.Dequeue(), q.Dequeue()}
				switch string(arr) {
				case "ヴァ":
					h.Out.WriteRune('バ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴィ":
					h.Out.WriteRune('ビ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴェ":
					h.Out.WriteRune('ベ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴォ":
					h.Out.WriteRune('ボ')
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				default:
					h.Out.WriteRune('ブ')
					q.Enqueue(arr[1])
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
				}
			} else {
				h.Out.WriteRune(q.Dequeue())

				if i+2 < len(r) {
					q.Enqueue(r[i+2])
				}
			}
		}
	}
}
