package internal

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbmock struct {
	SqlDB   *sql.DB
	GormDB  *gorm.DB
	SqlMock sqlmock.Sqlmock
}

// Mock4Test return dbmock for sql tests.
// Do not use it in production.
func Mock4Test(t *testing.T) *dbmock {
	t.Helper()
	dbMock, mock, err := sqlmock.New()
	assert.Nil(t, err)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	}))
	assert.Nil(t, err)
	return &dbmock{dbMock, gormDB, mock}
}
