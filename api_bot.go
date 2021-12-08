package main

import (
	"log"
	"time"
)

type ApiBot struct {
	interval      time.Duration
	stop          chan struct{}
	data          map[string]Market
	lastFetchTime time.Time
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
