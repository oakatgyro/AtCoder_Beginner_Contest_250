package main

import (
	"fmt"
	"math"
)

var maxP = int(math.Pow10(6))

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func makePrimeList() []int {
	all := make([]bool, maxP)
	primeList := []int{}
	for i := 2; i < maxP; i++ {
		if all[i] == true {
			continue
		}

		primeList = append(primeList, i)

		for j := i; j < maxP; j += i {
			all[j] = true
		}
	}
	// reverse(primeList)
	return primeList
}

func makeValue(p, q int) int {
	v := int(float64(p) * math.Pow(float64(q), 3))
	if v > 4e18 {
		return 4e18
	}
	return v
}

func main() {
	var N int
	fmt.Scan(&N)
	var count int

	primeList := makePrimeList()
	size := len(primeList)
	q := size - 1

	for p := 0; p < size; p++ {
		for p < q && makeValue(primeList[p], primeList[q]) > N {
			q--
		}
		if p >= q {
			break
		}
		count += q - p
	}

	fmt.Println(count)
}
