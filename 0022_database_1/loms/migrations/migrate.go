package db

import (
	"embed"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

// Иногда делают через отдельный контейнер

// Написана на Go
// Поддерживает двунаправленные миграции
// Поддерживает несколько sql диалектов
// Применяет только то, что нужно применять, используя версии
// Можно пользоваться прямо из кода

//go:embed migrations/*.sql
var embedMigrations embed.FS

func SetupPostgres(pool *pgxpool.Pool, logger *zap.Logger) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("can not set dialect in goose", zap.Error(err))
		os.Exit(-1)
	}

	db := stdlib.OpenDBFromPool(pool)
	if err := goose.Up(db, "migrations"); err != nil {
		logger.Error("can not setup migrations", zap.Error(err))
		os.Exit(-1)
	}
}
