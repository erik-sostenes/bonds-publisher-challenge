package banxico

import "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"

// DTO
type (
	Dato struct {
		Fecha string `json:"fecha"`
		Dato  string `json:"dato"`
	}
	Datos []*Dato

	Serie struct {
		IdSerie string `json:"idSerie"`
		Titulo  string `json:"titulo"`
		Datos   Datos  `json:"datos"`
	}
	Series []*Serie

	Bmx struct {
		Series Series `json:"series"`
	}

	BanxicoRequest struct {
		Bmx Bmx `json:"bmx"`
	}
)

func (b *BanxicoRequest) ToBusiness() *domain.Banxico {
	series := make(domain.Series, 0, len(b.Bmx.Series))

	for _, serie := range b.Bmx.Series {
		datos := make(domain.Datos, 0, len(serie.Datos))

		for _, dato := range serie.Datos {
			datos = append(datos, &domain.Dato{
				Fecha: dato.Fecha,
				Dato:  dato.Dato,
			})
		}

		series = append(series, &domain.Serie{
			IdSerie: serie.IdSerie,
			Titulo:  serie.Titulo,
			Datos:   datos,
		})
	}

	return &domain.Banxico{
		Bmx: domain.Bmx{
			Series: series,
		},
	}
}

func ToRequest(b *domain.Banxico) *BanxicoRequest {
	series := make(Series, 0, len(b.Bmx.Series))

	for _, serie := range b.Bmx.Series {
		datos := make(Datos, 0, len(serie.Datos))

		for _, dato := range serie.Datos {
			datos = append(datos, &Dato{
				Fecha: dato.Fecha,
				Dato:  dato.Dato,
			})
		}

		series = append(series, &Serie{
			IdSerie: serie.IdSerie,
			Titulo:  serie.Titulo,
			Datos:   datos,
		})
	}

	return &BanxicoRequest{
		Bmx: Bmx{
			Series: series,
		},
	}
}
