package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to BoardGame Search")
	skip := flag.Int("skip", 0, "Skip the no of search results")
	limit := flag.Int("limit", 20, "Limit the no of search results")
	clientId := flag.String("client-id", os.Getenv("CLIENT_ID"), "client-id for the board-game")
	search := flag.String("search", "", "search key")
	output := flag.String("output", "text", "preferred output method. text,json")
	flag.Parse()

	Validate(*clientId, "client-id cannot be empty")
	Validate(*search, "search key cannot be empty")

	bga := NewBoardGameAtlas(*clientId)
	payload, err := bga.search(*search, *limit, *skip)
	if err != nil {
		log.Fatalln(err)
	}

	switch *output {
	case "text":
		PrintText(&payload.Games)
	case "json":
		PrintJson(&payload.Games)
	default:
		PrintText(&payload.Games)
	}
}
