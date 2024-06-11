package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/ports"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/filter"
)

type userBondsRetriever struct {
	userBondsGetter ports.UserBondsGetter
}

func NewUserBondsRetriever(userBondsGetter ports.UserBondsGetter) ports.UserBondsRetriever {
	if userBondsGetter == nil {
		panic("missing User Bonds Getter dependency")
	}
	return &userBondsRetriever{
		userBondsGetter: userBondsGetter,
	}
}

func (r *userBondsRetriever) Retrieve(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	return r.userBondsGetter.Get(ctx, bcOwnerId, fltr)
}

type bondsRetriever struct {
	bondsGetter ports.BondsGetter
}

func NewBondsRetriever(bondsGetter ports.BondsGetter) ports.BondsRetriever {
	if bondsGetter == nil {
		panic("missing Bonds Getter dependency")
	}
	return &bondsRetriever{
		bondsGetter: bondsGetter,
	}
}

func (r *bondsRetriever) Retrieve(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	return r.bondsGetter.Get(ctx, bcOwnerId, fltr)
}
