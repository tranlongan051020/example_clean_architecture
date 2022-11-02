package usecase

import (
	"clean-architecture/internal/user/mock"
	"clean-architecture/internal/user/models"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_GetUsers(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockR := mock.NewMockRepository(ctrl)
	mockUc := NewUserUsecase(mockR)

	var err error

	mockR.EXPECT().GetUsers(context.Background()).Return([]models.User{}, err)

	result, err := mockUc.GetUsers(context.Background())

	require.Equal(t, result, []models.User{})
	require.NoError(t, err)
	require.Nil(t, err)
}

func Test_CreateUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockR := mock.NewMockRepository(ctrl)
	mockUc := NewUserUsecase(mockR)

	userData := models.User{
		Name: "user name 1",
		Email: "user email 1",
	}

	userData.GenerateID()
	userData.PrepareCreate()

	var errR error
	var errU error
	
	mockR.EXPECT().CreateUser(context.Background(), userData).Return(errR)

	errU = mockUc.CreateUser(context.Background(), userData.Name, userData.Email)

	require.Equal(t, errR, errU)
	require.NoError(t, errU)
	require.Nil(t, errU)
}
