package main

import "fmt"

type Persegi struct {
	sisi int
}

type HitungInterface interface {
	luas() int
}

func main() {
	var persegi HitungInterface
	persegi = Persegi{
		sisi: 3,
	}

	fmt.Println(persegi.luas())
}

func (persegi Persegi) luas() int { //Method
	return persegi.sisi * persegi.sisi
}
