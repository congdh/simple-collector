package main

import (
	"fmt"
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

type DataPasser struct {
	message chan string
}

func (p *DataPasser) log(outfile string) {
	f, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	for {
		data, ok := <-p.message
		if !ok {
			return
		}
		fmt.Fprintf(f, "%s %s\n", time.Now().Format("2006-01-02 15:04:05"), data)
	}
}

func (p *DataPasser) handler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	p.message <- r.URL.Query().Get("message")
}

var message = make(chan string)

func main() {
	err := initConfig("")
	if err != nil {
		fmt.Println(err)
		return
	}

	outfile := viper.GetString("outfile")
	passer := &DataPasser{message: make(chan string)}
	go passer.log(outfile)
	http.HandleFunc("/", passer.handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
