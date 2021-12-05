package main

import (
	"polls/api/server"
	"polls/internal/db"
	"polls/transport"
)

func main() {
	dbClient := db.NewClient()

	polls := db.NewPolls(dbClient)

	s := server.New(polls)

	amqper := transport.New()
	defer amqper.Stop()

	go amqper.Consume()

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
