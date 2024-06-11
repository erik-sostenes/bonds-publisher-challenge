package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
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

func (b *bondMemory) Save(ctx context.Context, bond *domain.Bond) error {
	ok := b.set.Exist(bond.ID())
	if ok {
		return fmt.Errorf("%w = Bond with id '%s' already exists", domain.DuplicateBond, bond.ID())
	}

	b.set.Add(bond.ID(), bond)

	return nil
}
