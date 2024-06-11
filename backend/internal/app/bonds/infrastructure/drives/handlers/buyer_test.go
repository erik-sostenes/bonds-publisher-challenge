package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_PutBondBuyerHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request   *http.Request
		urlParams struct {
			bondId,
			buyerUserId string
		}
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given the parameters are valid a status code 200 is expected": {
			request: httptest.NewRequest(http.MethodPut, "/api/v1/bonds/buy", http.NoBody),
			urlParams: struct {
				bondId      string
				buyerUserId string
			}{
				bondId:      "ba1dc545-90a0-4501-af99-8a5944ca38c4",
				buyerUserId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return PutBondBuyerHandler(), nil
			},
			expectedStatusCode: http.StatusOK,
		},
		"Given the parameter'bond_id' is invalid a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPut, "/api/v1/bonds/buy", http.NoBody),
			urlParams: struct {
				bondId      string
				buyerUserId string
			}{
				bondId:      "ba1dc545-90a0-4501-af99",
				buyerUserId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return PutBondBuyerHandler(), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},

		"Given the parameter'buyer_user_id' is invalid a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPut, "/api/v1/bonds/buy", http.NoBody),
			urlParams: struct {
				bondId      string
				buyerUserId string
			}{
				bondId:      "ba1dc545-90a0-4501-af99-8a5944ca38c4",
				buyerUserId: "580b87da-e389-4290-f6191467f401",
			},
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				return PutBondBuyerHandler(), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			req := tsc.request
			req.Header.Set("Content-type", "application/json; charset=utf-8")
			req.SetPathValue("bond_id", tsc.urlParams.bondId)
			req.SetPathValue("buyer_user_id", tsc.urlParams.buyerUserId)

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
