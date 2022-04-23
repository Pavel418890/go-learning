package main

import "fmt"

func main() {
	// var m1 map[string]string
	// fmt.Printf("%#v\n", m1)
	// m2 := map[int]string{}
	// fmt.Printf("%#v\n", m2)
	// m2[10] = "Abba"

	// fmt.Printf("%v\n", m2[10])
	// fmt.Printf("%v\n", m2[1])

	var m1 map[int]bool
	// m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}
	_, _, _ = m1, m2, m3
	//#ERROR
	// fmt.Println(m2 == m3)

	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)
	for k, v := range m {
		fmt.Printf("k:%d=%t\n", k, v)
	}
}
