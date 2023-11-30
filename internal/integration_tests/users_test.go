//go:build integration

package integration_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/convert"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/repository"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/response"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	usersRepo := repository.NewUsersRepo(db)
	usersService := service.NewUsersService(usersRepo)
	router := rest.New(handler.New(usersService, nil))

	testCases := []struct {
		name             string
		prepareDataFunc  func() any
		expectedResponse struct {
			code int
			obj  any
		}
	}{
		{
			name: "valid",
			prepareDataFunc: func() any {
				return map[string]string{
					"phone":    "87718665797",
					"password": "password",
					"email":    "artur.zhamaliev@gmail.com",
				}
			},
			expectedResponse: struct {
				code int
				obj  any
			}{
				code: http.StatusCreated,
				obj: response.SignUp{
					Phone: convert.StringToAddr("87718665797"),
					Email: "artur.zhamaliev@gmail.com",
				},
			},
		},
		{
			name: "invalid email address",
			prepareDataFunc: func() any {
				return map[string]string{
					"phone":    "87718665797",
					"password": "password",
					"email":    "wrong address",
				}
			},
			expectedResponse: struct {
				code int
				obj  any
			}{
				code: http.StatusInternalServerError,
				obj:  nil,
			},
		},
		{
			name:            "invalid request body",
			prepareDataFunc: func() any { return map[string]string{} },
			expectedResponse: struct {
				code int
				obj  any
			}{
				code: http.StatusBadRequest,
				obj:  nil,
			},
		},
		{
			name: "user already exists",
			prepareDataFunc: func() any {
				userEmail := "artur.zhamaliev@gmail.com"
				usersRepo.InsertUser(ctx, model.User{Email: userEmail})

				return map[string]string{
					"phone":    "87718665797",
					"password": "password",
					"email":    userEmail,
				}
			},
			expectedResponse: struct {
				code int
				obj  any
			}{
				code: http.StatusInternalServerError,
				obj:  nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reqBody := tc.prepareDataFunc()

			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(reqBody)
			require.NoError(t, err, "error on encoding request body")

			req, err := http.NewRequest(http.MethodPost, "/api/v1/users", b)
			require.NoError(t, err, "error on sending request")

			router.ServeHTTP(rec, req)
			require.Equal(t, tc.expectedResponse.code, rec.Code)

			if rec.Code != http.StatusCreated {
				return
			}

			var respObj response.SignUp
			err = json.NewDecoder(rec.Body).Decode(&respObj)
			require.NoError(t, err, "error on decoding response object")

			require.Equal(t, tc.expectedResponse.obj.(response.SignUp).Email, respObj.Email)
			require.Equal(t, tc.expectedResponse.obj.(response.SignUp).Phone, respObj.Phone)
		})
	}
}
