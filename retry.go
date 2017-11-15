package main

import "fmt"

var numberOfRetrys int
var previousFunc string

// func initializeRetry() {
//   numberOfRetrys = 0
// }

func retry(recentFunc string, passedArg string, passedArg1 string) string {
  numberOfRetrys += 1

  fmt.Println("Retrying: ", recentFunc)
  if recentFunc == "crypto" {
    return getPrice(passedArg)
  } else if recentFunc == "stock" {
    return getStockInfo(passedArg, passedArg1)
  }

  return "Invalid argument passed: " + recentFunc
}
