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

	var min, max int
	fmt.Fscan(in, &min, &max)
	writer(min, max, out)
}

func writer(min, max int, out *bufio.Writer) {
	res := []int{}
	for i := min; i <= max; i++ {
		ok := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, i)
		}

	}
	fmt.Fprint(out, res)
}
