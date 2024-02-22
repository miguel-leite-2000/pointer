package main

import "fmt"

type Car struct {
	Name string
}

func (c *Car) walked() {
	c.Name = "BMW"
	fmt.Println(c.Name, "walked")
}

func main() {
	// a := 10

	// var pointer *int = &a
	// *pointer = 50
	// b := &a
	// *b = 60
	// fmt.Println(*b)

	// variable := 10

	// abc(&variable)

	// fmt.Println(variable)

	car := Car{
		Name: "Jetur",
	}

	car.walked()
	fmt.Println(car.Name)
}

func abc(a *int) {
	*a = 200
}
