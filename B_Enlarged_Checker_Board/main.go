package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)
	for i := 0; i < N; i++ {
		var start, end, line string
		if i%2 == 0 {
			start = "."
			end = "#"
		} else {
			start = "#"
			end = "."
		}
		for j := 0; j < N; j++ {
			if j%2 == 0 {
				line = line + strings.Repeat(start, B)
			} else {
				line = line + strings.Repeat(end, B)
			}
		}

		for j := 0; j < A; j++ {
			fmt.Println(line)
		}
	}

}
