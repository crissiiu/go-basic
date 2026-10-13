package main

import "fmt"

//Bài 6: Array và slice
/*
Array là mảng có phần tử cố định
Slice là một vùng nhìn vào một mảng nền, với độ dài có thể thay đổi,
*/

func main() {
	productName := "Noodle"
	price := 4.2
	quantities := []int{0, 10, 25, 30, 50}
	quantities = append(quantities, 60)
	for _, quantity := range quantities {
		if quantity == 30 {
			quantity += 10
			continue
		}
		PrintBill(productName, price, quantity)
	}

}

func PrintBill(productName string, price float64, quantity int) {
	if quantity <= 0 {
		fmt.Println("So luong khong hop le")
		return
	}

	total, discount, result := CalculateBill(price, quantity)
	fmt.Printf("%s [%d]: \n"+
		"So tien truoc khi giam: %.2f\n"+
		"So tien giam: %.2f\n"+
		"So tien sau khi giam: %.2f\n",
		productName, quantity, total, discount, result)

}

func CalculateBill(price float64, quantity int) (float64, float64, float64) {

	total := Total(price, quantity)
	discount := Discount(quantity, total)

	return total, discount, total - discount
}
func Total(price float64, quantity int) float64 {
	return price * float64(quantity)
}

func Discount(quantity int, total float64) float64 {
	if quantity >= 50 {
		return total * 0.2
	}

	if quantity >= 20 {
		return total * 0.1
	}

	return 0
}
