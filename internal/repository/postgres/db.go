package postgres

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a small wrapper around the GORM database handle.
type DB struct {
	Gorm *gorm.DB // Main GORM connection object
}

// Connect opens a PostgreSQL connection using a DATABASE_URL-style string.
// It also sets simple connection pool limits to avoid too many open connections.
func Connect(databaseURL string) (*DB, error) {
	// Create a new GORM connection using the Postgres driver.
	gormDB, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Get the underlying *sql.DB to configure pooling.
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// Pool settings (simple, safe defaults for a small app).
	sqlDB.SetMaxOpenConns(10)                  // Max total open connections
	sqlDB.SetMaxIdleConns(5)                   // Max idle connections kept
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Recycle connections periodically

	return &DB{Gorm: gormDB}, nil
}

func (d *DB) Ping() error {
	sqlDB, err := d.Gorm.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
