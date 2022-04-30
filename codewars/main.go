package main

import (
	"fmt"
	"math"
)

func Gps(s int, x []float64) int {
	prevDistance := x[1:][0]
	var maxAvgSpeed float64
	for _, v := range x[2:] {
		delta := v - prevDistance
		avgSpeed := 3600.0 * delta / float64(s)
		if avgSpeed > maxAvgSpeed {
			maxAvgSpeed = avgSpeed
		}
		prevDistance = v
	}

	return int(math.Floor(maxAvgSpeed))
}
func test(args ...func([]int)) bool {
	result := true
	for _, fn := range args {
		arr := []int{10, 2, 5, 6, 5, 3, 9, 1, 8, 3}
		sortedArr := []int{1, 2, 3, 3, 5, 5, 6, 8, 9, 10}
		fn(arr)
		for i := 0; i < len(arr); i++ {
			if len(arr) != len(sortedArr) || arr[i] != sortedArr[i] {
				result = false
				break
			}
		}
		if result != true {
			fmt.Println("Test failed")
			break
		}

	}
	return result
}

func insert_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for k := i; k > 0 && arr[k-1] > arr[k]; k-- {
			arr[k], arr[k-1] = arr[k-1], arr[k]
		}
	}
}
func choice_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for k := i + 1; k < len(arr); k++ {
			if arr[k] < arr[i] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
}

func bubble_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for k := 0; k < len(arr)-i; k++ {

			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
				fmt.Println(arr)
			}
		}
	}
}

func generateNumbers(N, M int, prefix *[]int) {
	if M == 0 {
		fmt.Println(*prefix)
		return
	}

	for digit := 0; digit < N; digit++ {
		*prefix = append((*prefix), digit)
		generateNumbers(N, M-1, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]
	}
}

func main() {
	// var x = []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}
	// var s = 20
	// fmt.Println(Gps(s, x))
	fmt.Println(test(insert_sort, choice_sort, bubble_sort))
	prefix := []int{}
	generateNumbers(3, 3, &prefix)
}
