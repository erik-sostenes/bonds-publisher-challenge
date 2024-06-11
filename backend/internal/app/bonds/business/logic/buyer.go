package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
)

type buyerBond struct {
	bondOwnerUpdater ports.BondOwnerUpdater
}

func NewBondBuyer(bondOwnerUpdater ports.BondOwnerUpdater) ports.BondBuyer {
	if bondOwnerUpdater == nil {
		panic("missing Bond Owner Updater dependency")
	}

	return &buyerBond{
		bondOwnerUpdater: bondOwnerUpdater,
	}
}

func (b *buyerBond) Buy(ctx context.Context, bID *domain.BondID, bcOwnerID *domain.BondCurrentOwnerId) error {
	return b.bondOwnerUpdater.Update(ctx, bID, bcOwnerID)
}
