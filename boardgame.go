package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/http"
)

const BoardGameAtlasSearchUrl = "https://api.boardgameatlas.com/api/search" // ?name=Catan&client_id=Iq3EPvFHTI

type Game struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Price       string `json:"price"`
	URL         string `json:"url"`
	ImageUrl    string `json:"image_url"`
}

type Payload struct {
	Games []Game `json:"games"`
	Count int    `json:"count"`
}

type BoardGameAtlas struct {
	clientId string
}

func NewBoardGameAtlas(clientId string) *BoardGameAtlas {
	bga := BoardGameAtlas{clientId}
	return &bga
}

func (bga *BoardGameAtlas) search(searchKey string, limit int, skip int) (*Payload, error) {
	req, err := http.NewRequest(http.MethodGet, BoardGameAtlasSearchUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create HTTP request with URL %v", err)
	}
	req.Header.Add("Accept", "application/json")
	query := req.URL.Query()
	query.Set("name", searchKey)
	query.Set("client_id", bga.clientId)
	query.Set("limit", fmt.Sprintf("%d", limit))
	query.Set("skip", fmt.Sprintf("%d", skip))
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error %v", err)
	}
	var payload Payload
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshall payload %v", err)
	}
	return &payload, nil
}

func PrintText(games *[]Game) {
	titleDecorator := color.New(color.Bold).Add(color.BgGreen).Add(color.FgHiYellow).SprintFunc()
	descDecorator := color.New(color.Italic).Add(color.BgHiBlack).Add(color.FgHiWhite).SprintFunc()
	for _, game := range *games {
		fmt.Printf(titleDecorator(fmt.Sprintf("Title: %v \n", game.Name)))
		fmt.Printf("%v %v \n\n", titleDecorator("Description: "), descDecorator(game.Description))
	}
}

func PrintJson(games *[]Game) {
	gameBytes, err := json.Marshal(*games)
	if err != nil {
		log.Fatalln("Unable to Marshal the games slice", err)
	}
	fmt.Println(string(gameBytes))
}
