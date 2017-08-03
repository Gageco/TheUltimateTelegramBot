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
}

type quandlStock struct {
  Name    string    `json:"name"`
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
	checkErr(err)

  //parse quandl stock api
  quandlUrl := "https://www.quandl.com/api/v3/datasets/WIKI/" + stock + "/metadata.json?api_key=" + apiKey
  response, err = http.Get(quandlUrl)
  checkErr(err)
  defer response.Body.Close()
  body, err = ioutil.ReadAll(response.Body)
  checkErr(err)
  data = bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &quandl)
  checkErr(err)

  if stocks[0].Exchange != "OTCMKTS" || stocks[0].Exchange != "HKG" {
    stockName := quandl.Dataset.Name[:len(quandl.Dataset.Name)-45]
    stringToReturn = stockName + "\nPrice: " + stocks[0].Price + "\n24hr Change: " + stocks[0].Change + "\nExchange: " + stocks[0].Exchange
    fmt.Println("Info For Stock: " + stockName + " Shown.")
  } else {
    stringToReturn = stocks[0].Ticker + "\nPrice: " + stocks[0].Price + "\n24hr Change: " + stocks[0].Change + "\nExchange: " + stocks[0].Exchange
    fmt.Println("Info For Stock: " + stocks[0].Ticker + " Shown.")
  }

  return stringToReturn
}

func checkErr(err error) {
  if err != nil {
    fmt.Println(err)
  }
}
