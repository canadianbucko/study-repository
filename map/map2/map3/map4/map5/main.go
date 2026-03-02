/*
— Допустим мы с вами описываем автомобильную парковку
— На парковке есть несколько мест, каждое обозначено каким-то строковым значением (A4, B8, D32 и так далее)
— Каджое парковочное место имеет какую-то свою определённую стоимость за час парковки
— Нужно создать все парковочные места с их стоимостью, и вывести на экран только парковочные места стоимостью меньше 500 рублей
— Кажому месту стоимостью больше 900 рублей нужно сделать скидку 10% (было 1100 рублей стало 990)
— Необходимо выбрать наиболее подходящую для решения задачи структуры данных и реализовать описанные пункты
*/

package main

import "fmt"

type ParkingUnit struct {
	Name  string
	Price float64
}

func main() {
	parkingunits := map[string]ParkingUnit{
		"A4": ParkingUnit{
			Name:  "A4",
			Price: 400,
		},
		"A6": ParkingUnit{
			Name:  "A6",
			Price: 1000,
		},
		"A7": ParkingUnit{
			Name:  "A7",
			Price: 700,
		},
	}

	for i, v := range parkingunits {
		if v.Price <= 500 {
			fmt.Println("дешевые места", i, parkingunits[i].Price)
		}
	}
	for i, v := range parkingunits {
		if v.Price >= 900 {
			v.Price *= 0.9
			parkingunits[i] = v
			fmt.Println(parkingunits)

		}
	}
}
