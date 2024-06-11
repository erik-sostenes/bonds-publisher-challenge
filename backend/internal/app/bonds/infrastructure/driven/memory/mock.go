package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/set"
)

type bondMemory struct {
	set set.Set[string, *domain.Bond]
}

func NewBondMemory() bondMemory {
	return bondMemory{
		set: set.Set[string, *domain.Bond]{},
	}
}

func (b *bondMemory) Save(_ context.Context, bond *domain.Bond) error {
	ok := b.set.Exist(bond.ID())
	if ok {
		return fmt.Errorf("%w = Bond with id '%s' already exists", domain.DuplicateBond, bond.ID())
	}

	b.set.Add(bond.ID(), bond)

	return nil
}

func (b *bondMemory) Update(_ context.Context, bID *domain.BondID, bcOwnerId *domain.BondCurrentOwnerId) error {
	ok := b.set.Exist(bID.ID())
	if !ok {
		return fmt.Errorf("%w = Bond with id '%s' was not found", domain.BondNotFound, bID.ID())
	}

	crtBond := b.set.GetByItem(bID.ID())

	newBond, err := domain.NewBond(
		crtBond.ID(),
		crtBond.Name(),
		crtBond.CreatorUserID(),
		bcOwnerId.ID(),
		!crtBond.IsBought(),
		crtBond.QuantitySale(),
		crtBond.SalesPrice(),
	)
	if err != nil {
		return err
	}

	b.set.Add(newBond.ID(), newBond)

	return nil
}

func (b *bondMemory) Get(_ context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	return b.set.GetAll(), nil
}
