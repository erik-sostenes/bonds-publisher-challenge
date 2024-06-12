package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/driven/memory"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
)

func Test_PostUserHandler(t *testing.T) {
	type HttpHandlerFunc func() (response.HttpHandlerFunc, error)

	tdt := map[string]struct {
		request            *http.Request
		handlerFunc        HttpHandlerFunc
		expectedStatusCode int
	}{
		"Given a valid non-existing user, a status code 201 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Erik Sostenes Simon",
					"password": "password",
					"role": {
						"id": 1,
						"type": "USER"
					}
				}
				`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusCreated,
		},
		"Given a valid existing user, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Erik Sostenes Simon",
					"password": "password",
					"role": {
						"id": 1,
						"type": "USER"
					}
				}
				`,
			)),
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
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given an invalid non-existing user, a status code 422 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", http.NoBody),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		"Given a non-existent invalid user role, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Erik Sostenes Simon",
					"password": "password",
					"role": {
						"id": 1,
						"type": "ANY"
					}
				}
				`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given a non-existent invalid user id, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "",
					"name": "Erik Sostenes Simon",
					"password": "password",
					"role": {
						"id": 1,
						"type": "ANY"
					}
				}
				`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given a non-existent invalid user name, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "",
					"password": "password",
					"role": {
						"id": 1,
						"type": "ANY"
					}
				}
				`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		"Given a non-existent invalid user password, a status code 400 is expected": {
			request: httptest.NewRequest(http.MethodPost, "/api/v1/register", strings.NewReader(
				`{
					"id": "ba1dc545-90a0-4501-af99-8a5944ca38c4",
					"name": "Erik Sostenes Simon",
					"password": "",
					"role": {
						"id": 1,
						"type": "ANY"
					}
				}
				`,
			)),
			handlerFunc: func() (response.HttpHandlerFunc, error) {
				memo := memory.NewUserMemory()
				userCtr := logic.NewUserCreator(&memo)

				return PostUserHandler(userCtr), nil
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
