package main

import "fmt"

var numberOfRetrys int
var previousFunc string

// func initializeRetry() {
//   numberOfRetrys = 0
// }

func retry(recentFunc string, passedArg string) string {
  fmt.Println("Retrying")
  if recentFunc == "crypto" {
    numberOfRetrys += 1
    fmt.Println(numberOfRetrys)
    return getPrice(passedArg)
  }

  return "Invalid argument passed: " + recentFunc
}
