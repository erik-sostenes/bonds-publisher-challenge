package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/driven/memory"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_GetAuthenticator(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given an existing valid user, a status code 200 with his token is expected": {
			request: httptest.NewRequest(
				http.MethodPost,
				"/api/v1/login?user_id=ba1dc545-90a0-4501-af99-8a5944ca38c4&user_password=password",
				http.NoBody,
			),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()

				userRequest := UserRequest{
					ID:       "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					Name:     "Erik Sostenes Simon",
					Password: "password",
					Role: RoleRequest{
						ID:   1,
						Type: "USER",
					},
				}

				user, err := userRequest.ToBusiness()
				if err != nil {
					t.Fatal(err)
				}

				if err := memo.Save(context.TODO(), user); err != nil {
					t.Fatal(err)
				}

				privateKey := os.Getenv("PRIVATE_KEY")
				tokenGenerator := logic.NewTokenGenerator(privateKey)
				authorizer := logic.NewUserAuthorizer(tokenGenerator, &memo)

				return GetAuthenticator(authorizer), nil
			},
			expectedStatusCode: http.StatusOK,
		},
		"Given an existing valid user sends an incorrect password, a 401 status code is expected": {
			request: httptest.NewRequest(
				http.MethodPost,
				"/api/v1/login?user_id=ba1dc545-90a0-4501-af99-8a5944ca38c4&user_password=anyPassword",
				http.NoBody,
			),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()

				userRequest := UserRequest{
					ID:       "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					Name:     "Erik Sostenes Simon",
					Password: "password",
					Role: RoleRequest{
						ID:   1,
						Type: "USER",
					},
				}

				user, err := userRequest.ToBusiness()
				if err != nil {
					t.Fatal(err)
				}

				if err := memo.Save(context.TODO(), user); err != nil {
					t.Fatal(err)
				}

				privateKey := os.Getenv("PRIVATE_KEY")
				tokenGenerator := logic.NewTokenGenerator(privateKey)
				authorizer := logic.NewUserAuthorizer(tokenGenerator, &memo)

				return GetAuthenticator(authorizer), nil
			},
			expectedStatusCode: http.StatusUnauthorized,
		},
		"Given a valid non-existent user, a 404 status code is expected": {
			request: httptest.NewRequest(
				http.MethodPost,
				"/api/v1/login?user_id=ba1dc545-90a0-4501-af99-8a5944ca38c4&user_password=password",
				http.NoBody,
			),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()

				privateKey := os.Getenv("PRIVATE_KEY")
				tokenGenerator := logic.NewTokenGenerator(privateKey)
				authorizer := logic.NewUserAuthorizer(tokenGenerator, &memo)

				return GetAuthenticator(authorizer), nil
			},
			expectedStatusCode: http.StatusNotFound,
		},
		"Given a user with invalid id, a status code 400 is expected": {
			request: httptest.NewRequest(
				http.MethodPost,
				"/api/v1/login?user_id=ba1dc545-4501-af99-8a5944ca38c4&user_password=password",
				http.NoBody,
			),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()

				privateKey := os.Getenv("PRIVATE_KEY")
				tokenGenerator := logic.NewTokenGenerator(privateKey)
				authorizer := logic.NewUserAuthorizer(tokenGenerator, &memo)

				return GetAuthenticator(authorizer), nil
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

			UserErrorHandler(h).ServeHTTP(resp, req)
			if tsc.expectedStatusCode != resp.Code {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", tsc.expectedStatusCode, resp.Code)
			}
		})
	}
}
