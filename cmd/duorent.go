package main

import (
	"context"
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/transport/rest"
	"duorent.ru/internal/transport/rest/routes"
	"duorent.ru/pkg/logger"
	"duorent.ru/pkg/signal"
	"log"
	"os"

	_ "github.com/jackc/pgx"
)

func main() {
	go signal.ListenSignals()

	runMode := os.Getenv("RUN_MODE")
	if runMode == "" {
		runMode = "dev"
	}

	logService, err := logger.NewLogrusLogger()
	if err != nil {
		log.Fatalln(err)
	}

	configService := config.NewConfigService(logService)
	if err := configService.LoadConfig(runMode); err != nil {
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
