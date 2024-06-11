package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/memory"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_GetBondsPerUserHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given the query params are valid and there are valid bonds, a status code 200 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=25&page=1", http.NoBody),
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
				_ = memo.Save(context.Background(), bond)

				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
			},
			expectedStatusCode: http.StatusOK,
		},
		"Given the query param 'current_owner_id' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290&limit=10&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'limit' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=-1&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'page' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=15&page=-7", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'page' does not comply with business rules, a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=25&page=0", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given the query param 'limit' does not comply with business rules, a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/user?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=15&page=4", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewUserBondsRetriever(&memo)

				return GetBondsPerUserHandler(getter), nil
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

func Test_GetBondsHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given the query params are valid and there are valid bonds, a status code 200 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=25&page=1", http.NoBody),
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
				_ = memo.Save(context.Background(), bond)

				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
			},
			expectedStatusCode: http.StatusOK,
		},
		"Given the query param 'current_owner_id' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290&limit=10&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'limit' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=-1&page=1", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'page' is invalid, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=15&page=-7", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given the query param 'page' does not comply with business rules, a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=25&page=0", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given the query param 'limit' does not comply with business rules, a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodGet, "/api/v1/bonds/all?current_owner_id=580b87da-e389-4290-acbf-f6191467f401&limit=15&page=4", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewBondMemory()
				getter := logic.NewBondsRetriever(&memo)

				return GetBondsHandler(getter), nil
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
