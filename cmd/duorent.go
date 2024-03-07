package main

import (
	"duorent.ru/internal/repository/postgres"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	_, err := postgres.NewPostgresConnection("", 100)
	if err != nil {
		log.Fatalln(err)
	}
}
