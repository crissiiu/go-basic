package main

import "fmt"

func main() {
	// Demo()
	// Demo2()
	Demo3()
}

func Demo() {
	fmt.Println("Bat dau")
	defer fmt.Println("Don dep")
	fmt.Println("Dang xu ly")
}

func Demo2() {
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
}

func Demo3() {
	quantity := 4
	defer fmt.Println("A:", quantity)

	quantity = 6
	defer fmt.Println("B:", quantity)

	fmt.Println("C:", quantity)
}
