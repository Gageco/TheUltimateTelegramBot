package main

import (
	"fmt"
	"math/rand"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"bytes"
)

// var Blacklist[50]string

type firstMain []firstData

type firstData struct {
	Data    childrenData    `json:"data"`
}

type childrenData []postList

type postList struct {
  Hash    string    `json:"hash"`
  Title   string    `json:"title"`
}

func babeRetry(babe string) string {
	fmt.Println("babeRetrying")
	if babe == "gbabe" {
		return getBabe()
	}
	//  else if babe == "fbabe" {
	// 	return getFemaleBabe()
	// } else if babe == "mbabe" {
	// 	return getMaleBabe()
	// }
	return getBabe()
}

func getBabe() string {
	subDomain := [7]string {"shorthairedhotties", "PrettyGirls", "beautifulwomen", "SFWRedheads", "gentlemanboners", "sexysfw", "ladyboners"}
	Domain := "http://imgur.com/"

	randNum := rand.Intn(len(subDomain))
	getDomain := Domain + "r/" + subDomain[randNum] + "/new.json"

	return getPicture(getDomain, Domain, "gbabe")
}

func getPicture(getDomain string, Domain string, babeType string) string {
  var redditLink firstData

  response, err := http.Get(getDomain)
  if err != nil {
    fmt.Print("76: ",babeType)
    fmt.Println(err)
    return babeRetry(babeType)
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Print("83: ",babeType)
    fmt.Println(err)
    return babeRetry(babeType)
  }
  data := bytes.TrimSpace(body)
  data = bytes.TrimPrefix(data, []byte("// "))
  err = json.Unmarshal(data, &redditLink)
  if err != nil {
    fmt.Print("91: ",babeType)
    fmt.Println(err)
    return babeRetry(babeType)
  }

  randLink := rand.Intn(len(redditLink.Data))
  linkHash := redditLink.Data[randLink].Hash
  linkTitle := redditLink.Data[randLink].Title
  finalLink := Domain + linkHash + ".jpg"

  // matchFound := false

  // for i:=0; i < len(Blacklist); i++ {
  //   time.Sleep(time.Millisecond * 1)
  //   // fmt.Println(linkHash)
  //   // fmt.Println(Blacklist[i])
  //   if linkHash == Blacklist[i] {
  //     matchFound = true
  //     // fmt.Println(matchFound)
  //     break
  //   } else if Blacklist[i] == "" {
  //     break
  //   }
  //   // fmt.Println(matchFound, i)
  // }
	//
  // if matchFound {
  //   return babeRetry(babeType)
  // }
  fmt.Println("Babe: ", linkHash)
  return finalLink + "\nTitle: " + linkTitle + "\nID: " + linkHash
}
