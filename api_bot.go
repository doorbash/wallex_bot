package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type ApiBot struct {
	interval      time.Duration
	stop          chan struct{}
	data          map[string]Market
	lastFetchTime time.Time
	text          string
}

func (a *ApiBot) updateText() {
	tmnMarket := a.data["TMN"]
	var sb strings.Builder
	sb.WriteString("<b>آخرین قیمت‌ها در بازار والکس:</b>\n")
	// for _, s := range tmnMarket {
	// 	sb.WriteString(s.GetPricesTxt())
	// 	sb.WriteString("\n\n")
	// }
	s := tmnMarket["USDT"]
	sb.WriteString(s.GetPricesWith24chTxt())
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("@%s", os.Getenv("USERNAME")))
	a.text = sb.String()
}

func (a *ApiBot) fetch() {
	log.Println("Api: get()")
	var err error
	a.data, err = GetMarkets()
	if err != nil {
		log.Println(err)
		return
	}
	a.lastFetchTime = time.Now()
	a.updateText()
}

func (a *ApiBot) run() {
	log.Println("Api: run()")
	for {
		select {
		case <-time.Tick(a.interval):
			a.fetch()
		case <-a.stop:
			log.Println("returning run")
			return
		}
	}
}

func (a *ApiBot) Start() {
	log.Println("Api: Start()")
	a.Stop()
	a.stop = make(chan struct{})
	go a.run()
}

func (a *ApiBot) Stop() {
	log.Println("Api: Stop()")
	if a.stop != nil {
		close(a.stop)
	}
}

func NewApiBot(interval time.Duration) *ApiBot {
	return &ApiBot{
		interval: interval,
	}
}
