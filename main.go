package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	testcall()
}

const API_URL = "https://www.alphavantage.co/query"

type AlphaData struct {
	Function   string `json:"function"`
	Symbol     string `json:"symbol"`
	OutputSize string `json:"outputsize"`
	DataType   string `json:"datatype"`
	APIKey     string `json:"apikey"`
}

func testcall() {
	data := AlphaData{
		Function:   "TIME_SERIES_DAILY",
		Symbol:     "MSFT",
		OutputSize: "compact",

		APIKey: os.Getenv("API_TOKEN"),
	}
	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("%s?function=%s&symbol=%s&outputsize=%s&apikey=%s", API_URL, data.Function, data.Symbol, data.OutputSize, data.APIKey))
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("Error: " + err.Error())
	}
	fmt.Print(string(b))
}
