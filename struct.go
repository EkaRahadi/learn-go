package main

import "fmt"

type Person struct {
	Name, Nationality string
	Age               int
}

func main() {
	var p1 Person
	p1.Name = "John Doe"
	p1.Nationality = "England"
	p1.Age = 25

	p2 := Person{
		Name:        "Smith",
		Nationality: "France",
		Age:         30,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	p1.sayHelloTo("Eka")
	p2.sayHelloTo("Eki")
}

func (person Person) sayHelloTo(name string) {
	fmt.Println("Hello", name, "From", person.Name)
}
