package main

import "fmt"

type Person struct {
	firtName string
	lastName string
}

func changeName(p *Person) {
	p.firtName = "Bob"
}

func main() {
	person := Person{
		firtName: "Alice",
		lastName: "Dow",
	}

	changeName(&person)

	fmt.Println(person)
}
