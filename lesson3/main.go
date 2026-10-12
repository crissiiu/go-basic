// Bài 2: Bài 3: Điều kiện if và switch.
package main

import "fmt"

func main() {
	productName := "Noodle"
	price := 4.2

	quantity := 0
	PrintBill(productName, price, quantity)

	quantity = 10
	PrintBill(productName, price, quantity)

	quantity = 25
	PrintBill(productName, price, quantity)

	quantity = 50
	PrintBill(productName, price, quantity)
}

func PrintBill(productName string, price float64, quantity int) {
	if quantity <= 0 {
		fmt.Println("So luong khong hop le")
		return
	}

	fmt.Println(productName, ":", Total(price, quantity))
}

func Total(price float64, quantity int) float64 {
	total := price * float64(quantity)
	return Discount(quantity, total)
}

func Discount(quantity int, total float64) float64 {
	if quantity >= 50 {
		return total * 0.8
	}

	if quantity >= 20 {
		return total * 0.9
	}

	return total
}
