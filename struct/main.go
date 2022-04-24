package main

import "fmt"

func main() {
	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name           string
		age            int
		favoriteColors []string
		gradesInfo     grades
	}
	me := person{
		name: "Pavel", age: 28,
		favoriteColors: []string{"white", "blue"},
		gradesInfo: grades{
			course: "golang",
			grade:  98,
		},
	}
	you := person{
		name:           "Andrei",
		age:            32,
		favoriteColors: []string{"yellow", "red"},
		gradesInfo: grades{
			grade:  89,
			course: "Python 3 Deep Dive",
		},
	}
	fmt.Printf("%#v\n%#v\n", me, you)

	me.name = "Andrei"
	you.favoriteColors = []string{"white", "green"}
	you.favoriteColors = append(you.favoriteColors, "cian")
	fmt.Printf("%#v\n%#v\n", me, you)
	for _, v := range me.favoriteColors {
		fmt.Println(v)
	}

}
