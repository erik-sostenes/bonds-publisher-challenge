package bootstrap

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/drives/handlers"
)

func BondInjector(mux *http.ServeMux) {
	handlers.BondHandler(mux)
}
