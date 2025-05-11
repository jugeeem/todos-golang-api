package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigration は指定されたデータベースURLに対してマイグレーションを実行します
func RunMigration(databaseURL string) error {
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migration: %w", err)
	}
	log.Println("Migration completed successfully")

	return nil
}

// RollbackMigration は最後に適用されたマイグレーションをロールバックします
func RollbackMigration(databaseURL string) error {
	m, err := migrate.New(
		"file://migrations",
		databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Migration rollback completed successfully")
	return nil
}
