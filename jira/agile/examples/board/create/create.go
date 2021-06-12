package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/ctreminiom/go-atlassian/jira/agile"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	newBoard := &agile.BoardPayloadScheme{
		Name:     "DUMMY Board Name",
		Type:     "scrum", //scrum or kanban
		FilterID: 10016,

		// Omit the Location if you want to the board to yourself (location)
		Location: &agile.BoardPayloadLocationScheme{
			ProjectKeyOrID: "KP",
			Type:           "project",
		},
	}

	board, response, err := atlassian.Agile.Board.Create(context.Background(), newBoard)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(board)
}