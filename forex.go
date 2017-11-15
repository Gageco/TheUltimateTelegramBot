package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "bytes"
  // "time"
  "strings"
  "strconv"
)

var forexInterface interface{}


func getForex(forex string) string {
  fmt.Println("Forex: " + forex)
  var stringToReturn string

  forexArr := strings.Split(forex, " ")

  ApiURL := "https://api.fixer.io/latest?base=" + forexArr[0]

  response, err := http.Get(ApiURL)
  if err != nil {
    fmt.Print("stockbot.go 45: ")
    fmt.Println(err)
    return retry("forex", forex, "")
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Print("stockbot.go 52: ")
    fmt.Println(err)
    return retry("forex", forex, "")
  }

  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &forexInterface)
  if err != nil {
    fmt.Print("stockbot.go 61: ")
    fmt.Println(err)
    return retry("forex", forex, "")
  }

  inForexInterface := forexInterface.(map[string]interface{})
  for baseInterface, dataInterface := range inForexInterface {
      if baseInterface == "rates" {
        dataInterf := dataInterface.(map[string]interface{})
        for baseInter, dataInter := range dataInterf {
          if baseInter == strings.ToUpper(forexArr[1]) {
            // fmt.Println(forexArr[1], dataInter)
            stringToReturn = strings.ToUpper(forexArr[0]) + "/" + strings.ToUpper(forexArr[1]) + ": " +  strconv.FormatFloat(dataInter.(float64), 'E', -1, 64)[:6]
          }
        }
      }
  }

  return stringToReturn

}
