package postgres

import (
	"context"
	"duorent.ru/internal/config"
	"duorent.ru/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type pgxPool struct {
	db  *pgxpool.Config
	ctx context.Context
}

func NewPostgresConnection(cfgService config.ConfigService, ctx context.Context) (repository.SQLConnection, error) {
	db, err := pgxpool.ParseConfig(cfgService.GetDBConfig().Postgres.ConnString)
	if err != nil {
		return nil, fmt.Errorf("pg: failed to create a cfg: %v", err)
	}

	db.MaxConns = int32(cfgService.GetDBConfig().Postgres.MaxCons)
	db.MinConns = int32(cfgService.GetDBConfig().Postgres.MinCons)
	db.MaxConnLifetime = time.Minute * time.Duration(cfgService.GetDBConfig().Postgres.MaxConnLifetimeInMinutes)
	db.MaxConnIdleTime = time.Minute * time.Duration(cfgService.GetDBConfig().Postgres.MaxConnIdleTimeInMinutes)
	db.HealthCheckPeriod = time.Second * time.Duration(cfgService.GetDBConfig().Postgres.HealthCheckPeriodInSeconds)
	db.ConnConfig.ConnectTimeout = time.Second * time.Duration(cfgService.GetDBConfig().Postgres.ConnectTimeoutInSeconds)

	return &pgxPool{db: db, ctx: ctx}, err
}

func (p *pgxPool) GetConnection() (*pgxpool.Conn, error) {
	pool, err := pgxpool.NewWithConfig(p.ctx, p.db)
	if err != nil {
		return nil, fmt.Errorf("pg: error while creating conn to the db: %v", err)
	}

	conn, err := pool.Acquire(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("pg: error while acquiring conn from the db pool: %v", err)
	}

	if err := conn.Ping(p.ctx); err != nil {
		return nil, fmt.Errorf("could not ping db")
	}

	return conn, nil
}
