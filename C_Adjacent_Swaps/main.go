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
	for i := 1; i <= N; i++ {
		res = append(res, strconv.Itoa(i))
	}

	for _, X := range Xs {
		for i, r := range res {
			if r == X {
				if i == len(res)-1 {
					res[i], res[i-1] = res[i-1], res[i]
					break
				} else {
					res[i], res[i+1] = res[i+1], res[i]
					break
				}
			}
		}
	}

	fmt.Println(strings.Join(res, " "))
}
