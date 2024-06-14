package logic

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/ports"
)

type userAuthorizer struct {
	userGetter     ports.UserGetter
	tokenGenerator ports.TokenGenerator
}

func NewUserAuthorizer(
	tokenGenerator ports.TokenGenerator,
	userGetter ports.UserGetter,
) ports.UserAuthorizer {
	if userGetter == nil {
		panic("missing 'User Getter' dependency")
	}

	if tokenGenerator == nil {
		panic("missing 'Token Generator' dependency")
	}

	return &userAuthorizer{
		userGetter:     userGetter,
		tokenGenerator: tokenGenerator,
	}
}

func (a *userAuthorizer) Authorize(ctx context.Context, username *domain.UserName, userPassword *domain.UserPassword) (token string, err error) {
	user, permissions, err := a.userGetter.Get(ctx, username)
	if err != nil {
		return
	}

	if err := user.PasswordMatches(userPassword); err != nil {
		return "", err
	}

	return a.tokenGenerator.Generate(user, permissions)
}
