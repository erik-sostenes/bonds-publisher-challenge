package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/memory"
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
				memo := memory.NewBondMemory()
				bondCtr := logic.NewBondCreator(&memo)

				return PostBondHandler(bondCtr), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given an invalid non-existing bond a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				bondCtr := logic.NewBondCreator(&memo)

				return PostBondHandler(bondCtr), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given the name of an invalid bond, the status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Ultra Secure Global Investment Corporate Bond 2045",
					"quantity_sale": 1,
					"sales_price": 400.0000,
					"creator_user_id":"580b87da-e389-4290-acbf-f6191467f401",
					"current_owner_id":"580b87da-e389-4290-acbf-f6191467f401"
				}`,
			)),

			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				bondCtr := logic.NewBondCreator(&memo)

				return PostBondHandler(bondCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the quantity sale of an invalid bond, the status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Global Secure Corporate Bond 2024",
					"quantity_sale": 10001,
					"sales_price": 400.0000,
					"creator_user_id":"580b87da-e389-4290-acbf-f6191467f401",
					"current_owner_id":"580b87da-e389-4290-acbf-f6191467f401"
				}`,
			)),

			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				bondCtr := logic.NewBondCreator(&memo)

				return PostBondHandler(bondCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the sales price of an invalid bond, the status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/bonds/create", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Global Secure Corporate Bond 2024",
					"quantity_sale": 10000,
					"sales_price": 100000001.0000,
					"creator_user_id":"580b87da-e389-4290-acbf-f6191467f401",
					"current_owner_id":"580b87da-e389-4290-acbf-f6191467f401"
				}`,
			)),

			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				bondCtr := logic.NewBondCreator(&memo)

				return PostBondHandler(bondCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given an existing valid bond, a 400 status code is expected": {
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
				bondRequest := BondRequest{
					ID:             "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
					CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
					CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
				}

				bond, err := bondRequest.toBusiness()
				if err != nil {
					return nil, err
				}

				memo := memory.NewBondMemory()

				// save in memory a new Bond
				memo.Save(context.Background(), bond)

				return PostBondHandler(logic.NewBondCreator(&memo)), nil
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
