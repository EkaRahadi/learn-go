package main

import (
	"errors"
	"fmt"
)

func main() {
	err := getError()
	fmt.Println(getError())
	fmt.Println(err.Error())
}

func getError() error {
	return errors.New("Terjadi Kesalahan System")
}
