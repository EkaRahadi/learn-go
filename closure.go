package main

import "fmt"

func main() {
	counter := 0
	name := "Eka"

	increment := func() {
		name := "Eki"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
