package main

import (
	"fmt"
	"strings"

	"unicode/utf8"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love go programming", "love")
	p(result)
	result = strings.ContainsAny("success", "xys")
	p(result)
	result = strings.ContainsRune("golang", 'g')
	p(result)
	n := strings.Count("cheese", "e")
	p(n)

	var name string = "Pavel"
	country := "Russia"
	fmt.Printf("Name: %s\nCountry: %s\n", name, country)

	fmt.Printf(`a) He says: hello` + "\n")
	fmt.Printf("b) C:\\Users\\a.txt\n")

	r := 'ă'
	fmt.Printf("%T\n", r)

	ma, m := "ma", "m"
	mama := ma + m + string(r)
	p(mama)
	fmt.Printf("%T\n", mama)
	p(len(mama))

	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v", s1[i])
	}
	p()

	for i, r := range s1 {
		fmt.Printf("byte index: %d value: %c\n", i, r)

		s1 := "Go is cool!"
		x := s1[0]
		fmt.Println(x)

		// s1[0] = 'x'
		strings.Replace(s1, "G", "x", 1)

		// printing the number of runes in the string
		// fmt.Println(len(s1))
		p(utf8.RuneCountInString(s1))
	}
	s := "你好 Go!"
	b := []byte(s)
	fmt.Printf("%#v\n", b)
	for i, b := range b {
		fmt.Printf("byte index %d byte %d\n", i, b)
	}

	r2 := []rune(s)
	fmt.Printf("%#v\n", r2)
	for i, r := range r2 {
		fmt.Printf("index: %d rune: %c\n", i, r)
	}

}
