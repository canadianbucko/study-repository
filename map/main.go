/*
— Создать несколько разных массивов разных типов
— Проделать все изученные над массивом операции
(получение отдельного элемента, вывод на экран, изменение отдельного элемента, массивы + циклы)
*/

package main

import "fmt"

func main() {
	newInts := [5]int{1, 2, 3, 4, 5}

	//newStrings := [2]string{"hey", "hello"}

	fmt.Println(newInts[1])
	newInts[1] = 5
	for i := 0; i < len(newInts); i++ {
		fmt.Println(newInts[i])
	}

}
