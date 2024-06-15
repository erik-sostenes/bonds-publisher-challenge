// package in chargue of handling the http middlewares
package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/server/response"
	"golang.org/x/time/rate"
)

func RateLimiter(next http.HandlerFunc) http.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Minute/1000), 1)

	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			_ = response.JSON(w, http.StatusTooManyRequests, response.Response{
				Message: "Rate limit exceeded",
			})
			return
		}

		next.ServeHTTP(w, r)
	}
}

func Recovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Middleware] %s panic recovered:\n%s\n",
					time.Now().Format("2006/01/02 - 15:04:05"), err)

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}
}

// Logger shows every request made to the server
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`Host: %s, Uri: %s, Method: %s, Path: %s, User-Agent: %s`,
			r.Host,
			r.RequestURI,
			r.Method,
			r.URL.Path,
			r.UserAgent(),
		)

		next.ServeHTTP(w, r)
	}
}

// IsAuthenticated middleware that validates the token for each http request
// if the token is invalid the client is responded to with a StatusForbidden
//
// if the token is valid the requested HandlerFunc is executed
func IsAuthenticated(validator ports.TokenValidator, next http.HandlerFunc) http.HandlerFunc {
	if validator == nil {
		panic("missing 'Token Validator' dependency")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		strToken := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		_, err := validator.Validate(strToken)
		// TODO: validate permissions
		if err != nil {
			_ = response.JSON(w, http.StatusForbidden, response.Response{
				Message: err.Error(),
			})
			return
		}

		next(w, r)
	}
}
