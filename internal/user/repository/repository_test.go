package repository

import (
	"clean-architecture/internal/user/models"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	  }), &gorm.Config{})

	repo := NewUserRepository(gormDB)

	t.Run("CreateUser", func(t *testing.T) {
		userData := models.User{
			Name: "user name 01",
			Email: "user email 01",
		}
		userData.GenerateID()
		userData.PrepareCreate()

		rows := sqlmock.NewRows(
			[]string{"id", "name", "email","created_at","updated_at","delete_at"},
		).AddRow(userData.ID, userData.Name, userData.Email,userData.CreatedAt,userData.UpdatedAt,userData.DeleteAt)

		mock.ExpectQuery(createUser).WithArgs(userData.ID, userData.Name, userData.Email,userData.CreatedAt,userData.UpdatedAt,userData.DeleteAt).WillReturnRows(rows)

		err := repo.CreateUser(context.Background(), &userData)

		require.NoError(t, err)
		require.Nil(t, err)
	})
}
