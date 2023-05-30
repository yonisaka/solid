package sqlx

import (
	"fmt"
	"github.com/yonisaka/solid/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnection struct {
	cfg *config.Config
}

func NewPostgresConnection(cfg *config.Config) SQLConnectionStrategy {
	return &PostgresConnection{cfg: cfg}
}

func (p *PostgresConnection) Connect() (*gorm.DB, error) {
	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		p.cfg.PostgresConfig.Host,
		p.cfg.PostgresConfig.User,
		p.cfg.PostgresConfig.Password,
		p.cfg.PostgresConfig.DBName,
		p.cfg.PostgresConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *PostgresConnection) Migrate(db *gorm.DB, entities ...interface{}) error {
	if len(entities) == 0 {
		return fmt.Errorf("no entities to migrate")
	}

	return db.AutoMigrate(
		entities...,
	)
}
