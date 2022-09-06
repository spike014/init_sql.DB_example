package internal

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"how-to-init-sql.DB/config"
)

var once sync.Once

type sqlDB struct {
	dbConn *gorm.DB
}

var db *sqlDB

func DB() *gorm.DB {
	once.Do(func() {
		db = initDBConn(config.DataSourceName)
	})

	return db.dbConn
}

func initDBConn(dataSourceName string) *sqlDB {
	db, err := createClient(dataSourceName)
	if err != nil {
		// handle error
		// panic? retry? send message to sentry? Whatever...
		return nil
	}
	return &sqlDB{dbConn: db}
}

func createClient(dataSourceName string) (*gorm.DB, error) {
	dbConn, err := gorm.Open(mysql.Open(dataSourceName))
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
