package main

import (
  "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "strings"
)


type allCoinInfo []singleCoinInfo

type singleCoinInfo struct {
  Price      string `json:"price_usd"`
  Name       string `json:"name"`
  Symbol     string `json:"symbol"`
  Change24h  string `json:"percent_change_24h"`
}

func getPrice(coin string) string {
  var err error
  var cryptoCoins allCoinInfo
  var response *http.Response
  var body []byte

  response, err = http.Get("https://api.coinmarketcap.com/v1/ticker/")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// Read the data into a byte slice
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Remove whitespace from response
	data := bytes.TrimSpace(body)

	// Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &cryptoCoins)
	if err != nil {
		fmt.Println(err)
	}

  for i := 0; i < len(cryptoCoins); i++ {
    if strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Symbol) { //if the symbols are the same ie btc == btc
      stringToReturn := (cryptoCoins[i].Name + ": " + cryptoCoins[i].Symbol + "\nPrice: " + cryptoCoins[i].Price + "\n24hr Change: " + cryptoCoins[i].Change24h + "%")
      fmt.Println("Requested Info On: " + cryptoCoins[i].Name)
      return stringToReturn
    }

    if strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Name) { //if the names are the same ie bitcoin == bitcoin
      stringToReturn := (cryptoCoins[i].Name + ": " + cryptoCoins[i].Symbol + "\nPrice: " + cryptoCoins[i].Price + "\n24hr Change: " + cryptoCoins[i].Change24h + "%")
      fmt.Println("Requested Info On: " + cryptoCoins[i].Name)
      return stringToReturn
    }
  }
  fmt.Println("Requested Info On Invalid Token")
  return "Invalid Token"
}