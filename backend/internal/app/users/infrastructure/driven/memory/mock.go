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

func (u *userMemory) Save(_ context.Context, user *domain.User) error {
	ok := u.set.Exist(user.ID())
	if ok {
		return fmt.Errorf("%w = user with id '%s' already exists", domain.DuplicateUser, user.ID())
	}
	newRole, _ := domain.NewRole(user.Role().ID(), user.Role().Type())
	newUser, _ := domain.NewUser(user.ID(), user.Name(), user.Password(), newRole)
	u.set.Add(user.ID(), newUser)

	return nil
}

func (u *userMemory) Get(_ context.Context, userId *domain.UserID) (*domain.User, uint8, error) {
	ok := u.set.Exist(userId.ID())
	if !ok {
		return nil, 0, fmt.Errorf("%w = User with id '%s' not found", domain.UserNotFound, userId.ID())
	}

	return u.set.GetByItem(userId.ID()), 1, nil
}
