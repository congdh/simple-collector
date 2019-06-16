package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRules(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] =letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main()  {
	filename := "log.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err!=nil{
		fmt.Println(err)
		f.Close()
		return
	}

	for i := 0; i < 2; i++ {
		ts := time.Now()
		message := RandStringRules(10)
		fmt.Fprintf(f,"%s %s\n", ts.Format("2006-01-02 15:04:05"), message)
	}
	err = f.Close()
	if err!=nil{
		fmt.Println(err)
		return
	}
}