package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var InfoSlice []info

type info struct {
	Name    string  `json:"name"`
	Adress  string  `json:"adress"`
	Age     int     `json:"age"`
	Married bool    `json:"married"`
	Height  float64 `json:"height"`
	Time    time.Time
}

func postInfo(w http.ResponseWriter, r *http.Request) {

	var info info
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		msg := "error parsing JSON request" + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	info.Time = time.Now()
	InfoSlice = append(InfoSlice, info)
	fmt.Println(info)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(InfoSlice)
	if err != nil {
		fmt.Println("not the best marshalling I ever had.", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("error writing the response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/postinfo", postInfo)
	http.HandleFunc("/getinfo", getInfo)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("idk, smth bad happened to us", err)
		return
	}
}
