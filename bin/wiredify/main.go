package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eniehack/wiredize/pkg/queue"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		r := []rune(s)
		q := queue.New([]rune{r[0], r[1]})
		for i := 0; i < len(r); i++ {
			if string(q.Seek()) == "ヴ" {
				arr := []rune{q.Dequeue(), q.Dequeue()}
				switch string(arr) {
				case "ヴァ":
					fmt.Print("バ")
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴィ":
					fmt.Print("ビ")
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴェ":
					fmt.Print("ベ")
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				case "ヴォ":
					fmt.Print("ボ")
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
					if i+3 < len(r) {
						q.Enqueue(r[i+3])
					}
					i += 1
				default:
					fmt.Print("ブ")
					q.Enqueue(arr[1])
					if i+2 < len(r) {
						q.Enqueue(r[i+2])
					}
				}
			} else {
				fmt.Printf("%c", q.Dequeue())

				if i+2 < len(r) {
					q.Enqueue(r[i+2])
				}
			}
		}
	}
	fmt.Printf("\n")
}
