package main

import (
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	apiBot := NewApiBot(5 * time.Second)

	telegramBot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	telegramBot.Handle(tb.OnText, func(m *tb.Message) {
		tmnMarket := apiBot.data["TMN"]
		var sb strings.Builder
		sb.WriteString("آخرین قیمت‌ها در بازار والکس:\n")
		// for _, s := range tmnMarket {
		// 	sb.WriteString(s.GetPricesTxt())
		// 	sb.WriteString("\n\n")
		// }
		s := tmnMarket["USDT"]
		sb.WriteString(s.GetPricesTxt())

		telegramBot.Send(m.Sender, sb.String())
	})

	apiBot.Start()
	telegramBot.Start()
}
