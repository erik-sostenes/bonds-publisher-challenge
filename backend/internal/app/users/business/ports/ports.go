package ports

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
)

type (
	// right ports
	UserSaver interface {
		Save(context.Context, *domain.User) error
	}
	UserGetter interface {
		Get(context.Context, *domain.UserID) (*domain.User, uint8, error)
	}
)

type (
	// left ports
	UserCreator interface {
		Create(context.Context, *domain.User) error
	}

	TokenGenerator interface {
		Generate(*domain.User, uint8) (string, error)
	}

	UserAuthorizer interface {
		Authorize(context.Context, *domain.UserID, *domain.UserPassword) (string, error)
	}
)
