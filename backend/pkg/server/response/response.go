// response package in chargue of handling the http response and request
package response

import (
	"encoding/json"
	"mime"
	"net/http"
)

type (
	HttpHandlerFunc func(w http.ResponseWriter, r *http.Request) error

	Response struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func Bind(w http.ResponseWriter, r *http.Request, body any) error {
	content := r.Header.Get("Content-Type")

	if content == "" {
		return JSON(w, http.StatusUnsupportedMediaType, Response{
			Message: "missing content type",
		})
	}

	mediaType, _, err := mime.ParseMediaType(content)
	if err != nil {
		return JSON(w, http.StatusUnsupportedMediaType, Response{
			Message: err.Error(),
		})
	}

	switch mediaType {
	case "application/json; charset=utf-8", "application/json":
		err = json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			return JSON(w, http.StatusUnprocessableEntity, Response{
				Message: "the format of the body of the request is malformed",
			})
		}
	default:
		return JSON(w, http.StatusUnsupportedMediaType, Response{
			Message: "unsupported media type",
		})
	}

	return nil
}

func JSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
