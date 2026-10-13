package main

import "fmt"

//Bài 7: len, cap, make và cách append

func main() {
	a := make([]int, 2, 3)
	a[0] = 10
	a[1] = 20
	b := a
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(b))
	fmt.Println(cap(b))
	a = append(a, 30)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(b))
	fmt.Println(cap(b))
	a = append(a, 40)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(b))
	fmt.Println(cap(b))
	a[0] = 99
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println(a) //99 20 40
	fmt.Println(b) //99 20

}
