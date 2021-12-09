package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	messageHeader = "<b>آخرین قیمت‌ها در بازار والکس:</b>"
	messageFooter = fmt.Sprintf("@%s", os.Getenv("USERNAME"))
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
		log.Printf("message from %d (@%s): %s\n", m.Sender.ID, m.Sender.Username, m.Text)
		if apiBot.text == "" || time.Since(apiBot.lastFetchTime).Minutes() >= 2 {
			telegramBot.Send(m.Sender, "خطایی رخ داد. لطفا بعدا دوباره امتحان کنید.")
			return
		}
		telegramBot.Send(m.Sender, fmt.Sprintf("%s\n%s%s",
			messageHeader,
			apiBot.text,
			fmt.Sprintf("@%s", os.Getenv("USERNAME"))),
			tb.ModeHTML)
	})

	telegramBot.Handle(tb.OnQuery, func(q *tb.Query) {
		results := make(tb.Results, 0)
		market := apiBot.data["TMN"]
		keys := make([]string, 0)
		for k, _ := range market {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, i := range keys {
			v := market[i]
			ii := strings.ToLower(i)
			if q.Text != "" && !strings.Contains(ii, strings.ToLower(q.Text)) {
				continue
			}
			result := &tb.ArticleResult{
				Title:       v.Symbol,
				Description: fmt.Sprintf("%s %s\n%s %s", v.GetAsk(), v.FaQuoteAsset, v.GetBid(), v.FaQuoteAsset),
				HideURL:     true,
				URL:         fmt.Sprintf("https://raw.githubusercontent.com/spothq/cryptocurrency-icons/master/128/color/%s.png", ii),
				ThumbURL:    fmt.Sprintf("https://raw.githubusercontent.com/spothq/cryptocurrency-icons/master/32/color/%s.png", ii),
			}
			result.SetContent(&tb.InputTextMessageContent{
				Text:      fmt.Sprintf("<b>آخرین قیمت %s در والکس:</b>\n\n%s\n%s", v.FaBaseAsset, v.GetPricesTxt(), messageFooter),
				ParseMode: tb.ModeHTML,
			})
			result.SetResultID(i)
			results = append(results, result)
		}
		if len(results) > 0 {
			err := telegramBot.Answer(q, &tb.QueryResponse{
				Results:   results,
				CacheTime: 5,
			})

			if err != nil {
				log.Println(err)
			}
		}
	})

	apiBot.Start()
	telegramBot.Start()
}
