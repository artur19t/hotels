package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscan(in, &num)
	writer(num, out)
}

func writer(num int, out *bufio.Writer) {
	if num > 10 {
		suf := num % 100
		if suf > 10 && suf < 20 {
			fmt.Fprintln(out, num, "компьютеров")
			return
		}
	}
	suf := num % 10
	if suf == 1 {
		fmt.Fprintln(out, num, "компьютер")
		return
	}
	if suf > 1 && suf < 5 {
		fmt.Fprintln(out, num, "компьютера")
		return
	}
	if suf == 0 || (suf > 4) {
		fmt.Fprintln(out, num, "компьютеров")
		return
	}
}
