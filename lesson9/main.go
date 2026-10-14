package main

import "fmt"

//Bài 9: Struct và cách mô hình hóa sản phẩm

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func main() {
	products := map[int]Product{
		1: {ID: 1, Name: "Milk", Price: 2.5, Stock: 20},
		2: {ID: 2, Name: "Bread", Price: 3.6, Stock: 10},
		3: {ID: 2, Name: "Pizza", Price: 2.8, Stock: 14},
	}
	id := 3
	product, ok := FindProduct(products, id)
	if !ok {
		fmt.Println("San pham khong ton tai")
		return
	}

	if product.Stock <= 0 {
		fmt.Println("San pham da het hang")
		return
	}

	product.Stock -= 1
	products[id] = product // Ghi giá trị đã sửa trở lại map

	fmt.Printf(
		"Name: %s\nPrice: %.2f\nStock: %d\n",
		product.Name,
		product.Price,
		product.Stock,
	)

}

func FindProduct(products map[int]Product, id int) (Product, bool) {
	product, ok := products[id]
	if ok {
		product.Stock -= 1
		return product, true
	}
	return product, false
}
