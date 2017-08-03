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

type quandlDataSet struct {
  Dataset    quandlStock    `json"dataset"`
  Error      quandlError    `json"quandl_error"`
}

type quandlStock struct {
  Name    string    `json:"name"`
}

type quandlError struct {
  Code    string    `json"code"`
}

func getStockInfo(stock string, apiKey string) string {
  var stringToReturn string
  var err error
  var stocks stocksInfo
  var quandl quandlDataSet
  var response *http.Response
  var body []byte

  //parse google stocks api
  response, err = http.Get("http://finance.google.com/finance/info?client=ig&q=" + stock)
  checkErr(err)
  defer response.Body.Close()
  body, err = ioutil.ReadAll(response.Body)
	checkErr(err)
  data := bytes.TrimSpace(body)
	data = bytes.TrimPrefix(data, []byte("// "))
	err = json.Unmarshal(data, &stocks)
  if err != nil {
    fmt.Println(err)
    fmt.Println("An error occured on or near like 52, in stockbot.go. The arguement passed was: " + stock)
    return "An error was found. Please retry or contact @gageco"
  }

  //parse quandl api
  quandlUrl := "https://www.quandl.com/api/v3/datasets/WIKI/" + stock + "/metadata.json?api_key=" + apiKey
  response, err = http.Get(quandlUrl)
  checkErr(err)
  defer response.Body.Close()
  body, err = ioutil.ReadAll(response.Body)
  checkErr(err)
  data = bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &quandl)
  if err != nil {
    fmt.Println(err)
    fmt.Println("An error occured on or near like 67, in stockbot.go. The arguement passed was: " + stock)
    return "An error was found. Please retry or contact @gageco"
  }

  if len(quandl.Dataset.Name) != 0 {
    stockName := quandl.Dataset.Name[:len(quandl.Dataset.Name)-45]
    stringToReturn = stockName + "\nPrice: " + stocks[0].Price + "\n24hr Change: " + stocks[0].Change + "\nExchange: " + stocks[0].Exchange
    fmt.Println("Info For Stock: " + stockName + " Shown.")
  } else if len(stocks[0].Ticker) != 0 {
    // fmt.Println("in google stock")
    stringToReturn = stocks[0].Ticker + "\nPrice: " + stocks[0].Price + "\n24hr Change: " + stocks[0].Change + "\nExchange: " + stocks[0].Exchange
    fmt.Println("Info For Stock: " + stocks[0].Ticker + " Shown.")
  } else {
    return "Could not find stock, if you believe this was in error please contact @gageco"
  }

  return stringToReturn
}

func checkErr(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
