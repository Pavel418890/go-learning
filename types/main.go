package main

import (
	"fmt"
)
type money float64 

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * .09
}
func (b *book) discount() {
	(*b).price = b.price * .9
}

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string{
	return fmt.Sprintf("%.2f", m)
}

func (b *book) changePrice(p float64) {
    b.price = p
}
func main() {
	var salary money = 1.5 * 5.608
	salary.print() 
	s:=salary.printStr()
	fmt.Printf("%s %T\n", s, s)
	bestBook := book{title: "The best of the best", price: 9.9}
	
	vat := bestBook.vat()
	fmt.Printf("%v\n", vat)
	
	bestBook.discount()
	fmt.Printf("%v\n", bestBook)

	bestBook.changePrice(11.99)
	fmt.Printf("Book's price %.2f\n", bestBook.price)

}