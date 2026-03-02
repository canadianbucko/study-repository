package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Text     string `json:"text"`
	Urgent   bool   `json:"urgent"`
	sentTime time.Time
}

var messagesMap map[int]Message = make(map[int]Message)

func acceptMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		msg := "error decoding a json message " + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}

	fakeID := rand.Intn(666)
	message.sentTime = time.Now()

	messagesMap[fakeID] = message

}

var tempMapResponse map[int]Message = map[int]Message{}

func getMessages(w http.ResponseWriter, r *http.Request) { // тут тоже тогда query ?urgent = true
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if urgentQuery := r.URL.Query().Get("urgent"); urgentQuery == "" {
		bytesMap, err := json.Marshal(messagesMap)
		if err != nil {
			msg := "idk mate we fucked up while trying to marshal json " + err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(msg))
			return
		}
		w.Write([]byte(bytesMap))
		return

	}

	urgentQuery := r.URL.Query().Get("urgent")
	urgentBool, err := strconv.ParseBool(urgentQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if urgentBool == true {
		for key, msg := range messagesMap {
			if msg.Urgent == true {
				tempMapResponse[key] = msg
			}

		}
		bytesMappy, err := json.Marshal(tempMapResponse)
		if err != nil {
			msg := "idk mate we fucked up kinda lazy to deal with this stuff, so yeah fuck off"
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(msg))
		}
		w.Write([]byte(bytesMappy))
	}

}

func deleteMessage(w http.ResponseWriter, r *http.Request) { // я хочу это жрать в query параметрах /delete?id=20
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	idConverted, err := strconv.Atoi(id)
	if err != nil {
		msg := "error converting to INT, probably you sent some shit " + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}

	_, ok := messagesMap[idConverted]
	if !ok {
		msg := "it looks like ID wasn't found!"
		fmt.Println(msg)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	delete(messagesMap, idConverted)
}

func main() {
	http.HandleFunc("/message", acceptMessage)
	http.HandleFunc("/list", getMessages)
	http.HandleFunc("/delete", deleteMessage)

	if err := http.ListenAndServe(":11757", nil); err != nil {
		fmt.Println("error starting http server", err)
	}
}
