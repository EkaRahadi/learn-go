package main

import "fmt"

func main() {
	var ups interface{} = Ups(1)
	var ups2 bool = Ups(2).(bool)
	// var upsErr bool = Ups(1).(bool) //This will cause error in runtime not compile time
	fmt.Println(ups)
	fmt.Println(ups2)
	// fmt.Println(upsErr)
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Upss..."
	}
}
