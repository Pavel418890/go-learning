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

func main() {
	var x = []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}
	var s = 20
	fmt.Println(Gps(s, x))
}
