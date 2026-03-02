package main

import (
	"fmt"
	"io"
	"net/http"
)

/* — Напишите программу, которая предоставляет возможность принимать
 пришедшие по сети текстовые сообщения и сохранять их внутри себя
— Мы, как клиенты, можем отправлять в эту программу
на определённый эндпоинт свои сообщения по HTTP, а программа сохранит их и выведет на экран
— Необходимо протестировать написанную программу
*/

var messages []string = []string{}

func handler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error parsing HTTP body", err)
		return
	}

	messages = append(messages, string(httpRequestBody)) // а можно и так httpRequestBodyString := string(httpRequestBody)
	fmt.Println(messages)
}

func main() {
	http.HandleFunc("/msg", handler)

	err := http.ListenAndServe(":1317", nil)
	if err != nil {
		fmt.Println("someting wong", err)
	}
}
