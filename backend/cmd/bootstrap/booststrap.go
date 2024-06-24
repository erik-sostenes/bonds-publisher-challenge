package bootstrap

import (
	"database/sql"
	"net/http"
	"os"

	bondLgc "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/logic"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/banxico"
	bondPgrsql "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/driven/postgresql"
	bondHlr "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/drives/handlers"
	usrLgc "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/logic"
	usrPgrsql "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/driven/postgresql"
	usrHlr "github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/infrastructure/drives/handlers"
)

func BondInjector(db *sql.DB, mux *http.ServeMux) {
	bondSaver := bondPgrsql.NewBondSaver(db)
	bondCreator := bondLgc.NewBondCreator(bondSaver)

	bmx := banxico.NewBanxicoSearcher(os.Getenv("BMX_TOKEN"), os.Getenv("BMX_API_URL"))

	bondOwnerUpdater := bondPgrsql.NewBondOwnerUpdater(db)
	bondBuyer := bondLgc.NewBondBuyer(bondOwnerUpdater)

	userBondsGetter := bondPgrsql.NewUserBondsGetter(db)
	userBondsRetriever := bondLgc.NewUserBondsRetriever(userBondsGetter, bmx)

	bondsGetter := bondPgrsql.NewBondsGetter(db)
	bondsRetriever := bondLgc.NewBondsRetriever(bondsGetter, bmx)

	newTokenValidator := usrLgc.NewTokenValidator(os.Getenv("PUBLIC_KEY"))
	bondHlr.BondHandler(bondCreator, bondBuyer, userBondsRetriever, bondsRetriever, newTokenValidator, mux)
}

func UserInjector(db *sql.DB, mux *http.ServeMux) {
	usrSaver := usrPgrsql.NewUserSaver(db)
	usrCreator := usrLgc.NewUserCreator(usrSaver)

	usrGetter := usrPgrsql.NewUserGetter(db)
	tokenGenerator := usrLgc.NewTokenGenerator(os.Getenv("PRIVATE_KEY"))
	usrAuthorizer := usrLgc.NewUserAuthorizer(tokenGenerator, usrGetter)

	usrHlr.UserHandler(usrCreator, usrAuthorizer, mux)
}
