package main

import "fmt"

type Flat struct {
	Adress string
	Size   int
	Floor  int
	Price  int
}

func NewFlat(adress string, size int, floor int, price int) Flat {
	if adress == "" {
		fmt.Println("адресс не может быть пустой")
		return Flat{}
	}
	if size <= 10 {
		fmt.Println("размер не может быть меньше 0")
		return Flat{}
	}
	if floor < 0 || floor >= 100 {
		fmt.Println("нельзя этаж меньше 0 или больше 100")
		return Flat{}
	}
	if price <= 0 {
		fmt.Println("цеа не может быт меньше 0")
		return Flat{}
	}

	return Flat{
		Adress: adress,
		Size:   size,
		Floor:  floor,
		Price:  price,
	}
}
func (f *Flat) ChangePRICE(price int) {
	if f.Price+price <= 0 {
		fmt.Println("цена ниже 0 нет")
		return
	}
	f.Price += price
}

func main() {
	flat1 := NewFlat("залупки 10", 20, 5, 5991)

	fmt.Println(flat1)
	flat1.ChangePRICE(4000)
	fmt.Println(flat1.Price)
}
