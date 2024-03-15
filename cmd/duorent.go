package main

import (
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/transport/rest"
	"duorent.ru/internal/transport/rest/routes"
	"duorent.ru/pkg/signal"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	go signal.ListenSignals()

	configService := config.NewConfigService()
	if err := configService.LoadConfig(); err != nil {
		log.Fatalln(err)
	}

	conn, err := postgres.NewPostgresConnection(configService.GetDBConfig().Postgres.ConnString, 100)
	if err != nil {
		log.Fatalln(err)
	}

	router := routes.InitRestRoutes(conn)

	if err := rest.RunNewHTTPServer(configService.GetServerConfig().Port, router); err != nil {
		log.Fatalln(err)
	}
}
