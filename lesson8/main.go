package main

import "fmt"

//Bài 8: Map và cách tra cứu dữ liệu theo khóa.

func main() {
	prices := map[string]float64{
		"Noodle": 4.2,
		"Milk":   2.5,
		"Bread":  3.6,
	}
	PrintBill(prices, "Milk")
	PrintBill(prices, "Sit")
}

func FindPrice(prices map[string]float64, productName string) (float64, bool) {
	price, ok := prices[productName]
	return price, ok
}

func PrintBill(prices map[string]float64, productName string) {
	price, ok := FindPrice(prices, productName)
	if !ok {
		fmt.Println("Khong tim thay san pham")
	} else {
		fmt.Println("Gia:", price)
	}
}
