package sqlx

import "gorm.io/gorm"

// Seeder is an interface
type Seeder interface {
	Run(db *gorm.DB) error
}

// SQLSeeder is a struct
type SQLSeeder struct {
	seeders []Seeder
}

// NewSQLSeeder is a constructor
func NewSQLSeeder(seeders ...Seeder) Seeder {
	return &SQLSeeder{
		seeders: seeders,
	}
}

// Run is a method
func (s *SQLSeeder) Run(db *gorm.DB) error {
	for _, seeder := range s.seeders {
		if err := seeder.Run(db); err != nil {
			return err
		}
	}
	return nil
}
