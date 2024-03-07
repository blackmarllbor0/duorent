package postgres

import (
	"database/sql"
	"duorent.ru/internal/repository"
	"fmt"
	"sync"
)

type pgPool struct {
	pool             *sql.DB
	mutex            *sync.Mutex
	maxCons, numCons uint
}

func NewPostgresConnection(connString string, maxCons uint) (repository.SQLConnection, error) {
	pool, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("pg: opening err: %v", err)
	}

	pool.SetMaxOpenConns(int(maxCons))

	if err := pool.Ping(); err != nil {
		return nil, fmt.Errorf("pg: ping error: %v", err)
	}

	return &pgPool{
		pool:    pool,
		maxCons: maxCons,
		numCons: 0,
		mutex:   &sync.Mutex{},
	}, nil
}

func (p *pgPool) GetConnection() (*sql.DB, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.numCons == p.maxCons {
		return nil, fmt.Errorf("pg: connection pool is full")
	}

	p.numCons++

	return p.pool, nil
}

func (p *pgPool) ReleaseConnection() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.numCons--
}
