
package main

import (
  "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "strings"
  "strconv"
)

var timesSearchedJSON int

type allCoinInfo []singleCoinInfo

type singleCoinInfo struct {
  Price      string `json:"price_usd"`
  Name       string `json:"name"`
  Symbol     string `json:"symbol"`
  Change24h  string `json:"percent_change_24h"`
  BtcPrice   string `json:"price_btc"`
  Volume     string `json:"24h_volume_usd"`
}

func getPrice(coin string) string {
    var cryptoCoins allCoinInfo
  var body []byte

  response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?start=" + strconv.Itoa(timesSearchedJSON*100))
	if err != nil {
    fmt.Print("cryptobot.go 30: ")
		fmt.Println(err)
    return retry("crypto", coin, "")
	}
	defer response.Body.Close()

	// Read the data into a byte slice
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
    fmt.Print("cryptobot.go 39: ")
		fmt.Println(err)
    return retry("crypto", coin, "")
	}
	// Remove whitespace from response
	data := bytes.TrimSpace(body)

	// Remove leading slashes and blank space to get byte slice that can be unmarshaled from JSON
	data = bytes.TrimPrefix(data, []byte("// "))

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(data, &cryptoCoins)
	if err != nil {
    fmt.Print("cryptobot.go 49: ")
		fmt.Println(err)
    return retry("crypto", coin, "")
	}

  for i := 0; i < len(cryptoCoins); i++ {
    if strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Symbol) || strings.ToLower(coin) == strings.ToLower(cryptoCoins[i].Name) { //if the symbols are the same ie btc == btc
      ethPrice := calculateEthPrice(getPriceEthereum(), cryptoCoins[i].BtcPrice)

      stringToReturn := cryptoCoins[i].Name + ": " + cryptoCoins[i].Symbol + "\nPrice USD: " + cryptoCoins[i].Price  + "\nPrice BTC: " + cryptoCoins[i].BtcPrice + "\nPrice ETH: " + ethPrice + "\n24hr Change: " + cryptoCoins[i].Change24h + "%%" + "\nVolume USD: " + cryptoCoins[i].Volume
      fmt.Println("Requested Info On: " + cryptoCoins[i].Name)
      return stringToReturn
    }
  }
  timesSearchedJSON += 1;
  if timesSearchedJSON > 12 {
    timesSearchedJSON = 0
    return "Token Not Found"
  } else {
    return retry("crypto", coin, "")
  }
  return "Token Not Found 144"

}

func calculateEthPrice(eth float64, btcString string) string {
  btc, err := strconv.ParseFloat(btcString, 64)
  if err != nil {
    fmt.Println(err)
  }
  altEth := btc/eth
  return(FloatToString(altEth))
}

func getPriceEthereum() float64 {
  var ethInfo allCoinInfo
  response, err := http.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")

  if err != nil {
    fmt.Print("cryptobot.go 100: ")
    fmt.Println(err)
    // return retry("crypto", coin, "")
  }
  defer response.Body.Close()

  // Read the data into a byte slice
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Print("cryptobot.go 109: ")
    fmt.Println(err)
    // return retry("crypto", coin, "")
  }
  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &ethInfo)
  if err != nil {
    fmt.Print("cryptobot.go 122: ")
    fmt.Println(err)
    // return retry("crypto", coin, "")
  }
  ethPriceFloat, err := strconv.ParseFloat(ethInfo[0].BtcPrice, 64)
  if err != nil {
    fmt.Println(err)
  }
  // fmt.Println(f)
  return ethPriceFloat

}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
