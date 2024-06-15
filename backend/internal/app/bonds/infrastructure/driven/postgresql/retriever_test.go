package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/erik-sostenes/bonds-publisher-challenge/cmd/bootstrap/db"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/infrastructure/drives/handlers"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
)

func Test_UserBondsGetter(t *testing.T) {
	type FactoryFunc func() (ports.UserBondsGetter, *sql.DB)

	const (
		sqlQueryInsertUser = `INSERT INTO users(id, name, password) VALUES($1, $2, $3)`
		sqlQueryDeleteUser = `DELETE FROM users WHERE id = $1`
		sqlQueryDeleteBond = `DELETE FROM bonds WHERE id = $1`
	)

	tdt := map[string]struct {
		currentOwnerId,
		bondId string
		factoryFunc FactoryFunc
		filter      struct {
			page, limit uint64
		}
		numberBondsExpected int
		expectedError       error
	}{
		"Given the user has existing valid bonds, a postgresql bond will be obtained": {
			currentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			bondId:         "1424e770-7aae-4a22-a743-7317b914082d",
			factoryFunc: func() (ports.UserBondsGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()

				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
					IsBought:       false,
					CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
					CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
				}

				bond, err := bondRequest.ToBusiness()
				if err != nil {
					t.Fatal(err)
				}

				if err := saver.Save(context.TODO(), bond); err != nil {
					t.Fatal(err)
				}

				return NewUserBondsGetter(conn), conn
			},
			filter: struct {
				page,
				limit uint64
			}{
				page:  1,
				limit: 25,
			},
			numberBondsExpected: 1,
		},
		"Given the user has no valid existing bonds, no postgresql bond will be obtained": {
			currentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			bondId:         "1424e770-7aae-4a22-a743-7317b914082d",
			factoryFunc: func() (ports.UserBondsGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()

				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
					IsBought:       true,
					CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
					CurrentOwnerId: "1148ab29-132b-4df7-9acc-b42a32c42a9f",
				}

				bond, err := bondRequest.ToBusiness()
				if err != nil {
					t.Fatal(err)
				}

				if err := saver.Save(context.TODO(), bond); err != nil {
					t.Fatal(err)
				}

				return NewUserBondsGetter(conn), conn
			},
			filter: struct {
				page,
				limit uint64
			}{
				page:  1,
				limit: 25,
			},
			numberBondsExpected: 0,
		},
	}

	conn := db.PostgreSQLInjector()
	ctx := context.Background()
	// setUp
	if err := func() (err error) {
		_, err = conn.ExecContext(ctx, sqlQueryInsertUser, "580b87da-e389-4290-acbf-f6191467f401", "Erik Sostenes Simon", "12345")
		if err != nil {
			return
		}

		_, err = conn.ExecContext(ctx, sqlQueryInsertUser, "1148ab29-132b-4df7-9acc-b42a32c42a9f", "Estefany Sostenes Simon", "12345")
		if err != nil {
			return
		}
		return
	}(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_, err := conn.ExecContext(ctx, sqlQueryDeleteUser, "580b87da-e389-4290-acbf-f6191467f401")
		if err != nil {
			t.Fatal(err)
		}
		_, err = conn.ExecContext(ctx, sqlQueryDeleteUser, "1148ab29-132b-4df7-9acc-b42a32c42a9f")
		if err != nil {
			t.Fatal(err)
		}
	})

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			getter, conn := tsc.factoryFunc()

			ownerId, err := domain.BondCurrentOwnerId(tsc.currentOwnerId).Validate()
			if err != nil {
				t.Fatal(err)
			}

			filter, err := filter.NewFilter(tsc.filter.page, tsc.filter.limit)
			if err != nil {
				t.Fatal(err)
			}

			bId, err := domain.BondID(tsc.bondId).Validate()
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := conn.Exec(sqlQueryDeleteBond, bId.ID()); err != nil {
					t.Fatal(err)
				}
			})

			bonds, err := getter.Get(context.Background(), ownerId, filter)

			asBond := domain.BondError(0)

			if errors.As(err, &asBond) {
				if !errors.Is(asBond, tsc.expectedError) {
					t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
				}
			} else if err != nil {
				t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
			}

			if len(bonds) != tsc.numberBondsExpected {
				t.Errorf("number of bonds expected to be %v, but %v was obtained", tsc.numberBondsExpected, len(bonds))
			}
		})
	}
}

