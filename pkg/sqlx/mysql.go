package sqlx

import (
	"fmt"
	"github.com/yonisaka/solid/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnection struct {
	cfg *config.Config
}

func NewMySQLConnection(cfg *config.Config) SQLConnectionStrategy {
	return &MySQLConnection{cfg: cfg}
}

func (m *MySQLConnection) Connect() (*gorm.DB, error) {
	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.cfg.MysqlConfig.User,
		m.cfg.MysqlConfig.Password,
		m.cfg.MysqlConfig.Host,
		m.cfg.MysqlConfig.Port,
		m.cfg.MysqlConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *MySQLConnection) Migrate(db *gorm.DB, entities ...interface{}) error {
	if len(entities) == 0 {
		return fmt.Errorf("no entities to migrate")
	}

	return db.AutoMigrate(
		entities...,
	)
}
