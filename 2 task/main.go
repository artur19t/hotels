package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nums []int
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	numL := strings.Fields(line)
	for _, n := range numL {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	fmt.Fscan(in, &nums)
	writer(nums, out)
}

func writer(num []int, out *bufio.Writer) {
	delimeters := make(map[int]int)
	for i := 2; i <= num[0]; i++ {
		if num[0]%i == 0 {
			delimeters[i] = 1
		}
	}
	for i := 1; i < len(num); i++ {
		for j := 2; j <= num[i]; j++ {
			if num[i]%j == 0 {
				val := delimeters[j]
				if val == i {
					delimeters[j]++
				}
			}
		}
	}
	result := []int{}
	for key, val := range delimeters {
		if val == len(num) {
			result = append(result, key)
		}
	}
	sort.Ints(result)
	fmt.Fprintln(out, result)
}
