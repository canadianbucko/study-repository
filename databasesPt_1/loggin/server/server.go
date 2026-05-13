package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"study/helpers"
	"study/logger"
	"time"

	"github.com/gorilla/mux"
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
)

var ourString helpers.StringsDTO
var sliceOfStrings []helpers.StringsDTO

func StartServer() {
	pid := os.Getpid()
	pp.Println("we're starting the server, our PID is, %v. please take it", pid)

	logger, logFileClose, err := logger.NewLogger("error") // делаем экземпляр логгера
	if err != nil {
		panic(err)
	}
	defer logFileClose()
	withLogger := func(originalFunc func(http.ResponseWriter, *http.Request, *zap.Logger)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			originalFunc(w, r, logger) // вот только тут сука и работает замыкание (сука в анонимной функции? да это
			// да это блять итак понятно что она сука тут работает она же блять в этой же функции лежит ебаная пермеменая
		}
	}
	// logger.Debug("some debug text")
	// logger.Info("some info text")
	// logger.Error("some error text")

	router := mux.NewRouter()

	router.Path("/string").Methods("POST").HandlerFunc(withLogger(AcceptString))
	router.Path("/getstring").Methods("GET").HandlerFunc(withLogger(ReturnStrings))

	if err := http.ListenAndServe(":6969", router); err != nil {
		panic(err)
	}

}

func AcceptString(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	startTime := time.Now()
	if r.Method != http.MethodPost {
		err := errors.New("we're not doing that, that's a wrong method")
		helpers.SendError(w, r, err, http.StatusMethodNotAllowed, logger, startTime)

		return
	}

	if err := json.NewDecoder(r.Body).Decode(&ourString); err != nil {
		err := errors.New("someting wong with json")
		helpers.SendError(w, r, err, http.StatusBadRequest, logger, startTime)
		return
	}

	finalString, err := helpers.NewStringDTO(ourString.Text)
	if err != nil {
		helpers.SendError(w, r, err, http.StatusInternalServerError, logger, startTime)

		return
	}

	sliceOfStrings = append(sliceOfStrings, finalString) // тут мы добавляем в слайс все
	w.WriteHeader(http.StatusOK)
}

func ReturnStrings(w http.ResponseWriter, r *http.Request, logger *zap.Logger) {
	startTime := time.Now()
	if r.Method != http.MethodGet {
		err := errors.New("we're not doing that, that's a wrong method")
		helpers.SendError(w, r, err, http.StatusMethodNotAllowed, logger, startTime)
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
