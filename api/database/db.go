package database

import (
	"github.com/jinzhu/gorm"
	// loading mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/leandroudala/golang_jwt/api/config"
)

// Connect returns a database connection
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
