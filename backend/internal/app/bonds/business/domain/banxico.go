package domain

// Banxico is a domain object that for the moment its only functionality is to transfer data
// it may need to be enriched further in the future, coupling it to a business logic
type (
	Dato struct {
		Fecha string
		Dato  string
	}
	Datos []*Dato

	Serie struct {
		IdSerie string
		Titulo  string
		Datos   Datos
	}

	Series []*Serie

	Bmx struct {
		Series Series
	}

	Banxico struct {
		Bmx Bmx
	}
)
