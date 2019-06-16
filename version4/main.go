package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func logToFile(filepath string, message <-chan string) {

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err!=nil{
		fmt.Println(err)
		f.Close()
		return
	}

	for  {
		data := <- message
		fmt.Fprintf(f,"%s %s\n", time.Now().Format("2006-01-02 15:04:05"), data)
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {

	if r.Method != "GET"{
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/"{
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	message <- r.URL.Query().Get("message")
}

var message = make(chan string)

func main() {
	filepath := "log.txt"
	go logToFile(filepath, message)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

