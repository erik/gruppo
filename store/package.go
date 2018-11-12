package store

import (
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
	_ "github.com/mattn/go-sqlite3"
)

type Configuration struct {
	Addr string
	DB   int
}

type RedisStore struct {
	db *redis.Client
}

type SqliteStore struct {
	db *sql.DB
}

var (
	NotFound = errors.New("item not found")
)

type Store interface {
	// TODO: define this as a generic interface
}
