package bootstrap

import (
	"database/sql"
	"net/http"

	bondLgc "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/logic"
	bondPgrsql "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/postgresql"
	bondHlr "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/drives/handlers"
	usrLgc "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/logic"
	usrPgrsql "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/driven/postgresql"
	usrHlr "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/drives/handlers"
)

func BondInjector(db *sql.DB, mux *http.ServeMux) {
	bondSaver := bondPgrsql.NewBondSaver(db)
	bondCreator := bondLgc.NewBondCreator(bondSaver)

	bondOwnerUpdater := bondPgrsql.NewBondOwnerUpdater(db)
	bondBuyer := bondLgc.NewBondBuyer(bondOwnerUpdater)

	userBondsGetter := bondPgrsql.NewUserBondsGetter(db)
	userBondsRetriever := bondLgc.NewUserBondsRetriever(userBondsGetter)

	bondsGetter := bondPgrsql.NewBondsGetter(db)
	bondsRetriever := bondLgc.NewBondsRetriever(bondsGetter)

	bondHlr.BondHandler(bondCreator, bondBuyer, userBondsRetriever, bondsRetriever, mux)
}

func UserInjector(db *sql.DB, mux *http.ServeMux) {
	usrSaver := usrPgrsql.NewUserSaver(db)
	usrCreator := usrLgc.NewUserCreator(usrSaver)

	usrHlr.UserHandler(usrCreator, mux)
}
