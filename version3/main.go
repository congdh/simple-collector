package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func logToFile(filepath string, message string) {

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err!=nil{
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Fprintf(f,"%s %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func handler(w http.ResponseWriter, r *http.Request)  {

	if r.Method != "GET"{
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	filepath := "log.txt"
	message := r.URL.Query().Get("message")
	logToFile(filepath, message)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

