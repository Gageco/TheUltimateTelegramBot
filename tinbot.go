package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/mnzt/tinder"
  "github.com/bot-api/telegram"
  "github.com/bot-api/telegram/telebot"
  "golang.org/x/net/context"
  "flag"
  "log"
	"bufio"
)



// myLatitude is the current latitude for you.
var myLatitude = float32(43.58)

// myLongitude is the current longitude for you.
var myLongitude = float32(-116.16)

// diff is used to find user age. It finds the difference between two times.
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func getBabe(tin *tinder.Tinder) string {

  recs, err := tin.GetRecommendations(1)
  if err != nil {
    if strings.Contains(err.Error(), "recs timeout") {
      return "Not enough babes right now, wait a bit"
    }
  }

  usr := recs.Results[1]
  userAge, _, _, _, _, _ := diff(usr.Birth, time.Now())
  userName := usr.Name
  //userBio := usr.Bio
  userID := usr.ID
  userPhotos := usr.Photos[0].URL

  textToReturn :=  "Name: " + userName + "\nAge: " + string(userAge)+  "\nBabeID: " + userID + "\n" + userPhotos
  fmt.Println("Displaying info on ",userName)
  return textToReturn
}

func findBabe(id string, tin *tinder.Tinder) string {
  userInfo, err := tin.GetUser(id)
  if err != nil {
    fmt.Println("Error")
  }

  usr := userInfo.Results

  //userAge, _, _, _, _, _ := diff(usr.BirthDate, time.Now())
  userName := usr.Name
  //userBio := usr.Bio
  userID := usr.ID
  userPhotos := usr.Photos[0].URL

	textToReturn :=  "Name: " + userName + "\nBabeID: " + userID + "\n" + userPhotos
  fmt.Println("Finding info ID:",userID, "Name:",userName)
  return textToReturn

}
