//go:build unit

package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/mocks"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/convert"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/response"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name             string
		requestBody      any
		expectedResponse struct {
			code int
			obj  response.CreateUser
		}
	}{
		{
			name: "valid",
			requestBody: map[string]string{
				"phone":    "87718665797",
				"password": "password",
				"email":    "artur.zhamaliev@gmail.com",
			},
			expectedResponse: struct {
				code int
				obj  response.CreateUser
			}{
				code: http.StatusCreated,
				obj: response.CreateUser{
					Phone: convert.StringToAddr("87718665797"),
					Email: "artur.zhamaliev@gmail.com",
				},
			},
		},
	}

	usersService := mocks.NewUsersServiceMock()
	handler := handler.New(usersService)
	router := rest.New(handler)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			err := json.NewEncoder(b).Encode(tc.requestBody)
			require.NoError(t, err, "error on encoding request body")

			req, err := http.NewRequest(http.MethodPost, "/api/v1/users", b)
			require.NoError(t, err, "error on sending request")

			router.ServeHTTP(rec, req)
			require.Equal(t, tc.expectedResponse.code, rec.Code)

			var respObj response.CreateUser
			err = json.NewDecoder(rec.Body).Decode(&respObj)
			require.NoError(t, err, "error on decoding response object")

			require.Equal(t, tc.expectedResponse.obj.Email, respObj.Email)
			require.Equal(t, tc.expectedResponse.obj.Phone, respObj.Phone)
		})
	}
}
