package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "bytes"
)

type stocksInfo []singleStock

type singleStock struct {
	Price    string `json:"l"`
	Change   string `json:"c"`
  Exchange string `json:"e"`
  Ticker   string `json:"t"`
}

var err error
var stocks stocksInfo
var response *http.Response
var body []byte

func getStockInfo(stock string) string {
  response, err = http.Get("http://finance.google.com/finance/info?client=ig&q=" + stock)

  if err != nil {
		fmt.Println(err)
	}
  defer response.Body.Close()

  // Read the data into a byte slice
  body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

  //get rid of whitespace
  data := bytes.TrimSpace(body)

  // Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &stocks)
	if err != nil {
		fmt.Println(err)
	}

  //ok so straight up i have no idea what this all does but it works. I found it here: https://github.com/tjaensch/stockticker/blob/master/stockticker.go
  //So thanks to that dude. I apprecate your code, this is going to be helpful in the future
  stringToReturn := "Stock: " + stocks[0].Ticker + "\nCurrent Price: " + stocks[0].Price + "\n24hr Change: " + stocks[0].Change + "\nExchange: " + stocks[0].Exchange
  fmt.Println("Info For Stock: " + stocks[0].Ticker + " Shown.")

  return stringToReturn
}
