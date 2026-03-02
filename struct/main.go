package main

import "fmt"

type Car struct {
	Brand string
	Model string
	HP    int
}

func main() {
	car1 := Car{
		Brand: "Tesla",
		Model: "Model S",
		HP:    500,
	}

	fmt.Println(car1)
	car1.HP += 100
	fmt.Println(car1.HP)
}
