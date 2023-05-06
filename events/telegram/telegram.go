package telegram

import "anyone-of-the-day/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// хранилище бы бахнуть
}
