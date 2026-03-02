/* — Мы описываем операцию по изменению размера груди
— Пусть эта операция описана функцией, которая может изменить значение переменной, созданной в main-функции
— Размер груди задан дробной переменной в main'е
— Описать функцию изменения размера груди, протестировать её
*/

package main

import "fmt"

func changeTitsUP(t *float64, change float64) {
	if t != nil {
		*t += change
		fmt.Println("gud")
	} else {
		fmt.Println("fucked up")
		return
	}
}

func main() {
	titsSize := 4.5

	changeTitsUP(&titsSize, 5)
	fmt.Println(titsSize)
}
