package main

import (
	"encoding/json"
	"fmt"
	dto "lol/DTO"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var mtx sync.Mutex

type Book struct {
	Bookname       string `json:"bookname"`
	Autor          string `json:"author"`
	Pages          int    `json:"pages"`
	Read           bool
	AddedToLibrary time.Time
	ReadTime       *time.Time
}

func (b Book) NewBook(bookname string, author string, pages int) Book {
	return Book{
		Bookname:       bookname,
		Autor:          author,
		Pages:          pages,
		Read:           false,
		AddedToLibrary: time.Now(),
		ReadTime:       nil,
	}
}

var bookMap map[string]Book = make(map[string]Book)

func AddBook(w http.ResponseWriter, r *http.Request) { // приходит джсон и все?
	mtx.Lock()
	defer mtx.Unlock()

	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		msg := "something bad happened"
		dto := dto.NewDto(msg, err)
		dtoMarshalled, err := json.Marshal(dto)
		if err != nil {
			fmt.Println("все по пизде пошло похуй мне ", err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(dtoMarshalled)
		return
	}
	book.Read = false
	book.AddedToLibrary = time.Now()

	bookMap[book.Bookname] = book
	marshalledJson, err := json.Marshal(bookMap)
	if err != nil {
		fmt.Println("пошло в пизду", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(marshalledJson)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	defer mtx.Unlock()
	bookname := mux.Vars(r)["bookname"]
	_, ok := bookMap[bookname]
	if !ok {
		msg := "book doesn't exist."
		dto := dto.NewDto(msg, nil)
		dtoMarshalled, err := json.Marshal(dto)
		if err != nil {
			fmt.Println("все по пизде пошло похуй мне ", err)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(dtoMarshalled)
		return
	}
	book := bookMap[bookname]
	book.Read = true
	timeNow := time.Now()
	book.ReadTime = &timeNow

	bookMap[bookname] = book

	marshalledJsoni, err := json.Marshal(bookMap[bookname])
	if err != nil {
		fmt.Println("нахуй")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledJsoni)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	defer mtx.Unlock()
	bookname := mux.Vars(r)["bookname"]

	_, ok := bookMap[bookname]
	if !ok {
		msg := "book doesn't exist."
		dto := dto.NewDto(msg, nil)
		dtoMarshalled, err := json.Marshal(dto)
		if err != nil {
			fmt.Println("все по пизде пошло похуй мне ", err)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(dtoMarshalled)
		return
	}

	marshalledJsoni, err := json.Marshal(bookMap[bookname])
	if err != nil {
		fmt.Println("нахуй")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledJsoni)

}

func getBookFiltered(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	defer mtx.Unlock()
	author := r.URL.Query().Get("author")
	read := r.URL.Query().Get("read")
	if author != "" && read != "" {
		convRead, err := strconv.ParseBool(read)
		if err != nil {
			fmt.Println("fucking shit shoulda bool! and its not")
			return
		}
		var tempMap map[string]Book = make(map[string]Book)
		for k, v := range bookMap {
			if author == v.Autor {
				if v.Read == convRead {
					tempMap[k] = v
				}
			}
		}
		marshalledJson, err := json.Marshal(tempMap)
		if err != nil {
			fmt.Println("fuck off")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshalledJson)
		return
	}

	if author != "" {
		var tempMap map[string]Book = make(map[string]Book)
		for k, v := range bookMap {
			if author == v.Autor {
				tempMap[k] = v
			}

		}
		marshalledJson, err := json.Marshal(tempMap)
		if err != nil {
			fmt.Println("fuck off")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshalledJson)
		return
	}
	if read != "" {
		convRead, err := strconv.ParseBool(read)
		if err != nil {
			fmt.Println("fucking shit shoulda bool! and its not")
			return
		}
		var tempMap map[string]Book = make(map[string]Book)
		for k, v := range bookMap {
			if v.Read == convRead {
				tempMap[k] = v
			}

		}
		marshalledJson, err := json.Marshal(tempMap)
		if err != nil {
			fmt.Println("fuck off")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(marshalledJson)
		return
	}
	getAllBooks(w, r)
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	defer mtx.Unlock()
	marshalledJson, err := json.Marshal(bookMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledJson)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	mtx.Lock()
	defer mtx.Unlock()
	bookname := mux.Vars(r)["bookname"]

	_, ok := bookMap[bookname]
	if !ok {
		msg := "book doesn't exist."
		dto := dto.NewDto(msg, nil)
		dtoMarshalled, err := json.Marshal(dto)
		if err != nil {
			fmt.Println("все по пизде пошло похуй мне ", err)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(dtoMarshalled)
		return
	}
	delete(bookMap, bookname)

	w.WriteHeader(http.StatusOK)

}

// router.Path("/tasks{name}").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetAllUncompletedTasks)
func main() {
	router := mux.NewRouter()

	router.Path("/book").Methods("POST").HandlerFunc(AddBook)
	router.Path("/book/{bookname}").Methods("PATCH").HandlerFunc(ReadBook)
	router.Path("/book/{bookname}").Methods("GET").HandlerFunc(getBook)
	router.Path("/book").Methods("GET").HandlerFunc(getBookFiltered)
	router.Path("/book").Methods("GET").HandlerFunc(getAllBooks)
	router.Path("/book/{bookname}").Methods("DELETE").HandlerFunc(deleteBook)
	if err := http.ListenAndServe(":9999", router); err != nil {
		fmt.Println("error starting http server", err)
	}
}

// посчитать количество строк в проекте всего
