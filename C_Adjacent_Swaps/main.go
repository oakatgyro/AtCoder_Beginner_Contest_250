package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)
	Xs := make([]string, Q)
	for i := 0; i < Q; i++ {
		fmt.Scan(&Xs[i])
	}

	var res []string
	positions := map[string]int{}
	for i := 1; i <= N; i++ {
		num := strconv.Itoa(i)
		res = append(res, num)
		positions[num] = i - 1
	}

	for _, X := range Xs {
		if pos, ok := positions[X]; ok {
			var nextXPos, nextYPos int
			var Y string
			if pos == len(res)-1 {
				nextXPos = pos
				nextYPos = pos - 1
				Y = res[pos]
			} else {
				nextXPos = pos + 1
				nextYPos = pos
				Y = res[pos+1]
			}
			res[nextXPos], res[nextYPos] = res[nextYPos], res[nextXPos]
			positions[X] = nextXPos
			positions[Y] = nextYPos
		}
	}

	fmt.Println(strings.Join(res, " "))
}
