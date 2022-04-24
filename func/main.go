package main

import (
	"fmt"
	"math"
)

func cube(a float64) float64 {
	return math.Pow(a, 3)
}

func f1(n uint) (int, int) {

	if n < 0 {
		return 0, 0
	}
	fact := 1
	sum := 0
	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}

func main() {
	fmt.Println(cube(5))
	fmt.Println(f1(23))
}
