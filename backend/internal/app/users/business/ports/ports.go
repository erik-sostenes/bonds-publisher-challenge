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
)

type (
	// left ports
	UserCreator interface {
		Create(context.Context, *domain.User) error
	}
)
