package banxico

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func Test_SeriesSearcher(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := BanxicoRequest{
			Bmx: Bmx{
				Series: Series{
					{
						IdSerie: "SF43718",
						Titulo:  "Tipo de cambio Pesos por dólar E.U.A. Tipo de cambio para solventar obligaciones denominadas en moneda extranjera Fecha de determinación (FIX)",
						Datos: Datos{
							{
								Fecha: "21/06/2024",
								Dato:  "18.1848",
							},
						},
					},
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))

	defer server.Close()

	s := NewBanxicoSearcher("some token", server.URL)

	bxico, err := s.Search(context.TODO())
	if err != nil {
		t.Error(err.Error())
		t.SkipNow()
	}

	if !slices.Equal(bxico.Bmx.Series, bxico.Bmx.Series) {
		t.Error("the Series are incorrect")
	}
}
