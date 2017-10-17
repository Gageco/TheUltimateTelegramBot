package main

import (
	"log"
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
)

type firstMain []firstData

type firstData struct {
	Data    childrenData    `json:"data"`
}

type childrenData []postList

type postList struct {
  Hash    string    `json:"hash"`
  Title   string    `json:"title"`
}

func getBabe() string {
	fmt.Println("Babe Command")
  var redditLink firstData

  subDomain := [4]string {"shorthairedhotties", "PrettyGirls", "beautifulwomen", "SFWRedheads"}
  Domain := "http://imgur.com/"

  randNum := rand.Intn(len(subDomain))
  // randNum = 0
  // log.Println(randNum)
  getDomain := Domain + "r/" + subDomain[randNum] + "/new.json"

  response, err := http.Get(getDomain)
  if err != nil {
    log.Print("40: ")
    log.Println(err)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Print("46: ")
    log.Println(err)
  }
  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &redditLink)
  if err != nil {
    log.Print("53: ")
    log.Println(err)
  }

	randLink := rand.Intn(50)

  linkHash := redditLink.Data[randLink].Hash
  linkTitle := redditLink.Data[randLink].Title

  finalLink := Domain + linkHash + ".jpg"

	stringToReturn := finalLink + "\n" + linkTitle
  return stringToReturn

}
