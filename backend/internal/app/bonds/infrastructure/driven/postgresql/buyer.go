package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
)

type bondOwnerUpdater struct {
	conn *sql.DB
}

func NewBondOwnerUpdater(conn *sql.DB) ports.BondOwnerUpdater {
	return &bondOwnerUpdater{
		conn: conn,
	}
}

func (b *bondOwnerUpdater) Update(ctx context.Context, bID *domain.BondID, bCrOwId *domain.BondCurrentOwnerId) (err error) {
	tx, _ := b.conn.BeginTx(ctx, nil)

	err = transaction(tx, func() (err error) {
		const sqlQueryBondExist = `SELECT EXISTS(SELECT 1 FROM bonds WHERE id = $1)`

		var isExisting bool
		err = tx.QueryRow(sqlQueryBondExist, bID.ID()).Scan(&isExisting)
		if err != nil {
			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return errors.New("an error occurred while buying a bond")
		}

		if !isExisting {
			return fmt.Errorf("%w = Bond with id '%s' was not found", domain.BondNotFound, bID.ID())
		}

		const sqlQueryUpdateBondOwner = `UPDATE bonds
					SET current_owner_id = $1, is_bought = TRUE
					WHERE id = $2 AND is_bought = FALSE`

		result, err := b.conn.ExecContext(ctx, sqlQueryUpdateBondOwner,
			bCrOwId.ID(),
			bID.ID(),
		)
		if err != nil {
			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return errors.New("an error occurred while buying a bond")
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return errors.New("an error occurred while buying a bond")
		}

		if rowsAffected == 0 {
			return fmt.Errorf("%w = the bond with id '%s' has already been bought", domain.InvalidBondBought, bID.ID())
		}

		return
	})

	return
}
