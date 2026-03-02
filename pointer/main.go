/*
— Создайте указатели на переменные всех изученных типов (int, string, bool, float64)
— Выведите на экран адреса в памяти, которые в себе хранят эти указатели
— Выведите на экран значения переменных, на которые указатели указывают, при помощи разыменования указателей
*/

package main

import "fmt"

func main() {
	boolean := true
	string := "hello"

	stringPtr := &string
	boolPtr := &boolean

	fmt.Println(stringPtr)
	fmt.Println(boolPtr)

	fmt.Println(*stringPtr)
	fmt.Println(*boolPtr)

}
