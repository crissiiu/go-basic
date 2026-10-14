package main

import "fmt"

//12: Interface

// Interface mô tả những method mà một kiểu cần có để thực hiện một vai trò. Nó mô tả hành vi, không chứa các trường dữ liệu như struct.

type Named interface {
	GetName() string
}

type Product struct {
	id   int
	name string
}

type Service struct {
	id          int
	serviceName string
}

func main() {
	p := Product{id: 1, name: "Pizza"}
	s := Service{id: 1, serviceName: "Delivery"}
	Print(p)
	Print(s)
}

func Print(item Named) {
	fmt.Println(item.GetName())
}

func (p Product) GetName() string {
	return p.name
}

func (s Service) GetName() string {
	return s.serviceName
}
