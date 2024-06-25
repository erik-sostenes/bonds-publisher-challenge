package banxico

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
)

type banxicoSearcher struct {
	bmxToken, bmxApiURL string
}

func NewBanxicoSearcher(bmxToken, bmxApiURL string) ports.BanxicoSearcher {
	if strings.TrimSpace(bmxToken) == "" {
		panic("missing Bmx-Token")
	}

	if strings.TrimSpace(bmxApiURL) == "" {
		panic("missing Bmx-ApiURL")
	}

	return &banxicoSearcher{
		bmxToken:  bmxToken,
		bmxApiURL: bmxApiURL,
	}
}

func (s *banxicoSearcher) Search(ctx context.Context) (*domain.Banxico, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, s.bmxApiURL, nil)
	if err != nil {
		slog.ErrorContext(ctx, "banxico error", "msg", err.Error())
		return nil, errors.New("an error occurred while obtaining the SF43718 series from Banxico")
	}
	request.Header.Set("Bmx-Token", s.bmxToken)

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, err := client.Do(request)
	if err != nil {
		slog.ErrorContext(ctx, "banxico error", "msg", err.Error())
		return nil, errors.New("an error occurred while obtaining the SF43718 series from Banxico")
	}

	banxicoRequest := &BanxicoRequest{}
	err = json.NewDecoder(response.Body).Decode(banxicoRequest)
	if err != nil {
		slog.ErrorContext(ctx, "banxico error", "msg", err.Error())
		return nil, errors.New("an error occurred while obtaining the SF43718 series from Banxico")
	}

	return banxicoRequest.ToBusiness(), nil
}
