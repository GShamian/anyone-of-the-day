package main

import (
	"anyone-of-the-day/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"токен для получения доступа к телеграму",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("токен не задан")
	}

	return *token
}
