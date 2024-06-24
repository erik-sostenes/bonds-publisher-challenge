package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
)

type userBondsRetriever struct {
	userBondsGetter ports.UserBondsGetter
	banxicoSearcher ports.BanxicoSearcher
}

func NewUserBondsRetriever(userBondsGetter ports.UserBondsGetter, banxicoSearcher ports.BanxicoSearcher) ports.UserBondsRetriever {
	if userBondsGetter == nil {
		panic("missing 'User Bonds Getter' dependency")
	}

	if banxicoSearcher == nil {
		panic("missing 'Banxico Searcher' dependency")
	}

	return &userBondsRetriever{
		userBondsGetter: userBondsGetter,
		banxicoSearcher: banxicoSearcher,
	}
}

func (r *userBondsRetriever) Retrieve(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, *domain.Banxico, error) {
	banxico, err := r.banxicoSearcher.Search(ctx)
	if err != nil {
		return nil, nil, err
	}

	bonds, err := r.userBondsGetter.Get(ctx, bcOwnerId, fltr)
	if err != nil {
		return nil, nil, err
	}

	return bonds, banxico, nil
}

type bondsRetriever struct {
	bondsGetter     ports.BondsGetter
	banxicoSearcher ports.BanxicoSearcher
}

func NewBondsRetriever(bondsGetter ports.BondsGetter, banxicoSearcher ports.BanxicoSearcher) ports.BondsRetriever {
	if bondsGetter == nil {
		panic("missing Bonds Getter dependency")
	}

	if banxicoSearcher == nil {
		panic("missing 'Banxico Searcher' dependency")
	}

	return &bondsRetriever{
		bondsGetter:     bondsGetter,
		banxicoSearcher: banxicoSearcher,
	}
}

func (r *bondsRetriever) Retrieve(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, *domain.Banxico, error) {
	banxico, err := r.banxicoSearcher.Search(ctx)
	if err != nil {
		return nil, nil, err
	}

	bonds, err := r.bondsGetter.Get(ctx, bcOwnerId, fltr)
	if err != nil {
		return nil, nil, err
	}

	return bonds, banxico, nil
}
