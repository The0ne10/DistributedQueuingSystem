package postgres

import (
	"DistributedQueueSystem/internal/config"
	"database/sql"
	"fmt"
	"log/slog"
)

type Storage struct {
	db *sql.DB
}

func New(
	cfg config.Config,
	log *slog.Logger,
) (*Storage, error) {
	const op = "storage.Postgres.New"

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Database)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Connected to PostgreSQL")
	return &Storage{db: db}, nil
}
