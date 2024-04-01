package sql

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

// Write a connection Manager to handle DB operations make it thread safe

type ConnectionManager struct {
	conn *sql.DB
	mu   sync.Mutex
}

// NewConnectionManager creates a new ConnectionManager instance
func NewConnectionManager(connectionStr string) (*ConnectionManager, error) {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}
	return &ConnectionManager{conn: db}, nil
}

func (cm *ConnectionManager) GetConnection() *sql.DB {
	return cm.conn
}

func (cm *ConnectionManager) Close() error {
	return cm.conn.Close()
}
func (cm *ConnectionManager) Ping() error {
	return cm.conn.Ping()
}
