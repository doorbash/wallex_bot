package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var interval int
	flag.IntVar(&interval, "i", 5, "fetch interval in seconds")
	flag.Parse()

	if interval < 5 || interval > 60 {
		log.Fatalln("interval must be between 5 and 60")
	}

	apiBot := NewApiBot(time.Duration(interval) * time.Second)

	telegramBot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})

	if err != nil {
		log.Fatalln(err)
	}

	telegramBot.Handle(tb.OnText, func(m *tb.Message) {
		if time.Since(apiBot.lastFetchTime).Minutes() >= 2 {
			telegramBot.Send(m.Sender, "مشکلی رخ داد. لطفا بعدا دوباره امتحان کنید.")
			return
		}
		tmnMarket := apiBot.data["TMN"]
		var sb strings.Builder
		sb.WriteString("*آخرین قیمت‌ها در بازار والکس:*\n")
		// for _, s := range tmnMarket {
		// 	sb.WriteString(s.GetPricesTxt())
		// 	sb.WriteString("\n\n")
		// }
		s := tmnMarket["USDT"]
		sb.WriteString(s.GetPricesWith24chTxt())
		sb.WriteString("\n")
		sb.WriteString(fmt.Sprintf("@%s", os.Getenv("USERNAME")))

		telegramBot.Send(m.Sender, sb.String(), tb.ModeMarkdown)
	})

	apiBot.Start()
	telegramBot.Start()
}
