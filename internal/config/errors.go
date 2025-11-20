package config

import (
	"errors"
)

var (
	ErrPostgresHost = errors.New("POSTGRES_HOST is required")
	ErrPostgresUser = errors.New("POSTGRES_USER is required")
	ErrPostgresDB   = errors.New("POSTGRES_DB is required")
)
