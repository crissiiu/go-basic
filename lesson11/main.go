package main

import "fmt"

// 11. Method, value receiver và pointer receiver
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	pizza := Product{ID: 2, Name: "Pizza", Price: 2.8, Stock: 14}

	fmt.Println(pizza)

	pizza.Print(pizza.ReduceStock(4))
	pizza.Print(pizza.Restock(6))
	pizza.Print(pizza.ReduceStock(-2))
}

func (pizza Product) Print(ok bool) {
	if !ok {
		fmt.Println("Khong duoc")
	} else {
		fmt.Println(pizza)
	}

}

func (pizza *Product) ReduceStock(quantity int) bool {
	if pizza == nil || quantity <= 0 || pizza.Stock < quantity {
		return false
	}
	pizza.Stock -= quantity
	return true
}

func (pizza *Product) Restock(quantity int) bool {
	if quantity <= 0 || pizza.Stock < quantity {
		return false
	}
	pizza.Stock += quantity
	return true
}
