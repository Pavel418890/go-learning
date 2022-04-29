package main

import "fmt"

func swap(arg1, arg2 *float64) {
	*arg1, *arg2 = *arg2, *arg1

}


func main(){
	// x:= 10.10
	// fmt.Println(&x)
	// ptr:= &x
	// fmt.Printf("%v %T\n", ptr, ptr)
	// fmt.Printf("%v %v\n", *ptr, &ptr)
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := *ptrx / *ptry
	fmt.Printf("%v\n", z)
	a, b:= 5.5, 8.8
	swap(&a, &b)
	fmt.Printf("%v %v\n", a, b)
	fmt.Printf("%v\n", int(a / b))
}