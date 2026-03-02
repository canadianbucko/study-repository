/* — Создайте указатели на переменные всех изученных типов (int, string, bool, float64)
— Для каждого указателя создайте ему в пару второй указатель с nil-значением
— Опишите для каждого указателя функцию, которая бы принимала указатель, выводила его на экран, а затем проверяла является ли указатель nil-указателем
— Если указатель nil, то выведите на экран сообщение об этом
— Если указатель не nil, то выведите значение переменной, которая лежит под этим указателем
*/

package main

import "fmt"

func main() {
	price := 100

	intPtr := &price

	var intNilPtr *int
	changePrice(intNilPtr)
	changePrice(intPtr)
	fmt.Println(price)

}

func changePrice(n *int) {
	if n != nil {
		*n = 10
	} else {
		fmt.Println("nil pointer! ")

	}
}
