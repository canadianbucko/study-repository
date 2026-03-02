package main

import (
	"fmt"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	words := "Я собака и я говорю 'Гав'"
	converted := []byte(words)
	_, err := w.Write(converted)
	if err != nil {
		fmt.Println("something went wrong", err.Error())
	}

}

func cat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я кошка и я говрою 'Мяу'"))
	if err != nil {
		fmt.Println("something went wrong", err.Error())
	}

}

func cow(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Я корова и я говорю 'Мууу'"))
	if err != nil {
		fmt.Println("something went wrong", err.Error())
	}

}

func main() {
	http.HandleFunc("/cow", cow)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	err := http.ListenAndServe(":1317", nil)
	if err != nil {
		fmt.Println("someting wong", err.Error())
	}
}
