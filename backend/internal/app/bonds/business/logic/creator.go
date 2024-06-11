package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
)

type bondCreator struct {
	saver ports.BondSaver
}

func NewBondCreator(saver ports.BondSaver) ports.BondCreator {
	if saver == nil {
		panic("missing Bond Saver dependency")
	}
	return &bondCreator{
		saver: saver,
	}
}

func (b *bondCreator) Create(ctx context.Context, bond *domain.Bond) error {
	return b.saver.Save(ctx, bond)
}
