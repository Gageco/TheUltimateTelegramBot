package main

import "fmt"

func startText() string {
  var text string
  fmt.Println("Start Text")
  text = ("The Ultimate Bot For Telegram!!!\nIncorporates:\n- Tinbot\n- Cryptoprices\n- Stockbot\n- TheTime\nThe code for this is opensource and can be found at https://github.com/gageco/TheUltimateTelegramBot.\n\nIf you want to buy me a coffee or you liked this bot, please feel free to send me some coin at the following:\nEthereum: 0x2A7D65d4F6C148c7dfeEc76836E8c5EE02dc5f83\nBitcoin: 1NcWtC3McnZwebyhkruadg7zaWKNYAZgmW\n")

  return text
}

func invalidInput(arg string) string {
  fmt.Println("Invalid Input:",arg)
  return "Invalid Input"
}

func rimshotCommand() string {
  fmt.Println("Rimshot command")
  return "https://giphy.com/gifs/c8bJDVz7i9KRW/html5"
}

func helpCommand(arg string) string {
  if arg == "stock" {
    return "/stock [stock symbol]\nExample: /stock goog"
  } else if arg == "coin" {
    return "/coin [coin symbol]\nExample: /coin eth"
  } else if arg == "findbabe" {
    return "/findbabe [babeid]\nExample: /findbabe asdf35FJLDK10af"
  } else {
    return "Cryptocoin: /coin eth\nStocks: /stock goog\nThe Time: /whattimeisit\nBabe: /babe\nFind Babe: /findbabe [BabeID]\nRimshot: /rimshot\n\nIf you want to buy me a coffee or you liked this bot, please feel free to send me some coin at the following:\nEthereum: 0x2A7D65d4F6C148c7dfeEc76836E8c5EE02dc5f83\nBitcoin: 1NcWtC3McnZwebyhkruadg7zaWKNYAZgmW\n"
  }

}

func abeCommand() string {
  fmt.Println("Abe Command")
  return "http://img.quotery.com/pictures/2013/02/abraham-lincoln.jpg"
}
