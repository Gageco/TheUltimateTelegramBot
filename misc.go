package main

func startText() string {
  var text string
  fmt.Println("Start Text")
  text = ("The Ultimate Bot For Telegram!!!\nIncorporates:\n- Tinbot\n- Cryptoprices\n- Stockbot\n- TheTime\nThe commands vary between them but now you wont need to have like 4 bots doing different things.\nThe code for this is opensource and can be found at https://github.com/gageco/TheUltimateTelegramBot.\nIf you want to buy me a coffee or you liked this bot, please feel free to send me some coin at the following:\nEthereum: 0x2A7D65d4F6C148c7dfeEc76836E8c5EE02dc5f83\nBitcoin: 1NcWtC3McnZwebyhkruadg7zaWKNYAZgmW\n")

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
