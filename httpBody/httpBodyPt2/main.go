package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var Messages map[int]string = make(map[int]string)

func message(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := "sorry, wrong method. Please, use method POST, you just used - " + r.Method
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error parsing http request", err)
		return
	}

	randIntId := rand.Intn(100000)
	Messages[randIntId] = string(httpRequestBody)
	w.Write([]byte("Сообщение успешно добавлено и ему присвоен ID"))
	string := strconv.Itoa(randIntId) + " " + string(httpRequestBody)
	w.Write([]byte(string))

}

var stringSlice []string = make([]string, 0)

func getMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := "sorry, wrong method. Please, use method GET, you just used - " + r.Method
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}
	for k, _ := range Messages { // на каждой итерации я хочу че делать ключ + значение
		tempValue := Messages[k]
		tempKey := k
		tempKeyConverted := strconv.Itoa(tempKey)
		concatenatedString := tempKeyConverted + " " + tempValue
		stringSlice = append(stringSlice, concatenatedString)
	}
	convertedSlice := strings.Join(stringSlice, ", ")
	_, err := w.Write([]byte(convertedSlice))
	if err != nil {
		fmt.Println("you fukin douchbag", err)
		return
	}

}

func deleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := "sorry, wrong method. Please, use method DELETE, you just used - " + r.Method
		fmt.Println(msg)
		w.Write([]byte(msg))
		return
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("idk someting bad happened ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	httpRequestBodyString := string(httpRequestBody)
	id, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "bad request, fuck off, we are waiting for the ID" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest) // сначала хедер а потом уже да пишем!!! сначалаа ставим бед реквест а потом да
		w.Write([]byte(msg))

		return
	}

	_, ok := Messages[id]
	if !ok {
		msg := "there's no such id, you fucking bastard"
		fmt.Println(msg)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}
	delete(Messages, id)

	fmt.Println("yaaay, we deleted something")
	w.Write([]byte("yaaay, we deleted something, go on and try /list to check if its gone (it should be lol)"))

}

func getHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("for this endpoint it is necessary to use method GET!"))
	}
	msg := "thanks for using this handler with GET method , we love you"
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/message", message)
	http.HandleFunc("/list", getMessages)
	http.HandleFunc("/delete", deleteById)
	http.HandleFunc("/get", getHeader)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("something really bad happened", err)
		return
	}
}
