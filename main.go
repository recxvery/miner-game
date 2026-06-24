package main

import (
	"context"
	"log"
	"miner-game/miner"
	"miner-game/webhttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mine := miner.NewMine(ctx, cancel)
	handlers := webhttp.NewHTTPHandlers(mine)
	server := webhttp.NewHTTPServer(*handlers)

	go mine.BackgroundMine(ctx)
	go mine.Collector(ctx)
	err := server.StartServer()

	if err != nil {
		log.Fatal(err)
	}
}
