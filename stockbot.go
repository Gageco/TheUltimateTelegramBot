package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "bytes"
  "time"
  "strings"
  "strconv"
)

var stockInterface interface{}

func getCurrentDate() string {
  CurrentDate := time.Now().UTC()
  newDate := CurrentDate

  if CurrentDate.Weekday() == 0 {
    newDate = CurrentDate.AddDate(0, 0, -2)
  } else if CurrentDate.Weekday() == 6 {
    newDate = CurrentDate.AddDate(0, 0, -1)
  }
  ReturnedString := newDate.Format("2006-01-02")
  return ReturnedString
}

func tempStockFunc(arg string, api string) string {
  fmt.Println(api)
  getCurrentDate()
  return "Sorry stock functionality is not working right now."
}

func getStockInfo(stock string, apiKey string) string {
  fmt.Println("Stock Function: " + stock)
  var stringToReturn string
  dateFound := false

  ApiURL := "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + stock + "&apikey=" + apiKey

  //parse api response
  response, err := http.Get(ApiURL)
  if err != nil {
    fmt.Print("stockbot.go 49: ")
    fmt.Println(err)
    return "An error was found. Please retry or contact @gageco"
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Print("stockbot.go 52: ")
    fmt.Println(err)
    return "An error was found. Please retry or contact @gageco"
  }

  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &stockInterface)
  if err != nil {
    fmt.Print("stockbot.go 70: ")
    fmt.Println(err)
    return "An error was found. Please retry or contact @gageco"
  }

  inStockInterface := stockInterface.(map[string]interface{})
  for possibleError, inMap00 := range inStockInterface {

    if possibleError == "Error Message" {
      stringToReturn = "Stock Not Found"
      break
    }
    inMap01 := inMap00.(map[string]interface{})
    for mapDate, mapDateStock  := range inMap01 {
      if mapDate == getCurrentDate() {
        mapDateStockData := mapDateStock.(map[string]interface{})
        openStock := mapDateStockData["1. open"]
        // highStock := mapDateStockData["2. high"]
        // lowStock := mapDateStockData["3. low"]
        closeStock := mapDateStockData["4. close"]

        pullLength := len(closeStock.(string))-2

        intOpenStock, err := strconv.Atoi(openStock.(string)[:pullLength-3])
        if err != nil {
          fmt.Println(err)
        }
        intCloseStock, err := strconv.Atoi(closeStock.(string)[:pullLength-3])
        if err != nil {
          fmt.Println(err)
        }

        changeStock := (float64(intCloseStock) - float64(intOpenStock))/float64(intOpenStock) * 100

        stringToReturn = "Stock $" + strings.ToUpper(stock) + "\nPrice: " + closeStock.(string)[:pullLength] + "\nOpen: " + openStock.(string)[:pullLength] + "\nChange: " + strconv.FormatFloat(changeStock, 'f', 2, 64) + "%%"
        dateFound = true
      }
    }
    if dateFound == false{
      stringToReturn = "Date Not Found"
    }
  }

  return stringToReturn
}
