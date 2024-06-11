package ports

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
)

type (
	// right ports
	BondSaver interface {
		Save(context.Context, *domain.Bond) error
	}

	BondOwnerUpdater interface {
		Update(context.Context, *domain.BondID, *domain.BondCurrentOwnerId) error
	}
)

type (
	// left ports
	BondCreator interface {
		Create(context.Context, *domain.Bond) error
	}

	BondBuyer interface {
		Buy(context.Context, *domain.BondID, *domain.BondCurrentOwnerId) error
	}
)
