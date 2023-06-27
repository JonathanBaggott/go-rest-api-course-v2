package db

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// MigrateDB performs the database migration using the golang-migrate library.
func (d *Database) MigrateDB() error {
	fmt.Println("migrating our database")

	// Create a PostgreSQL driver instance.
	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %w", err)
	}

	// Create a new migration instance with the PostgreSQL driver.
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Apply pending migrations to the database.
	if err := m.Up(); err != nil {
		// If there are no pending migrations, migrate.ErrNoChange is returned, and the migration process continues.
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("could not run up migrations: %w", err)
		}
	}
	fmt.Println("successfully migrated the database")

	// nil is returned to indicate that the migration was successful.
	return nil
}
