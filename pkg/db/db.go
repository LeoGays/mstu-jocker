// Package db содержит объекты и функции для работы с реляционными базами данных.
package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// Connect возвращает открытое соединение с БД и проверяет его. Если во время проверки произошла ошибка, вернет error.
func Connect(cfg ConnectionDescriber) (*sql.DB, error) {
	db, err := sql.Open(cfg.DriverName(), cfg.DataSourceName())
	if err != nil {
		return nil, fmt.Errorf("failed to open %s connection (%v): %w", cfg.DriverName(), cfg, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't check db connection (%v): %w", cfg, err)
	}

	return db, nil
}

// DefaultConnect возвращает открытое соединение с БД и проверяет его. Если во время проверки произошла ошибка, вернет error.
// Данный метод можно использовать для создания или удаления БД, когда инициализируешь тестовое окружение.
func DefaultConnect(cfg ConnectionDescriber) (*sql.DB, error) {
	db, err := sql.Open(cfg.DriverName(), cfg.DefaultDataSourceName())
	if err != nil {
		return nil, fmt.Errorf("failed to open %s connection (%v): %w", cfg.DriverName(), cfg, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't check db connection (%v): %w", cfg, err)
	}

	return db, nil
}

// LoadFixtures загружает фикстуры из указанных файлов и возвращает id созданных записей.
func LoadFixtures(conn *sql.DB, fileNames ...string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, 0, len(fileNames))

	for _, file := range fileNames {
		script, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file (%s): %w", file, err)
		}

		rows, err := conn.Query(string(script))
		if err != nil {
			return nil, fmt.Errorf("can't query db: %w", err)
		}

		for rows.Next() {
			var id string
			if err = rows.Scan(&id); err != nil {
				return nil, fmt.Errorf("can't read query results: %w", err)
			}

			guid, err := uuid.Parse(id)
			if err != nil {
				return nil, fmt.Errorf("id value is incorrect: %w", err)
			}

			result = append(result, guid)
		}
	}

	return result, nil
}
