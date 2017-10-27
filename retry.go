package main

import "fmt"

var numberOfRetrys int
var previousFunc string

// func initializeRetry() {
//   numberOfRetrys = 0
// }

func retry(recentFunc string, passedArg string) string {
  fmt.Println("Retrying")
  if previousFunc != recentFunc {
    numberOfRetrys = 0
  } else if recentFunc == "crypto" {
    numberOfRetrys += 1
    if numberOfRetrys <= 5 {
      return getPrice(passedArg)
    }
    numberOfRetrys = 0
    return "Something went wrong with crypto functions"
  }

  return "Invalid argument passed: " + recentFunc
}
