package main

import "fmt"

func main() {
	increment := 0
	doIncrement(&increment)
	fmt.Println("Outside doIncrement", increment)
}

func doIncrement(i *int) {
	*i++
	fmt.Println("Inside doIncrement", *i)
}
