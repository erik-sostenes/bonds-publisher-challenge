package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api", http.NoBody)
	request.Header.Set("Content-type", "application/json; charset=utf-8")

	response := httptest.NewRecorder()

	HealthCheck().ServeHTTP(response, request)

	expected := http.StatusOK
	if expected != response.Code {
		t.Log(response.Body.String())
		t.Errorf("status code was expected %d, but it was obtained %d", expected, response.Code)
	}
}
