package main

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitConfig(t *testing.T) {
	tests := []struct{
		configName string
		expect error
	}{
		{"", nil},
		{"config", nil},
	}
	for _, test := range tests {
		assert.Equal(t, initConfig(test.configName), test.expect)
	}

	configName := "test"
	assert.NotNil(t, initConfig(configName))
}

func TestLog(t *testing.T) {
	initConfig("")
	outfile := viper.GetString("outfile")


	passer := &DataPasser{message: make(chan string)}
	go func() {
		passer.message <- "1"
		passer.message <- "2"
		close(passer.message)
	}()
	passer.log(outfile)
}

func TestHandler(t *testing.T)  {
	initConfig("")
	outfile := viper.GetString("outfile")
	passer := &DataPasser{message: make(chan string)}

	go passer.log(outfile)
	var req *http.Request
	var res *httptest.ResponseRecorder
	var err error

	tests := [] struct{
		method string
		url string
		expected int
	}{
		{"GET","/?message=test", http.StatusOK},
		{"POST", "/", http.StatusMethodNotAllowed},
		{"GET", "/abc", http.StatusBadRequest},
		{"GET","/?fdsafa=fdsfs", http.StatusOK},
	}
	for _, test := range tests {
		req, err = http.NewRequest(test.method, test.url, nil)
		if err!=nil{
			t.Fatal(err)
		}
		res = httptest.NewRecorder()
		passer.handler(res, req)
		assert.Equal(t, test.expected, res.Code)
	}

}