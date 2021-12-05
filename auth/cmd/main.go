package main

import (
	"auth/api/server"
	"auth/internal/db"
	"auth/transport"
)

func main() {
	dbClient := db.NewClient()

	amqper := transport.New()
	defer amqper.Stop()

	users := db.NewUsers(dbClient)

	s := server.New(users, amqper)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
