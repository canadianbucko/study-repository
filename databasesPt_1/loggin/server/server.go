package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"study/helpers"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
)

var ourString helpers.StringsDTO
var sliceOfStrings []helpers.StringsDTO

func StartServer() {
	pid := os.Getpid()
	pp.Println("we're starting the server, our PID is, %v. please take it", pid)

	router := mux.NewRouter()

	if err := http.ListenAndServe(":6969", router); err != nil {
		panic(err)
	}
}

func AcceptString(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := errors.New("we're not doing that, that's a wrong method")
		helpers.SendError(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&ourString); err != nil {
		err := errors.New("someting wong with json")
		helpers.SendError(w, err, http.StatusBadRequest)
		return
	}

	finalString, err := helpers.NewStringDTO(ourString.Text)
	if err != nil {
		helpers.SendError(w, err, http.StatusBadRequest)
		return
	}

	sliceOfStrings = append(sliceOfStrings, finalString) // тут мы добавляем в слайс все
	w.WriteHeader(http.StatusOK)
}

func ReturnStrings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New("we're not doing that, that's a wrong method")
		helpers.SendError(w, err, http.StatusMethodNotAllowed)
		return
	}

	// просто отдать slice of sliceOfStrings
	sliceMarshalled, _ := json.Marshal(sliceOfStrings)
	w.WriteHeader(http.StatusOK)
	w.Write(sliceMarshalled)
}

// — Опишите на Golang программу, которая создаёт
//  https://github.com/uber-go/zap логгер, пишущий в консоль и в файл (аналогично логгеру из видео-урока)
// — Создайте HTTP сервер на 2 эндпоинта:
//   1. `POST /strings`, который просто из входящего тела HTTP запроса читает строку,
//   и кладёт эту строку в слайс внутри программы
//   2. `GET /strings`, который в HTTP ответе возвращает
//   в JSON-виде слайс сохранённых ранее строк
// — Все входящие и исходящие запросы должны логгироваться:
// когда пришёл запрос (time.Now), на какой эндпоинт пришёл запрос,
// какой метод у входящего запроса,
// сколько строк в слайсе оказалось после добавления новой строки `POST /string`,
// сколько времени заняло выполнение HTTP эндпоинта,
// какой статус-код был отправлен в HTTP ответе, время отправки HTTP ответа (time.Now)

// — Запустить написанную программу,
// отправить в неё несколько HTTP запросов,
// проверить корректность логгирования работы программы
