package main

import "fmt"

func main() {
	var result interface{} = giveMeAThing(1)
	var resultCast int = result.(int)
	var resultCast2 string = result.(string) //panic in runtime
	fmt.Println(resultCast)
	fmt.Println(resultCast2)
}

func giveMeAThing(i int) interface{} {
	if i == 0 {
		return false
	} else if i == 1 {
		return 1
	} else {
		return "Hello"
	}
}
