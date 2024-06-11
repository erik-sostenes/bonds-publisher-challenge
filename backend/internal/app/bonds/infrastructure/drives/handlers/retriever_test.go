package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_GetBondsPerUserHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given the query params are valid, a status code 200 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=10&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return GetBondsPerUserHandler(), nil
			},
			expectedStatusCode: http.StatusOK,
		},
		"Given the query param 'current_owner_id' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290&limit=10&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return GetBondsPerUserHandler(), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'limit' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=-1&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return GetBondsPerUserHandler(), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'page' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=15&page=-7", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return GetBondsPerUserHandler(), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			req := tsc.request
			req.Header.Set("Content-type", "application/json; charset=utf-8")

			resp := httptest.NewRecorder()

			h, err := tsc.handlerFunc()
			if err != nil {
				t.Fatal(err)
			}

			BondErrorHandler(h).ServeHTTP(resp, req)

			if tsc.expectedStatusCode != resp.Code {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", tsc.expectedStatusCode, resp.Code)
			}
		})
	}
}
