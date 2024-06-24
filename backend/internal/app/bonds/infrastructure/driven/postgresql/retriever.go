package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
)

type userBondsGetter struct {
	conn *sql.DB
}

func NewUserBondsGetter(conn *sql.DB) ports.UserBondsGetter {
	return &userBondsGetter{
		conn: conn,
	}
}

func (b *userBondsGetter) Get(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	const sqlQueryGetUserBonds = `
		SELECT *
		FROM bonds
		WHERE current_owner_id = $1
		LIMIT $2
		OFFSET $3
	`

	limit := fltr.Stop()
	offset := fltr.Start()

	rows, err := b.conn.QueryContext(ctx, sqlQueryGetUserBonds, bcOwnerId.ID(), limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "database error", "msg", err.Error())
		return nil, errors.New("an error occurred while retrieving the bonds")
	}

	defer rows.Close()

	var bondsSchema BondsSchema

	for rows.Next() {
		var bondSchema BondSchema

		if err := rows.Scan(
			&bondSchema.ID,
			&bondSchema.Name,
			&bondSchema.QuantitySale,
			&bondSchema.SalesPrice,
			&bondSchema.IsBought,
			&bondSchema.CreatorUserId,
			&bondSchema.CurrentOwnerId,
		); err != nil {
			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return nil, errors.New("an error occurred while retrieving the bonds")
		}

		bondsSchema = append(bondsSchema, &bondSchema)
	}

	return bondsSchema.ToBusiness()
}

type bondsGetter struct {
	conn *sql.DB
}

func NewBondsGetter(conn *sql.DB) ports.BondsGetter {
	return &bondsGetter{
		conn: conn,
	}
}

func (b *bondsGetter) Get(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	const sqlQueryGetUserBonds = `
		SELECT *
		FROM bonds
		WHERE current_owner_id != $1 AND is_bought = FALSE
		LIMIT $2
		OFFSET $3
	`

	limit := fltr.Stop()
	offset := fltr.Start()

	rows, err := b.conn.QueryContext(ctx, sqlQueryGetUserBonds, bcOwnerId.ID(), limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "database error", "msg", err.Error())
		return nil, errors.New("an error occurred while retrieving the bonds")
	}

	defer rows.Close()

	var bondsSchema BondsSchema

	for rows.Next() {
		var bondSchema BondSchema

		if err := rows.Scan(
			&bondSchema.ID,
			&bondSchema.Name,
			&bondSchema.QuantitySale,
			&bondSchema.SalesPrice,
			&bondSchema.IsBought,
			&bondSchema.CreatorUserId,
			&bondSchema.CurrentOwnerId,
		); err != nil {
			slog.ErrorContext(ctx, "database error", "msg", err.Error())
			return nil, errors.New("an error occurred while retrieving the bonds")
		}

		bondsSchema = append(bondsSchema, &bondSchema)
	}

	return bondsSchema.ToBusiness()
}
