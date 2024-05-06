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
	fmt.Fprint(out, "  ")
	for i := 1; i <= num; i++ {
		fmt.Fprint(out,  i, " ")
	}
	fmt.Fprintln(out)
	for i := 1; i <= num; i++ {
		fmt.Fprint(out, i, " ")
		for j := 1; j <= num; j++ {
			fmt.Fprint(out, i*j, " ")
		}
		fmt.Fprintln(out)
	}
}
