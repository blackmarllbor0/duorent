package main

import (
	"context"
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository/postgres"
	"duorent.ru/internal/transport/rest"
	"duorent.ru/internal/transport/rest/routes"
	"duorent.ru/pkg/logger"
	run_mode "duorent.ru/pkg/run-mode"
	"duorent.ru/pkg/signal"
	_ "github.com/jackc/pgx"
	"log"
)

func main() {
	go signal.ListenSignals()

	runMode := run_mode.GetAppRunMode()

	logService, err := logger.NewLogrusLogger(runMode)
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
