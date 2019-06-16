package main

import (
	"fmt"
	"os"
)

func main()  {
	filename := "log.txt"
	f, err := os.Create(filename)
	if err!=nil{
		fmt.Println(err)
		f.Close()
		return
	}

	lines := []string{
		"this is a line of log",
		"2nd line",
	}
	for _, line := range lines {
		fmt.Fprintln(f, line)
	}
	err = f.Close()
	if err!=nil{
		fmt.Println(err)
		return
	}
}