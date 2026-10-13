package main

import "fmt"

// Bài 5: Hàm, tham số và giá trị trả về

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
