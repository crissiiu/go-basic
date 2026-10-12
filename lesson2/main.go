// Bài 2: Biến, hằng số và kiểu dữ liệu.
package main

import "fmt"

func main() {
	productName := "Noodle"
	price := 4.2
	quantity := 25

	total := price * float64(quantity)

	fmt.Println(productName, ":", total)
}

// := dùng dể khai báo biến
// = gán giá trị cho biến
// const là khai báo hằng -> không thể thay đổi giá trị của biến có dạng const
