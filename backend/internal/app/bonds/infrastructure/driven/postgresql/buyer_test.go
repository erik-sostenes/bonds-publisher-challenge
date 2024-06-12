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

func Test_BondOwnerUpdater(t *testing.T) {
	type FactoryFunc func() (ports.BondOwnerUpdater, *sql.DB)

	const (
		sqlQueryDeleteBond = `DELETE FROM bonds WHERE id = $1`
		sqlQueryInsertUser = `INSERT INTO users(id, name, password) VALUES($1, $2, $3)`
		sqlQueryDeleteUser = `DELETE FROM users WHERE id = $1`
	)

	tdt := map[string]struct {
		bond          handlers.BondRequest
		factoryFunc   FactoryFunc
		expectedError error
	}{
		"Given an existing valid unpurchased bond, its buy will be registered in postgresql": {
			bond: handlers.BondRequest{
				ID:             "1424e770-7aae-4a22-a743-7317b914082d",
				CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			factoryFunc: func() (ports.BondOwnerUpdater, *sql.DB) {
				conn := db.PostgreSQLInjector()
				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					QuantitySale:   1,
					SalesPrice:     400.0000,
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
				return NewBondOwnerUpdater(conn), conn
			},
		},
		"Given a non-existing bond, your buy will not be registered in postgresql": {
			bond: handlers.BondRequest{
				ID:             "1424e770-7aae-4a22-a743-7317b914082d",
				CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			factoryFunc: func() (ports.BondOwnerUpdater, *sql.DB) {
				conn := db.PostgreSQLInjector()
				return NewBondOwnerUpdater(conn), conn
			},
			expectedError: domain.BondNotFound,
		},
		"Given an existing valid bond purchased, the purchase will not be registered in postgresql": {
			bond: handlers.BondRequest{
				ID:             "1424e770-7aae-4a22-a743-7317b914082d",
				CurrentOwnerId: "580b87da-e389-4290-acbf-f6191467f401",
			},
			factoryFunc: func() (ports.BondOwnerUpdater, *sql.DB) {
				conn := db.PostgreSQLInjector()
				saver := NewBondSaver(conn)

				bondRequest := handlers.BondRequest{
					ID:             "1424e770-7aae-4a22-a743-7317b914082d",
					Name:           "Global Secure Corporate Bond 2024",
					IsBought:       true,
					QuantitySale:   1,
					SalesPrice:     400.0000,
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
				return NewBondOwnerUpdater(conn), conn
			},
			expectedError: domain.InvalidBondBought,
		},
	}

	conn := db.PostgreSQLInjector()
	ctx := context.Background()

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
			updater, conn := tsc.factoryFunc()

			bId, err := domain.BondID(tsc.bond.ID).Validate()
			if err != nil {
				t.Fatal(err)
			}

			currentOwnerID, err := domain.BondCurrentOwnerId(tsc.bond.CurrentOwnerId).Validate()
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if _, err := conn.Exec(sqlQueryDeleteBond, bId.ID()); err != nil {
					t.Fatal(err)
				}
			})

			err = updater.Update(context.Background(), bId, currentOwnerID)
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
