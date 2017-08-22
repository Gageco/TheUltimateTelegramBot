/*
Bots this includes:
- Tinbot
- TheTime
- Cryptopricebot
- tgStockBot
*/

package main

import (
  "fmt"
  "flag"
  "log"
  "github.com/bot-api/telegram"
  "github.com/bot-api/telegram/telebot"
  "golang.org/x/net/context"
  "bufio"
  "os"
  "github.com/mnzt/tinder"
)

func main() {
  //Getting the config information
  inFile, _ := os.Open("./config")
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
  scanner.Scan()
  scanner.Scan()
  telegramToken := scanner.Text()                                               //Line 2
  scanner.Scan()
  scanner.Scan()
  facebookID := scanner.Text()                                                  //Line 4
  scanner.Scan()
  scanner.Scan()
  facebookToken := scanner.Text()                                               //Line 6
  scanner.Scan()
  scanner.Scan()
  quandlAPI := scanner.Text()                                                   //Line 8
  inFile.Close()

  //Telegram Authentication
	token := flag.String("token", telegramToken, "telegram bot token")
	debug := flag.Bool("debug", false, "show debug information")
	flag.Parse()
	if *token == "" {
		log.Fatal("token flag is required")
	}
	api := telegram.New(*token)
	api.Debug(*debug)
	bot := telebot.NewWithAPI(api)
	bot.Use(telebot.Recover()) // recover if handler panics

	netCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
  fmt.Println("Bot has started working.")

  //Facebook Authentication
  tin := tinder.Init(facebookID, facebookToken)
  err := tin.Auth()
  if err != nil {
    fmt.Println("Auth failed%s", err.Error())
  }

  userInfo, err := tin.GetUser(tin.Me.User.ID)
  if userInfo.Results.Name == "" {
    fmt.Println("Token expired")
  }

  if err != nil {
    fmt.Println("Failed to get your user info:%s",err.Error())
  }
  fmt.Println("Logged into Facebook")

	bot.Use(telebot.Commands(map[string]telebot.Commander{

    //Misc Commands
		"start": telebot.CommandFunc(                                               //start
			func(ctx context.Context, arg string) error {
				api := telebot.GetAPI(ctx)
				update := telebot.GetUpdate(ctx)
				_, err := api.SendMessage(ctx,
					telegram.NewMessagef(update.Chat().ID, startText(),
					))
				return err
			}),

    "help": telebot.CommandFunc(                                               //start
  		func(ctx context.Context, arg string) error {
  			api := telebot.GetAPI(ctx)
  			update := telebot.GetUpdate(ctx)
  			_, err := api.SendMessage(ctx,
  				telegram.NewMessagef(update.Chat().ID, helpCommand(arg),
  				))
  			return err
  		}),

      "kbnt": telebot.CommandFunc(                                               //start
    		func(ctx context.Context, arg string) error {
    			api := telebot.GetAPI(ctx)
    			update := telebot.GetUpdate(ctx)
    			_, err := api.SendMessage(ctx,
    				telegram.NewMessagef(update.Chat().ID, kbntCommand(),
    				))
    			return err
    		}),

    "rimshot": telebot.CommandFunc(                                              //rimshot
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, rimshotCommand(),
          ))
        return err
      }),

    "abe": telebot.CommandFunc(                                              //rimshot
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, abeCommand(),
          ))
        return err
      }),

      //Tinbot Commands
      "babe": telebot.CommandFunc(                                              //babe
  			func(ctx context.Context, arg string) error {
  				api := telebot.GetAPI(ctx)
  				update := telebot.GetUpdate(ctx)
  				_, err := api.SendMessage(ctx,
  					telegram.NewMessagef(update.Chat().ID, getBabe(tin),
  					))
  				return err
  			}),

      "findbabe": telebot.CommandFunc(                                          //findbabe
        func(ctx context.Context, arg string) error {
          api := telebot.GetAPI(ctx)
          update := telebot.GetUpdate(ctx)
          _, err := api.SendMessage(ctx,
            telegram.NewMessagef(update.Chat().ID, findBabe(arg, tin),
            ))
          return err
        }),

      //Cryptobot Commands
      "coin": telebot.CommandFunc(                                             //coin
    		func(ctx context.Context, arg string) error {
    			api := telebot.GetAPI(ctx)
    			update := telebot.GetUpdate(ctx)
    			_, err := api.SendMessage(ctx,
    				telegram.NewMessagef(update.Chat().ID, getPrice(arg),
    				))
    			return err
    		}),

      //Stockbot Commands
      "stock": telebot.CommandFunc(                                             //stock
        func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, getStockInfo(arg, quandlAPI),
          ))
        return err
      }),

      //Thetimebot Commands
      "whattimeisit": telebot.CommandFunc(                                      //whattimeisit
    		func(ctx context.Context, arg string) error {
    			api := telebot.GetAPI(ctx)
    			update := telebot.GetUpdate(ctx)
    			_, err := api.SendMessage(ctx,
    				telegram.NewMessagef(update.Chat().ID, getTheTime(),
    			))
  			return err
  		}),


		"": telebot.CommandFunc(                                                    //not valid
			func(ctx context.Context, arg string) error {

				api := telebot.GetAPI(ctx)
				update := telebot.GetUpdate(ctx)
				//command, arg := update.Message.Command()
				_, err := api.SendMessage(ctx,
					telegram.NewMessagef(update.Chat().ID, invalidInput(arg),
					))
				return err
			}),
	}))

	err = bot.Serve(netCtx)
	if err != nil {
		log.Fatal(err)
	}




}
