package main

import (
	"context"
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/transport/rest"
	"duorent.ru/internal/transport/rest/routes"
	"duorent.ru/pkg/signal"
	"log"

	_ "github.com/jackc/pgx"
)

func main() {
	go signal.ListenSignals()

	configService := config.NewConfigService()
	if err := configService.LoadConfig(); err != nil {
		log.Fatalln(err)
	}

	// todo: decide something with the ctx.
	conn, err := postgres.NewPostgresConnection(configService, context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	router := routes.InitRestRoutes(conn, configService)

	if err := rest.RunNewHTTPServer(configService.GetServerConfig().Port, router); err != nil {
		log.Fatalln(err)
	}
}
