package main

import (
	"fmt"
)

func main() {
	var R, H, C, W int
	fmt.Scan(&H, &W)
	fmt.Scan(&R, &C)

	var res int

	if C != 1 {
		res++
	}

	if C != W {
		res++
	}

	if R != 1 {
		res++
	}

	if R != H {
		res++
	}

	fmt.Println(res)

}
