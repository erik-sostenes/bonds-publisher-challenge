package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_PostBondHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing bond a status code 201 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Global Secure Corporate Bond 2024",
					"quantity_sale": 1,
					"sales_price": 400.0000,
					"creator_user_id":"580b87da-e389-4290-acbf-f6191467f401",
					"current_owner_id":"580b87da-e389-4290-acbf-f6191467f401"
				}`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return PostBondHandler(), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing bond a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Global Secure Corporate Bond 2024",
					"quantity_sale": 1.5,
					"sales_price": 400.0000,
					"creator_user_id":"580b87da-e389-4290-acbf-f6191467f401",
					"current_owner_id":"580b87da-e389-4290-acbf-f6191467f401"
				}`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return PostBondHandler(), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
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
