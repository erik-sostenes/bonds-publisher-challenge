package ports

import (
	"context"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/bonds/business/domain"
)

type (
	// right ports
	CategorySaver interface {
		Save(context.Context, *domain.Bond) error
	}
)

type (
	// left ports
	BondCreator interface {
		Create(context.Context, *domain.Bond) error
	}
)
