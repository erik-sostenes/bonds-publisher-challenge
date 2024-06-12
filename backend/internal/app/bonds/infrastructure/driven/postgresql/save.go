package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/lib/pq"
)

type bondSaver struct {
	conn *sql.DB
}

func NewBondSaver(conn *sql.DB) ports.BondSaver {
	return &bondSaver{
		conn: conn,
	}
}

func (b *bondSaver) Save(ctx context.Context, bond *domain.Bond) (err error) {
	const sqlQueryInsertBond = `INSERT INTO bonds(
									id,
									name,
									quantity_sale,
									sales_price,
									creator_user_id,
									current_owner_id
								) VALUES($1, $2, $3, $4, $5, $6)`

	_, err = b.conn.ExecContext(ctx, sqlQueryInsertBond,
		bond.ID(),
		bond.Name(),
		bond.QuantitySale(),
		bond.SalesPrice(),
		bond.CreatorUserID(),
		bond.CurrentOwnerID(),
	)
	// Check if the error is a unique constraint violation
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // 23505 is the PostgreSQL error code for unique violation
				return fmt.Errorf("%w = Bond with id '%s' already exists", domain.DuplicateBond, bond.ID())
			}
		}

		slog.ErrorContext(ctx, "database error", "error", err.Error())

		return errors.New("an error occurred while creating a bond")
	}

	return
}
