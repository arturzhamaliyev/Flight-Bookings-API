//go:build unit

package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/mocks"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/convert"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		user          model.User
		expectedError error
	}{
		{
			name: "valid",
			user: model.User{
				ID:        uuid.New(),
				Phone:     convert.StringToAddr("87718665797"),
				Email:     "artur.zhamaliev@gmail.com",
				Password:  "password",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name: "invalid email address",
			user: model.User{
				ID:        uuid.New(),
				Phone:     convert.StringToAddr("87718665797"),
				Email:     "wrong address",
				Password:  "password",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: service.ErrInvalidEmailAddress,
		},
		{
			name: "user already exists",
			user: model.User{
				ID:        uuid.New(),
				Phone:     convert.StringToAddr("87718665797"),
				Email:     "artur.zhamaliev@gmail.com",
				Password:  "password",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: service.ErrUserExists,
		},
	}

	ctx := context.Background()
	usersRepo := mocks.NewUsersMockRepo()
	usersService := service.NewUsersService(usersRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := usersService.CreateUser(ctx, tc.user)
			if !errors.Is(err, tc.expectedError) {
				t.Error(err)
			}
		})
	}
}
