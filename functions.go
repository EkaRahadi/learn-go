package main

import "fmt"

type Blacklist func(string) bool

func main() {
	sayHelloTo("Eka")
	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)
	fName, lName, age := getPerson()
	fmt.Println(fName, lName, age)

	fmt.Println(sumAll(10, 20, 30, 40, 50))

	numsSLice := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAll(numsSLice...))

	// Function values
	goodBye := sayGoodBye
	fmt.Println(goodBye("Eka"))

	//Functions As Parameters
	filter := spamWordFilter

	fmt.Println(sayWord("Fuck", filter))

	fmt.Println(sayWord("Eka", spamWordFilter))

	checkBlocked := func(name string) bool {
		return name == "root"
	}

	fmt.Println(registerUser("Eka", checkBlocked))
	fmt.Println(registerUser("root", checkBlocked))
	fmt.Println(registerUser("admin", func(name string) bool {
		return name == "admin"
	}))

	recursive := factorialRecursive(3)
	fmt.Println(recursive)

}

func sayHelloTo(firstname string) {
	fmt.Println("Hello", firstname)
	result := multiplyInt(1, 2)
	fmt.Println(result)
}

func multiplyInt(a, b int) int {
	return a * b
}

// Multiple Return Values
func getFullName() (string, string) {
	return "Eka", "Rahadi"
}

// Named Return Values Function
func getPerson() (firstname, lastname string, age int) {
	firstname = "Eka"
	lastname = "Rahadi"
	age = 24
	return
}

// Variadic Function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func spamWordFilter(word string) string {
	if word == "Fuck" {
		return "..."
	} else {
		return word
	}
}

func sayWord(word string, filter func(string) string) string {
	filtered := filter(word)

	return "The Word is " + filtered
}

func registerUser(name string, isBlocked Blacklist) string {
	if isBlocked(name) {
		return "You are blocked"
	} else {
		return "Welcom " + name
	}
}

func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorialRecursive(n-1)
}
