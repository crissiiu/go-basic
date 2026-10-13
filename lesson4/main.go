package main

import "fmt"

//Bài 4: Vòng lặp for

func main() {
	productName := "Noodle"
	price := 4.2

	quantity := 0
	for quantity <= 60 {
		if quantity == 30 {
			quantity += 10
			continue
		}
		PrintBill(productName, price, quantity)
		
		quantity += 10
	}

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
