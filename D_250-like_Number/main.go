package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(num int) bool {
	var isPrime bool
	if num == 0 || num == 1 {
		return isPrime
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return isPrime
		}
	}
	isPrime = true
	return isPrime
}

func main() {
	var N int
	fmt.Scan(&N)
	var count int

	elem := N / 2
	loop := int(math.Cbrt(float64(elem)))

	for i := 1; i <= loop; i++ {
		if !isPrimeNumber(i) {
			continue
		}
		for j := i + 1; j <= loop; j++ {
			if !isPrimeNumber(j) {
				continue
			}
			if i*int(math.Pow(float64(j), 3)) <= N {
				count++
			}
		}
	}
	fmt.Println(count)
}
