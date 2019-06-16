package main

import (
	"fmt"
	"github.com/cskr/pubsub"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"time"
)

func initConfig(configName string) error {
	if configName == "" {
		configName = "config"
	}
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Config not found: %s \n", configName)
	}
	return nil
}

func logToConsole(ch chan interface{})  {
	for {
		if msg, ok := <-ch; ok {
			fmt.Printf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
		} else {
			break
		}
	}
}

func logToFile(filepath string, ch chan interface{})  {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err!=nil{
		fmt.Println(err)
		f.Close()
		return
	}

	for {
		if msg, ok := <-ch; ok {
			fmt.Fprintf(f,"%s %s\n", time.Now().Format("2006-01-02 15:04:05"), msg)
		} else {
			break
		}
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
	msg := r.URL.Query().Get("message")
	ps.Pub(msg, topic)
}

const topic = "topic"
var ps  *pubsub.PubSub

func main() {
	err := initConfig("")
	if err != nil {
		fmt.Println(err)
		return
	}

	ps = pubsub.New(5)
	outfile := viper.GetString("output.file.filename")
	keys := viper.AllKeys()
	keys[0] = "afdsaf"
	if outfile != "" {
		sub := ps.Sub(topic)
		go logToFile(outfile, sub)
	}
	console := viper.GetBool("output.console")
	if console{
		sub2 := ps.Sub(topic)
		go logToConsole(sub2)
	}


	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
