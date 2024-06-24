package ports

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
)

type (
	// right ports
	BondSaver interface {
		Save(context.Context, *domain.Bond) error
	}

	BondOwnerUpdater interface {
		Update(context.Context, *domain.BondID, *domain.BondCurrentOwnerId) error
	}

	UserBondsGetter interface {
		Get(context.Context, *domain.BondCurrentOwnerId, *filter.Filter) (domain.Bonds, error)
	}

	BondsGetter interface {
		Get(context.Context, *domain.BondCurrentOwnerId, *filter.Filter) (domain.Bonds, error)
	}

	BanxicoSearcher interface {
		Search(context.Context) (*domain.Banxico, error)
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

	UserBondsRetriever interface {
		Retrieve(context.Context, *domain.BondCurrentOwnerId, *filter.Filter) (domain.Bonds, *domain.Banxico, error)
	}

	BondsRetriever interface {
		Retrieve(context.Context, *domain.BondCurrentOwnerId, *filter.Filter) (domain.Bonds, *domain.Banxico, error)
	}
)
