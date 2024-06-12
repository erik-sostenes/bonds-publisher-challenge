package memory

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/bonds-publisher-challenge/internal/app/users/business/domain"
	"github.com/erik-sostenes/bonds-publisher-challenge/pkg/set"
)

type userMemory struct {
	set set.Set[string, *domain.User]
}

func NewUserMemory() userMemory {
	return userMemory{
		set: set.Set[string, *domain.User]{},
	}
}

func (b *userMemory) Save(_ context.Context, user *domain.User) error {
	ok := b.set.Exist(user.ID())
	if ok {
		return fmt.Errorf("%w = user with id '%s' already exists", domain.DuplicateUser, user.ID())
	}

	b.set.Add(user.ID(), user)

	return nil
}
