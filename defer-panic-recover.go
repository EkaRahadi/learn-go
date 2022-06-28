package main

import "fmt"

func main() {
	runApplication(true)
}

func logging() {
	message := recover()
	fmt.Println(message)
	fmt.Println("Log Berjalan")
}

func runApplication(status bool) {
	defer logging()
	if status {
		panic("Application Shutdown")
	}
	fmt.Println("Aplication Running")
}
