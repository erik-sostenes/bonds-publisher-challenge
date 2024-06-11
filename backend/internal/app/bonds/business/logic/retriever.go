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
	return &userBondsRetriever{
		userBondsGetter: userBondsGetter,
	}
}

func (r *userBondsRetriever) Retrieve(ctx context.Context, bcOwnerId *domain.BondCurrentOwnerId, fltr *filter.Filter) (domain.Bonds, error) {
	return r.userBondsGetter.Get(ctx, bcOwnerId, fltr)
}
