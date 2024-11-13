package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type Storage struct {
	log *slog.Logger
	db  *gorm.DB
}

func New(log *slog.Logger, dsn string) (*Storage, error) {
	const operation = "storage.postgres.New"

	log = log.With(
		slog.String("operation", operation),
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	log.Info("database opened")

	return &Storage{
		log: log,
		db:  db,
	}, nil
}

func (s *Storage) Stop() {
	const operation = "storage.postgres.Stop"

	s.log.With(
		slog.String("operation", operation),
	)

	sqlDB, err := s.db.DB()
	if err != nil {
		s.log.Warn("failed to get *sql.DB instance")
		return
	}

	if err = sqlDB.Close(); err != nil {
		s.log.Warn("close database failed")
		return
	}

	s.log.Info("database closed")
}
