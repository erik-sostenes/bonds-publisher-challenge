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
)

func Test_BondSaver(t *testing.T) {
	type FactoryFunc func() (ports.BondSaver, *sql.DB)

	const sqlQueryDeleteBond = `DELETE FROM bonds WHERE id = $1`

	tdt := map[string]struct {
		bond          handlers.BondRequest
		factoryFunc   FactoryFunc
		expectedError error
	}{
		"Given a valid non-existing bond, it will be registered in postgresql": {
			bond: handlers.BondRequest{
				ID:             "1424e770-7aae-4a22-a743-7317b914082d",
				Name:           "Global Secure Bond 2024",
				QuantitySale:   1,
				SalesPrice:     400.0000,
				CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
				CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			factoryFunc: func() (ports.BondSaver, *sql.DB) {
				conn := db.PostgreSQLInjector()
				return NewBondSaver(conn), conn
			},
		},
		"Given an valid existing bond, it will not be registered in postgresql": {
			bond: handlers.BondRequest{
				ID:             "ba1dc545-90a0-4501-af99-8a5944ca38c4",
				Name:           "Global Secure Corporate Bond 2024",
				QuantitySale:   1,
				SalesPrice:     400.0000,
				CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
				CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			factoryFunc: func() (ports.BondSaver, *sql.DB) {
				conn := db.PostgreSQLInjector()
				return NewBondSaver(conn), conn
			},
			expectedError: domain.DuplicateBond,
		},
	}

	conn := db.PostgreSQLInjector()
	bondRequest := handlers.BondRequest{
		ID:             "ba1dc545-90a0-4501-af99-8a5944ca38c4",
		Name:           "Global Secure Corporate Bond 2024",
		QuantitySale:   1,
		SalesPrice:     400.0000,
		CreatorUserId:  "580b87da-e389-4290-acbf-f6191467f401",
		CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
	}

	// SetUP
	func() {
		saver := NewBondSaver(conn)
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
	}()

	t.Cleanup(func() {
		if _, err := conn.Exec(sqlQueryDeleteBond, bondRequest.ID); err != nil {
			t.Fatal(err)
		}
	})

	for name, tsc := range tdt {
		t.Run(name, func(t *testing.T) {
			saver, _ := tsc.factoryFunc()

			bond, err := domain.NewBond(
				tsc.bond.ID,
				tsc.bond.Name,
				tsc.bond.CreatorUserId,
				tsc.bond.CurrentOwnerId,
				tsc.bond.IsBought,
				tsc.bond.QuantitySale,
				tsc.bond.SalesPrice,
			)
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := conn.Exec(sqlQueryDeleteBond, bond.ID()); err != nil {
					t.Fatal(err)
				}
			})

			err = saver.Save(context.Background(), bond)
			asBond := domain.BondError(0)

			if errors.As(err, &asBond) {
				if !errors.Is(asBond, tsc.expectedError) {
					t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
				}
			} else if err != nil {
				t.Errorf("'%v' error was expected, but '%s' error was obtained", tsc.expectedError, asBond)
			}
		})
	}
}
