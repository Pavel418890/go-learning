package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func myFunc(n string) int {
	total := 0
	for i := 1; i < 4; i++ {
		nn, err := strconv.Atoi(strings.Repeat(n, i))
		if err != nil {
			fmt.Printf("Cannot convert %q to a string\n", n)
			os.Exit(1)
		} else {
			total += nn
		}

	}
	return total
}
func sum(n ...int) (s int) {
	for _, v := range n {
		s += v
	}
	return s
}

func searchItem(s []string, b string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func f2(n int) {
	fmt.Println(n)
}
func main() {
	fmt.Println(cube(5))
	fmt.Println(f1(23))
	fmt.Println(myFunc("5"))
	fmt.Println(sum(1, 2, 3, 4, 5))
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
	x := f2
	fmt.Printf("%T\n", x)
	x(8)
}
