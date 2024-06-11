package bootstrap

import (
	"net/http"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/memory"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/drives/handlers"
)

func BondInjector(mux *http.ServeMux) {
	memo := memory.NewBondMemory()

	bondCreator := logic.NewBondCreator(&memo)
	bondBuyer := logic.NewBondBuyer(&memo)
	userBondsRetriever := logic.NewUserBondsRetriever(&memo)

	handlers.BondHandler(bondCreator, bondBuyer, userBondsRetriever, mux)
}
