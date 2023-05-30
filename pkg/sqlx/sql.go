package sqlx

import "gorm.io/gorm"

type SQLConnection struct {
	strategy SQLConnectionStrategy
}

// NewSQL is a constructor
func NewSQL(strategy SQLConnectionStrategy) SQLConnectionStrategy {
	return &SQLConnection{
		strategy: strategy,
	}
}

// Connect is a method
func (s *SQLConnection) Connect() (*gorm.DB, error) {
	return s.strategy.Connect()
}

// Migrate is a method
func (s *SQLConnection) Migrate(db *gorm.DB, entities ...interface{}) error {
	return s.strategy.Migrate(db, entities)
}
