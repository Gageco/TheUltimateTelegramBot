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
)

func main() {
  //Getting the config information
  inFile, _ := os.Open("./config")
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
  scanner.Scan()
  scanner.Scan()
  telegramToken := scanner.Text()                                               //Line 2                                               //Line 6
  scanner.Scan()
  scanner.Scan()
  stockAPI := scanner.Text()                                                   //Line 8
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

    "help": telebot.CommandFunc(                                                //help
  		func(ctx context.Context, arg string) error {
  			api := telebot.GetAPI(ctx)
  			update := telebot.GetUpdate(ctx)
  			_, err := api.SendMessage(ctx,
  				telegram.NewMessagef(update.Chat().ID, helpCommand(arg),
  				))
  			return err
  		}),

    "kbnt": telebot.CommandFunc(                                                //kbnt
    	func(ctx context.Context, arg string) error {
    		api := telebot.GetAPI(ctx)
    		update := telebot.GetUpdate(ctx)
    		_, err := api.SendMessage(ctx,
    			telegram.NewMessagef(update.Chat().ID, kbntCommand(),
    			))
    		return err
    	}),

    "bees": telebot.CommandFunc(                                                //bees
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, beeCommand(),
            ))
        return err
      }),

    "bees2": telebot.CommandFunc(                                               //bees2
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, bee2Command(),
            ))
        return err
      }),

    "rimshot": telebot.CommandFunc(                                             //rimshot
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, rimshotCommand(),
          ))
        return err
      }),

    "abe": telebot.CommandFunc(                                                 //abe
      func(ctx context.Context, arg string) error {
        api := telebot.GetAPI(ctx)
        update := telebot.GetUpdate(ctx)
        _, err := api.SendMessage(ctx,
          telegram.NewMessagef(update.Chat().ID, abeCommand(),
          ))
        return err
      }),

    "babe": telebot.CommandFunc( //NEEDSRETRY                                                //babe
  		func(ctx context.Context, arg string) error {
  			api := telebot.GetAPI(ctx)
  			update := telebot.GetUpdate(ctx)
  			_, err := api.SendMessage(ctx,
  				telegram.NewMessagef(update.Chat().ID, getBabe(),
  				))
  			return err
  		}),

    "coin": telebot.CommandFunc(//NEEDSRETRY                                                //coin
    	func(ctx context.Context, arg string) error {
    		api := telebot.GetAPI(ctx)
    		update := telebot.GetUpdate(ctx)
    		_, err := api.SendMessage(ctx,
    			telegram.NewMessagef(update.Chat().ID, getPrice(arg),
    			))
    		return err
    	}),

    "stock": telebot.CommandFunc(                                               //stock
      func(ctx context.Context, arg string) error {
      api := telebot.GetAPI(ctx)
      update := telebot.GetUpdate(ctx)
      _, err := api.SendMessage(ctx,
        telegram.NewMessagef(update.Chat().ID, getStockInfo(arg, stockAPI),
        ))
      return err
    }),

    "whattimeisit": telebot.CommandFunc(                                        //whattimeisit
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

	err := bot.Serve(netCtx)
	if err != nil {
		log.Fatal(err)
	}




}
