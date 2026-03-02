package main

import (
	"fmt"
	"study2/module"
	"study2/module/methods"
)

/* — Вы разрабатываете систему отправки уведомлений через разные каналы связи:
    1. email
    2. SMS
    3. push-уведомления
— Необходимо разработать модуль отправки уведомлений,
который одним методом может принять массив уведомлений
на отправку по различным каналам связи, после чего, собственно, отправить все эти уведомления
— Вместо настоящей отправки уведомлений необходимо использовать заглушки,
на манер заглушек методов оплаты для модуля проведения оплат из видео-урока
*/

func main() {

	bulkMessages := []string{"govno", "sosi"}

	fmt.Println(bulkMessages)

	method := methods.NewEmail()

	sendModule := module.NewSendModule(method)
	id := sendModule.Send(bulkMessages)

	fmt.Println(id)

	fmt.Println(sendModule.AllInfo())
	temp := sendModule.ExactMessageById(id)

	fmt.Println(temp)
}
