package main

import (
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/transport/rest"
	"duorent.ru/internal/transport/rest/routes"
	"duorent.ru/pkg/signal"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	go signal.ListenSignals()

	pool, err := postgres.NewPostgresConnection("", 100)
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := pool.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	router := routes.InitRestRoutes(conn)

	if err := rest.RunNewHTTPServer(8080, router); err != nil {
		log.Fatalln(err)
	}
}