func Test_BondsGetter(t *testing.T) {
	type FactoryFunc func() (ports.BondsGetter, *sql.DB)

	const (
		sqlQueryInsertUser = `INSERT INTO users(id, name, password) VALUES($1, $2, $3)`
		sqlQueryDeleteUser = `DELETE FROM users WHERE id = $1`
		sqlQueryDeleteBond = `DELETE FROM bonds WHERE id = $1`
	)

	tdt := map[string]struct {
		currentOwnerId,
		bondId string
		factoryFunc FactoryFunc
		filter      struct {
			page, limit uint64
		}
		numberBondsExpected int
		expectedError       error
	}{
		"Given the user has existing valid bonds, no postgresql bond will be obtained": {
			currentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			bondId:         "1424e770-7aae-4a22-a743-7317b914082d",
			factoryFunc: func() (ports.BondsGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()

				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
					IsBought:       true,
					CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
					CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
				}

				bond, err := domain.NewBond(
					bondRequest.ID,
					bondRequest.Name,
					bondRequest.CreatorUserId,
					bondRequest.CurrentOwnerId,
					bondRequest.IsBought,
					bondRequest.QuantitySale,
					bondRequest.SalesPrice,
				)
				if err != nil {
					t.Fatal(err)
				}

				if err := saver.Save(context.TODO(), bond); err != nil {
					t.Fatal(err)
				}

				return NewBondsGetter(conn), conn
			},
			filter: struct {
				page,
				limit uint64
			}{
				page:  1,
				limit: 25,
			},
			numberBondsExpected: 0,
		},
		"Given the user has no valid existing bonds, a postgresql bond will be obtained": {
			currentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			bondId:         "1424e770-7aae-4a22-a743-7317b914082d",
			factoryFunc: func() (ports.BondsGetter, *sql.DB) {
				conn := db.PostgreSQLInjector()

				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
					IsBought:       true,
					CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
					CurrentOwnerId: "1148ab29-132b-4df7-9acc-b42a32c42a9f",
				}

				bond, err := domain.NewBond(
					bondRequest.ID,
					bondRequest.Name,
					bondRequest.CreatorUserId,
					bondRequest.CurrentOwnerId,
					bondRequest.IsBought,
					bondRequest.QuantitySale,
					bondRequest.SalesPrice,
				)
				if err != nil {
					t.Fatal(err)
				}

				if err := saver.Save(context.TODO(), bond); err != nil {
					t.Fatal(err)
				}

				return NewBondsGetter(conn), conn
			},
			filter: struct {
				page,
				limit uint64
			}{
				page:  1,
				limit: 25,
			},
			numberBondsExpected: 0,
		},
	}

	conn := db.PostgreSQLInjector()
	ctx := context.Background()
	// setUp
	if err := func() (err error) {
		_, err = conn.ExecContext(ctx, sqlQueryInsertUser, "580b87da-e389-4290-acbf-f6191467f401", "Erik Sostenes Simon", "12345")
		if err != nil {
			return
		}

		_, err = conn.ExecContext(ctx, sqlQueryInsertUser, "1148ab29-132b-4df7-9acc-b42a32c42a9f", "Estefany Sostenes Simon", "12345")
		if err != nil {
			return
		}
		return
	}(); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_, err := conn.ExecContext(ctx, sqlQueryDeleteUser, "580b87da-e389-4290-acbf-f6191467f401")
		if err != nil {
			t.Fatal(err)
		}
		_, err = conn.ExecContext(ctx, sqlQueryDeleteUser, "1148ab29-132b-4df7-9acc-b42a32c42a9f")
		if err != nil {
			t.Fatal(err)
		}
	})

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			getter, conn := tsc.factoryFunc()

			ownerId, err := domain.BondCurrentOwnerId(tsc.currentOwnerId).Validate()
			if err != nil {
				t.Fatal(err)
			}

			filter, err := filter.NewFilter(tsc.filter.page, tsc.filter.limit)
			if err != nil {
				t.Fatal(err)
			}

			bId, err := domain.BondID(tsc.bondId).Validate()
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := conn.Exec(sqlQueryDeleteBond, bId.ID()); err != nil {
					t.Fatal(err)
				}
			})

			bonds, err := getter.Get(context.Background(), ownerId, filter)

			asBond := domain.BondError(0)

			if errors.As(err, &asBond) {
				if !errors.Is(asBond, tsc.expectedError) {
					t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
				}
			} else if err != nil {
				t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
			}

			if len(bonds) != tsc.numberBondsExpected {
				t.Errorf("number of bonds expected to be %v, but %v was obtained", tsc.numberBondsExpected, len(bonds))
			}
		})
	}
}
