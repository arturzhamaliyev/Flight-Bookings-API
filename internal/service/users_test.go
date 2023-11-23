//go:build unit

package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/mocks"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/convert"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name            string
		prepareDataFunc func() (*service.Users, model.User)
		expectedError   error
	}{
		{
			name: "valid",
			prepareDataFunc: func() (*service.Users, model.User) {
				return service.NewUsersService(mocks.NewUsersRepoMock()), model.User{
					ID:        uuid.New(),
					Phone:     convert.StringToAddr("87718665797"),
					Email:     "artur.zhamaliev@gmail.com",
					Password:  "password",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			},
			expectedError: nil,
		},
		{
			name: "invalid email address",
			prepareDataFunc: func() (*service.Users, model.User) {
				return service.NewUsersService(mocks.NewUsersRepoMock()), model.User{
					ID:        uuid.New(),
					Phone:     convert.StringToAddr("87718665797"),
					Email:     "wrong address",
					Password:  "password",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			},
			expectedError: service.ErrInvalidEmailAddress,
		},
		{
			name: "user already exists",
			prepareDataFunc: func() (*service.Users, model.User) {
				usersRepo := mocks.NewUsersRepoMock()
				usersService := service.NewUsersService(usersRepo)

				userEmail := "artur.zhamaliev@gmail.com"
				usersRepo.InsertUser(ctx, model.User{Email: userEmail})

				return usersService, model.User{
					ID:        uuid.New(),
					Phone:     convert.StringToAddr("87718665797"),
					Email:     userEmail,
					Password:  "password",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
			},
			expectedError: service.ErrUserExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			usersService, user := tc.prepareDataFunc()

			err := usersService.CreateUser(ctx, user)
			if !errors.Is(err, tc.expectedError) {
				require.NoError(t, err)
			}
		})
	}
}
