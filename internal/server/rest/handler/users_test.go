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
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/helper"
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
			obj  any
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
				obj  any
			}{
				code: http.StatusCreated,
				obj: response.SignUp{
					Phone: helper.StringToAddr("87718665797"),
					Email: "artur.zhamaliev@gmail.com",
				},
			},
		},
		{
			name:        "invalid request body",
			requestBody: "invalid body",
			expectedResponse: struct {
				code int
				obj  any
			}{
				code: http.StatusBadRequest,
				obj:  nil,
			},
		},
	}

	router := rest.New(handler.New(mocks.NewUsersServiceMock(), nil))

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
