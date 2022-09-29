package dao

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"how-to-init-sql.DB/dao/internal"
	"how-to-init-sql.DB/dao/model"
)

func TestInsertuser(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		userTest := model.User{Username: "testuser_Name", Password: "test"}
		mock := internal.Mock4Test(t)

		mock.SqlMock.ExpectBegin()
		mock.SqlMock.ExpectExec(regexp.QuoteMeta(
			"INSERT INTO `users` (`username`,`password`,`email`) VALUES (?,?,?)")).
			WithArgs(userTest.Username, userTest.Password, userTest.Email).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.SqlMock.ExpectCommit()
		defer mock.SqlDB.Close()

		user = &userDao{mock.GormDB}
		err := User().Create(&userTest)
		assert.Nil(t, err)
		err = mock.SqlMock.ExpectationsWereMet()
		assert.Nil(t, err)
	})
}
