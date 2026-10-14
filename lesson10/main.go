package main

import "fmt"

//Bài 10: Con trỏ, & và *

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	pizza := Product{ID: 2, Name: "Pizza", Price: 2.8, Stock: 14}

	fmt.Println(pizza)

	Print(&pizza, 4)

	Print(&pizza, 20)

	Print(nil, 20)

}

func Print(pizza *Product, quantity int) {
	ok := ReduceStock(pizza, quantity)
	if !ok {
		fmt.Println("Khong duoc")
	} else {
		fmt.Println(pizza)
	}

}

func ReduceStock(pizza *Product, quantity int) bool {
	if pizza == nil || quantity <= 0 {
		return false
	}
	if pizza.Stock < quantity {
		return false
	}
	pizza.Stock -= quantity
	return true
}
