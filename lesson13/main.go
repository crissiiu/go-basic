package main

import (
	"errors"
	"fmt"
)

//13: Xử lý lỗi bằng error

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	pizza := Product{ID: 2, Name: "Pizza", Price: 2.8, Stock: 14}

	fmt.Println(pizza)
	err := pizza.ReduceStock(4)
	pizza.Print(err)

	err = pizza.Restock(6)
	pizza.Print(err)

	err = pizza.Restock(-2)
	pizza.Print(err)

	var missing *Product
	err = missing.Restock(6)
	if err != nil {
		fmt.Println("Loix:", err)
	}
}

func (pizza Product) Print(e error) {
	if e != nil {
		fmt.Println("Loi:", e)
	} else {
		fmt.Println(pizza)
	}

}

func (pizza *Product) ReduceStock(quantity int) error {
	if pizza == nil {
		return errors.New("Sản phẩm không tồn tại")
	}

	if quantity <= 0 {
		return errors.New("So luong phai lon hon khong")
	}
	if pizza.Stock < quantity {
		return fmt.Errorf("khong du hang: con %d, can %d",
			pizza.Stock, quantity,
		)
	}
	pizza.Stock -= quantity
	return nil
}

func (pizza *Product) Restock(quantity int) error {
	if pizza == nil {
		return errors.New("san pham khong ton tai")
	}
	if quantity <= 0 {
		return errors.New("so luong phai lon hon khong")
	}

	pizza.Stock += quantity
	return nil
}
