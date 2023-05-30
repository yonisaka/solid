package sqlx

import "gorm.io/gorm"

// SQLConnectionStrategy is an interface
type SQLConnectionStrategy interface {
	Connect() (*gorm.DB, error)
	Migrate(db *gorm.DB, entities ...interface{}) error
}
