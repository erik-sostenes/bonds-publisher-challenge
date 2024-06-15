package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func Baxico(ctx context.Context) (*BanxicoRequest, error) {
	const apiURL = "https://www.banxico.org.mx/SieAPIRest/service/v1/series/SF43718/datos/oportuno"
	apiToken := strings.TrimSpace(os.Getenv("BMX-TOKEN"))

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Bmx-Token", apiToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	serieResponse := &BanxicoRequest{}
	err = json.NewDecoder(response.Body).Decode(serieResponse)
	if err != nil {
		return nil, err
	}

	return serieResponse, nil
}
