package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
)

type userCreator struct {
	saver ports.UserSaver
}

func NewUserCreator(saver ports.UserSaver) ports.UserCreator {
	if saver == nil {
		panic("missing User Saver dependency")
	}
	return &userCreator{
		saver: saver,
	}
}

func (u *userCreator) Create(ctx context.Context, user *domain.User) error {
	return u.saver.Save(ctx, user)
}
