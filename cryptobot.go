
package main

import (
  "bytes"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
  "strings"
  "fmt"
)


type allCoinInfo []singleCoinInfo

type singleCoinInfo struct {
  Price      string `json:"price_usd"`
  Name       string `json:"name"`
  Symbol     string `json:"symbol"`
  Change24h  string `json:"percent_change_24h"`
  BtcPrice   string `json:"price_btc"`
}

func getPrice(coin string) string {
  var cryptoCoins allCoinInfo
  var body []byte

  response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/")
	if err != nil {
    log.Print("cryptobot.go 30: ")
		log.Println(err)
    return retry("crypto", coin)
	}
	defer response.Body.Close()

	// Read the data into a byte slice
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
    log.Print("cryptobot.go 39: ")
		log.Println(err)
    return retry("crypto", coin)
	}
	// Remove whitespace from response
	data := bytes.TrimSpace(body)

	// Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &cryptoCoins)
	if err != nil {
    log.Print("cryptobot.go 49: ")
		log.Println(err)
    return retry("crypto", coin)
	}

  for i := 0; i < len(cryptoCoins); i++ {
    if strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Symbol) { //if the symbols are the same ie btc == btc
      stringToReturn := (cryptoCoins[i].Name + ": " + cryptoCoins[i].Symbol + "\nPrice USD: " + cryptoCoins[i].Price  + "\nPrice BTC: " + cryptoCoins[i].BtcPrice + "\n24hr Change: " + cryptoCoins[i].Change24h + "%")
      fmt.Println("Requested Info On: " + cryptoCoins[i].Name)
      return stringToReturn
    } else if strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Name) { //if the names are the same ie bitcoin == bitcoin
      stringToReturn := (cryptoCoins[i].Name + ": " + cryptoCoins[i].Symbol + "\nPrice USD: " + cryptoCoins[i].Price  + "\nPrice BTC: " + cryptoCoins[i].BtcPrice + "\n24hr Change: " + cryptoCoins[i].Change24h + "%")
      fmt.Println("Requested Info On: " + cryptoCoins[i].Name)
      return stringToReturn
    }
  }
  return "Invalid Token"
}
